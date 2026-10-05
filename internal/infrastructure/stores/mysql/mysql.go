/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: MySQL 客户端初始化
 *   - 从 viper 读 toml 的 [mysql] 段,拼成 gdb.Config
 *   - 内部复用 pkg/stores/gdb,保持 ORM 层的一致性
 *   - 读 address(host:port)格式,内部拆分
 **/

package mysql

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/weitrue/Seckill/pkg/stores/gdb"
)

// 默认值
const (
	defaultMaxIdle     = 10
	defaultMaxOpen     = 100
	defaultLogLevel    = "warn"
	defaultMaxLifetime = 300 // 秒
)

var (
	db     *gorm.DB
	initMu sync.Mutex
)

/*Init
 *@Description: 初始化 MySQL 连接池;幂等
 *@return error
 */
func Init() error {
	logType := "MysqlInit"
	initMu.Lock()
	defer initMu.Unlock()
	if db != nil {
		return nil
	}

	addr := viper.GetString("mysql.address")
	if addr == "" {
		return errors.New("mysql.address is empty")
	}
	host, port, err := parseHostPort(addr)
	if err != nil {
		return fmt.Errorf("parse mysql.address %q: %w", addr, err)
	}

	cfg := &gdb.Config{
		User:               viper.GetString("mysql.username"),
		Password:           viper.GetString("mysql.password"),
		Host:               host,
		Port:               port,
		Database:           viper.GetString("mysql.database"),
		MaxIdleConns:       viper.GetInt("mysql.max_idle_conns"),
		MaxOpenConns:       viper.GetInt("mysql.max_open_conns"),
		MaxConnMaxLifetime: viper.GetInt64("mysql.max_conn_max_lifetime"),
		LogLevel:           viper.GetString("mysql.log_level"),
	}
	if cfg.Database == "" {
		return errors.New("mysql.database is empty")
	}
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = defaultMaxIdle
	}
	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = defaultMaxOpen
	}
	if cfg.MaxConnMaxLifetime <= 0 {
		cfg.MaxConnMaxLifetime = defaultMaxLifetime
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = defaultLogLevel
	}

	db, err = gdb.NewDB(cfg)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, addr:%s, database:%s",
			logType, err.Error(), addr, cfg.Database)
		return err
	}
	logrus.Infof("logType:%s, msg:connected, addr:%s, database:%s, maxIdle:%d, maxOpen:%d",
		logType, addr, cfg.Database, cfg.MaxIdleConns, cfg.MaxOpenConns)
	return nil
}

/*GetDB
 *@Description: 获取 *gorm.DB;未初始化返回 nil
 *@return *gorm.DB
 */
func GetDB() *gorm.DB {
	return db
}

/*Close
 *@Description: 关闭底层连接池;幂等
 *@return error
 */
func Close() error {
	initMu.Lock()
	defer initMu.Unlock()
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	db = nil
	return sqlDB.Close()
}

// parseHostPort 拆分 "host:port" 格式
func parseHostPort(addr string) (string, int, error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port %q: %w", portStr, err)
	}
	return host, port, nil
}

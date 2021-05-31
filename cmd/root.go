/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:26
 * Description:
 **/

package cmd

import (
	"fmt"
	"os"

	"github.com/weitrue/Seckill/infrastructure/config/cluster"
	"github.com/weitrue/Seckill/infrastructure/stores/etcd"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seckill",
	Short: "Seckill server.",
	Long:  `Seckill server.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logrus.Info(cfgFile)
}

func init() {
	// 设置initConfig在调用rootCmd的Execute()方法时运行
	logrus.Info("------------------ init config -------------------")
	cobra.OnInitialize(initConfig)
	flags := rootCmd.PersistentFlags()
	logrus.Info("------------------ find config -------------------")
	flags.StringVarP(&cfgFile, "config", "c", "./config/Seckill.toml", "config file (default is $HOME/.Seckill.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// 从flag中获取配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 主目录
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 从主目录下搜索后缀名为".seckill"文件 (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".seckill")
	}
	logrus.Info("------------------ load env ----------------------")
	viper.AutomaticEnv() // 读取匹配的环境变量
	logrus.Info("------------------ load config -------------------")
	// 读取找到的配置文件
	if err := viper.ReadInConfig(); err == nil {
		logrus.Info("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}
	logrus.Info("------------------ init etcd ---------------------")
	// 初始化 etcd 客户端连接
	if err := etcd.Init(); err != nil {
		panic(err)
	}
	logrus.Info("------------------ init logger -------------------")
	// logger.InitLogger()
	logrus.Info("------------------ watch cluster ------------------")
	// 监听集群配置信息
	if err := cluster.WatchClusterConfig(); err != nil {
		panic(err)
	}
}

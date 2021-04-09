/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:26
 * Description:
 **/

package cmd

import (
	"Seckill/infrastructure/stores/etcd"
	"fmt"
	"os"

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
	cobra.OnInitialize(initConfig)
	flags := rootCmd.PersistentFlags()
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

	viper.AutomaticEnv() // 读取匹配的环境变量

	// 读取找到的配置文件
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}

	// 初始化 etcd 客户端连接
	if err := etcd.Init(); err != nil {
		panic(err)
	}

	//logger.InitLogger()
	//if err := cluster.WatchClusterConfig(); err != nil {
	//	panic(err)
	//}
}

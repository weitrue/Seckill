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
	//Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&cfgFile, "config", "c", "./config/Seckill.toml", "config file (default is $HOME/.Seckill.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".seckill" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".seckill")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}
	//if err := etcd.Init(); err != nil {
	//	panic(err)
	//}
	//logger.InitLogger()
	//if err := cluster.WatchClusterConfig(); err != nil {
	//	panic(err)
	//}
}
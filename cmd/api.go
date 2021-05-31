/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:08
 * Description:
 **/

package cmd

import (
	"github.com/weitrue/Seckill/interfaces/api"
	"github.com/weitrue/Seckill/interfaces/rpc"

	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// apiCmd api服务命令
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Seckill api server.",
	Long:  `Seckill api server.`,
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		wg.Add(2)
		// api退出信号通知chan
		onApiExit := make(chan error, 1)
		// rpc退出信号通知chan
		onRpcExit := make(chan error, 1)
		go func() {
			defer wg.Done()
			if err := api.Run(); err != nil {
				logrus.Error(err)
				onApiExit <- err
			}
			//
			close(onApiExit)
		}()
		go func() {
			defer wg.Done()
			if err := rpc.Run(); err != nil {
				logrus.Error(err)
				onRpcExit <- err
			}
			//
			close(onRpcExit)
		}()
		// 信号通知chan
		onSignal := make(chan os.Signal)
		// 优雅退出
		// 监听指定信号 ctrl+c 结束程序(可以被捕获、阻塞或忽略)
		signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-onSignal:
			logrus.Info("exit by signal ", sig)
			api.Exit()
			rpc.Exit()
		case err := <-onApiExit:
			rpc.Exit()
			logrus.Info("exit by error ", err)
		case err := <-onRpcExit:
			api.Exit()
			logrus.Info("exit by error ", err)
		}
		wg.Wait()
	},
}

func init() {
	// 将api初始化命令添加到主命令中
	rootCmd.AddCommand(apiCmd)
}

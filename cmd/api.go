/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:08
 * Description:
 **/

package cmd

import (
	"Seckill/interfaces/api"
	"Seckill/interfaces/rpc"

	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Seckill api server.",
	Long:  `Seckill api server.`,
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		wg.Add(2)
		onApiExit := make(chan error, 1)
		onRpcExit := make(chan error, 1)
		go func() {
			defer wg.Done()
			if err := api.Run(); err != nil {
				onApiExit <- err
			}
			close(onApiExit)
		}()
		go func() {
			defer wg.Done()
			if err := rpc.Run(); err != nil {
				logrus.Error(err)
				onRpcExit <- err
			}
			close(onRpcExit)
		}()
		onSignal := make(chan os.Signal)
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
	rootCmd.AddCommand(apiCmd)
}


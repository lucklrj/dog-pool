package listen

import (
	"github.com/fatih/color"
	"os"
	"os/signal"
	"service/mysql"
	"syscall"
)

func StartSignalListen() {
	listenSignal := make(chan os.Signal)
	signal.Notify(listenSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for Signal := range listenSignal {
			switch Signal {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				RpcListener.Close()

				if len(mysql.CoonPool) > 0 {
					color.Green("正在关闭Mysql连接")
					for _, mysqlCoon := range mysql.CoonPool {
						mysqlCoon.Close()
					}
					color.Green("Mysql连接关闭结束。")
				}

				//todo 关闭memcache
				//todo 关闭redis

				os.Exit(0)
			}
		}
	}()
}

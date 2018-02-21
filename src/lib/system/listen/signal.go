package listen

import (
	"github.com/fatih/color"
	"os"
	"os/signal"
	"service/mysql"
	"service/memcache"
	"service/redis"
	"syscall"
	"lib/system"
)

func StartSignalListen() {
	listenSignal := make(chan os.Signal)
	signal.Notify(listenSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for Signal := range listenSignal {
			switch Signal {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				system.Horizontaline()
				if len(mysql.CoonPool) > 0 {
					color.Green("正在关闭Mysql连接")
					for _, mysqlCoon := range mysql.CoonPool {
						mysqlCoon.Close()
					}
					color.Green("成功关闭Mysql连接。")
					system.Horizontaline()
				}

				//关闭memcache
				if len(memcache.MyMemcachePool.Pools)>0{
					color.Green("正在关闭Memcache连接")
					memcache.MyMemcachePool.Close()
					color.Green("成功关闭Memcache连接。")
					system.Horizontaline()
				}

				//关闭redis
				if len(redis.MyRedisPool.Pools)>0{
					color.Green("正在关闭Redis连接")
					redis.MyRedisPool.Close()
					color.Green("成功关闭Redis连接。")
					system.Horizontaline()
				}
				color.Green("正在关闭RPC服务")
				RpcListener.Close()
				color.Green("成功关闭RRC服务")
				system.Horizontaline()
				os.Exit(0)
			}
		}
	}()
}

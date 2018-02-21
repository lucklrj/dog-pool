package main

import (
	"net/rpc"

	"service/memcache"
	"service/mysql"
	"service/redis"

	"lib/system/listen"
)

func main() {
	//信号监听
	listen.StartSignalListen()

	//注册服务
	rpc.Register(new(mysql.Mysql))
	rpc.Register(new(memcache.Memcache))
	rpc.Register(new(redis.Redis))

	//go func() {
	//	for {
	//		time.Sleep(5 * time.Second)
	//		fmt.Println(redis.MyRedisPool.Stats())
	//	}
	//}()

	//监听客户端连接
	listen.StarRpctListen()

}

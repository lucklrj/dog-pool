package listen

import (
	"fmt"
	"github.com/fatih/color"
	"lib/config"
	"lib/system"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"strconv"
)

var RpcListener *net.TCPListener

func StarRpctListen() {
	//监听连接
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(config.SystemConfig.Port))
	system.Exit(err)

	RpcListener, err = net.ListenTCP("tcp", tcpAddr)
	system.Exit(err)

	color.Green("服务运行中,监听端口：" + strconv.Itoa(config.SystemConfig.Port))

	for {
		conn, err := RpcListener.Accept()
		if err != nil {
			fmt.Fprint(os.Stderr, "accept err: %s", err.Error())
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

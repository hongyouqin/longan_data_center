package main

import (
	"google.golang.org/grpc"
	"log"
	"longan_data_center/config"
	"longan_data_center/proto"
	"longan_data_center/server"
	"net"
)

func main() {
	addr := config.DCConfig.Server.ListenAddr
	lisen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	log.Println("龙眼数据中心服务启动")
	log.Println("服务地址: ", addr)

	//设置最大消息接收字节为20M
	option := grpc.MaxHeaderListSize(20971520);
	s := grpc.NewServer(option)
	proto.RegisterLonganDataCenterServer(s, &server.DataCenterServer{})
	s.Serve(lisen)
	log.Println("服务退出")
}

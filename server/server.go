package server

import (
	"context"
	"fmt"
	"log"
	"longan_data_center/proto"
	"longan_data_center/sql"
)

//数据中心远程服务
type DataCenterServer struct {
	//用于通知的客户端
	clients []proto.LonganDataCenter_RegisterClientServer
}

// 存储员工人脸信息
func (s *DataCenterServer) StorageEmployeeFace(ct context.Context, user *proto.UserParam) (*proto.StorageReply, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("录入出错: ", r)
		}
	}()

	var dc sql.DC
	dc.Connect()
	dc.SingularTable(true)
	defer dc.Close()

	var employee sql.FacesRegisterTable
	employee.UserId = user.UserId
	employee.Name = user.Name
	employee.Photo = user.FacePhoto
	employee.Feature = user.FaceFeature
	employee.Size = user.FaceSize
	employee.AcqTime = user.AcquisitionTime
	employee.Score = user.FaceScore
	result := dc.AddRecord(&employee)

	var reply proto.StorageReply
	reply.IsSuccess = result

	if result {
		s.NotifyUpdate();
	}

	fmt.Println("user=", user.Name, "入库");

	return &reply, nil
}
// 存储陌生人脸信息
func (s *DataCenterServer) StorageStrangerFace(context.Context, *proto.UserParam) (*proto.StorageReply, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("陌生人录入出错: ", r)
		}
	}()
	fmt.Println("strange face");
	var reply proto.StorageReply
	reply.IsSuccess = true
	return &reply, nil
}

// 拉取人脸注册表的所有数据
func (s *DataCenterServer) ExtractFaceRegTableDatas(empty *proto.Empty, stream proto.LonganDataCenter_ExtractFaceRegTableDatasServer) error {

	defer func() {
		if r := recover(); r != nil {
			log.Println("拉取员工数据出错: ", r)
		}
	}()

	fmt.Println("拉取员工数据")

	var dc sql.DC
	if err := dc.Connect(); err != nil {
		return err
	}
	defer dc.Close()
	dc.SingularTable(true)
	employees := dc.GetAllEmployeesRecord();

	for _, e := range employees {
		result := proto.UserParam{}
		result.Name = e.Name
		result.UserId = e.UserId
		result.FacePhoto = e.Photo
		result.FaceFeature = e.Feature
		result.FaceSize = e.Size
		result.AcquisitionTime = e.AcqTime
		result.FaceScore = e.Score
		var feature proto.Feature
		feature.Face = &result
		if err := stream.Send(&feature); err != nil {
			return err
		}
	}
	return nil
}

// 注册客户端，用于之后的更新通知接收
func (s *DataCenterServer) RegisterClient(empty *proto.Empty, stream proto.LonganDataCenter_RegisterClientServer) error {

	s.clients = append(s.clients, stream)
	log.Println(stream,"新客户端注册")
	ctx := stream.Context()
	for {
		select {
			case <- ctx.Done():
				log.Println(stream,"客户端断开连接")
				//移除客户端
				for index, client := range s.clients {
					if client == stream {
						s.clients = append(s.clients[:index], s.clients[index+1:]...)
						log.Println(stream, "客户端移除成功")
					}
				}

				return ctx.Err()
		}
	}

	return nil
}

//通知所有的客户端有更新
func (s *DataCenterServer) NotifyUpdate() error {
	log.Println("需要通知", len(s.clients), "个客户")

	for _, client := range s.clients {
		var notify proto.UpdateNotity
		notify.EmployeeTableUpdate = true
		client.Send(&notify)
	}

	return nil;
}

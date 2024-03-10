package services

import (
	"gitlab.zixel.cn/go/framework"
	"google.golang.org/grpc"
)

var conns = make([]*grpc.ClientConn, 0, 10)

var StorageServiceService FileManagerServiceClient
var UserServiceV2 UserServiceV2Client
var OrgMagService OrgMagServiceClient

func Init() {
	StorageServiceService = connectStorageServiceService()
	UserServiceV2 = connectUserServiceV2()
	OrgMagService = connectOrgMagService()
}

func connectStorageServiceService() FileManagerServiceClient {
	conn := framework.GetGrpcConnection("storage")

	Service := NewFileManagerServiceClient(conn)

	conns = append(conns, conn)
	return Service
}

func connectUserServiceV2() UserServiceV2Client {
	conn := framework.GetGrpcConnection("user")

	Service := NewUserServiceV2Client(conn)

	conns = append(conns, conn)
	return Service
}

func connectOrgMagService() OrgMagServiceClient {
	conn := framework.GetGrpcConnection("organization")

	Service := NewOrgMagServiceClient(conn)

	conns = append(conns, conn)
	return Service
}

func ExitGrpc() {
	for _, conn := range conns {
		conn.Close()
	}
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: OrgMagService.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrgMagServiceClient is the client API for OrgMagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgMagServiceClient interface {
	// 创建公司
	CreateCompany(ctx context.Context, in *CompanyCreateRequest, opts ...grpc.CallOption) (*CompanyCreateReply, error)
	// 删除公司
	DeleteCompany(ctx context.Context, in *CompanyDeleteRequest, opts ...grpc.CallOption) (*CompanyDeleteReply, error)
	// 更新公司
	UpdateCompany(ctx context.Context, in *CompanyUpdateRequest, opts ...grpc.CallOption) (*CompanyUpdateReply, error)
	// 查询公司
	QueryCompany(ctx context.Context, in *C2S_QueryCompanyReqT, opts ...grpc.CallOption) (*C2S_QueryCompanyRpnT, error)
	// 查询用户公司列表
	GetCompanyListById(ctx context.Context, in *GetCompanyListByIdRequest, opts ...grpc.CallOption) (*GetCompanyListByIdReply, error)
	GetCompanyListByUid(ctx context.Context, in *C2S_GetCompanyListByUidReqT, opts ...grpc.CallOption) (*GetCompanyListByIdReply, error)
	// 根据openId查询用户信息
	ListUserByOpenId(ctx context.Context, in *C2S_ListUserByOpenIdReq, opts ...grpc.CallOption) (*C2S_ListUserByOpenIdRpn, error)
	// 根据openId查询用户信息
	ListUserByUid(ctx context.Context, in *C2S_ListUserByUIdReq, opts ...grpc.CallOption) (*C2S_ListUserByOpenIdRpn, error)
	// 添加企业用户
	AddCompanyUser(ctx context.Context, in *C2S_AddCompanyUserReq, opts ...grpc.CallOption) (*C2S_AddCompanyUserRpn, error)
	// 添加企业外部联系人
	AddCompanyOutsideUser(ctx context.Context, in *C2S_AddCompanyOutsideUserReq, opts ...grpc.CallOption) (*C2S_AddCompanyOutsideUserRpn, error)
	DeleteUser(ctx context.Context, in *C2S_DeleteUserReq, opts ...grpc.CallOption) (*C2S_DeleteUserRpn, error)
	QueryUsers(ctx context.Context, in *C2S_UserQueryReq, opts ...grpc.CallOption) (*C2S_UserQueryRpn, error)
	// 根据企业id获得企业信息
	GetCompanyById(ctx context.Context, in *C2S_GetCompanyByIdReq, opts ...grpc.CallOption) (*CompanyMsg, error)
	// 添加已经存在的用户到企业中
	AddExistUser(ctx context.Context, in *C2S_AddExistUserReq, opts ...grpc.CallOption) (*C2S_AddExistUserRpn, error)
}

type orgMagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgMagServiceClient(cc grpc.ClientConnInterface) OrgMagServiceClient {
	return &orgMagServiceClient{cc}
}

func (c *orgMagServiceClient) CreateCompany(ctx context.Context, in *CompanyCreateRequest, opts ...grpc.CallOption) (*CompanyCreateReply, error) {
	out := new(CompanyCreateReply)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) DeleteCompany(ctx context.Context, in *CompanyDeleteRequest, opts ...grpc.CallOption) (*CompanyDeleteReply, error) {
	out := new(CompanyDeleteReply)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) UpdateCompany(ctx context.Context, in *CompanyUpdateRequest, opts ...grpc.CallOption) (*CompanyUpdateReply, error) {
	out := new(CompanyUpdateReply)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) QueryCompany(ctx context.Context, in *C2S_QueryCompanyReqT, opts ...grpc.CallOption) (*C2S_QueryCompanyRpnT, error) {
	out := new(C2S_QueryCompanyRpnT)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/QueryCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) GetCompanyListById(ctx context.Context, in *GetCompanyListByIdRequest, opts ...grpc.CallOption) (*GetCompanyListByIdReply, error) {
	out := new(GetCompanyListByIdReply)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/GetCompanyListById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) GetCompanyListByUid(ctx context.Context, in *C2S_GetCompanyListByUidReqT, opts ...grpc.CallOption) (*GetCompanyListByIdReply, error) {
	out := new(GetCompanyListByIdReply)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/GetCompanyListByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) ListUserByOpenId(ctx context.Context, in *C2S_ListUserByOpenIdReq, opts ...grpc.CallOption) (*C2S_ListUserByOpenIdRpn, error) {
	out := new(C2S_ListUserByOpenIdRpn)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/listUserByOpenId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) ListUserByUid(ctx context.Context, in *C2S_ListUserByUIdReq, opts ...grpc.CallOption) (*C2S_ListUserByOpenIdRpn, error) {
	out := new(C2S_ListUserByOpenIdRpn)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/listUserByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) AddCompanyUser(ctx context.Context, in *C2S_AddCompanyUserReq, opts ...grpc.CallOption) (*C2S_AddCompanyUserRpn, error) {
	out := new(C2S_AddCompanyUserRpn)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/addCompanyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) AddCompanyOutsideUser(ctx context.Context, in *C2S_AddCompanyOutsideUserReq, opts ...grpc.CallOption) (*C2S_AddCompanyOutsideUserRpn, error) {
	out := new(C2S_AddCompanyOutsideUserRpn)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/addCompanyOutsideUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) DeleteUser(ctx context.Context, in *C2S_DeleteUserReq, opts ...grpc.CallOption) (*C2S_DeleteUserRpn, error) {
	out := new(C2S_DeleteUserRpn)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/deleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) QueryUsers(ctx context.Context, in *C2S_UserQueryReq, opts ...grpc.CallOption) (*C2S_UserQueryRpn, error) {
	out := new(C2S_UserQueryRpn)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/queryUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) GetCompanyById(ctx context.Context, in *C2S_GetCompanyByIdReq, opts ...grpc.CallOption) (*CompanyMsg, error) {
	out := new(CompanyMsg)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/getCompanyById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgMagServiceClient) AddExistUser(ctx context.Context, in *C2S_AddExistUserReq, opts ...grpc.CallOption) (*C2S_AddExistUserRpn, error) {
	out := new(C2S_AddExistUserRpn)
	err := c.cc.Invoke(ctx, "/structure.OrgMagService/addExistUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgMagServiceServer is the server API for OrgMagService service.
// All implementations must embed UnimplementedOrgMagServiceServer
// for forward compatibility
type OrgMagServiceServer interface {
	// 创建公司
	CreateCompany(context.Context, *CompanyCreateRequest) (*CompanyCreateReply, error)
	// 删除公司
	DeleteCompany(context.Context, *CompanyDeleteRequest) (*CompanyDeleteReply, error)
	// 更新公司
	UpdateCompany(context.Context, *CompanyUpdateRequest) (*CompanyUpdateReply, error)
	// 查询公司
	QueryCompany(context.Context, *C2S_QueryCompanyReqT) (*C2S_QueryCompanyRpnT, error)
	// 查询用户公司列表
	GetCompanyListById(context.Context, *GetCompanyListByIdRequest) (*GetCompanyListByIdReply, error)
	GetCompanyListByUid(context.Context, *C2S_GetCompanyListByUidReqT) (*GetCompanyListByIdReply, error)
	// 根据openId查询用户信息
	ListUserByOpenId(context.Context, *C2S_ListUserByOpenIdReq) (*C2S_ListUserByOpenIdRpn, error)
	// 根据openId查询用户信息
	ListUserByUid(context.Context, *C2S_ListUserByUIdReq) (*C2S_ListUserByOpenIdRpn, error)
	// 添加企业用户
	AddCompanyUser(context.Context, *C2S_AddCompanyUserReq) (*C2S_AddCompanyUserRpn, error)
	// 添加企业外部联系人
	AddCompanyOutsideUser(context.Context, *C2S_AddCompanyOutsideUserReq) (*C2S_AddCompanyOutsideUserRpn, error)
	DeleteUser(context.Context, *C2S_DeleteUserReq) (*C2S_DeleteUserRpn, error)
	QueryUsers(context.Context, *C2S_UserQueryReq) (*C2S_UserQueryRpn, error)
	// 根据企业id获得企业信息
	GetCompanyById(context.Context, *C2S_GetCompanyByIdReq) (*CompanyMsg, error)
	// 添加已经存在的用户到企业中
	AddExistUser(context.Context, *C2S_AddExistUserReq) (*C2S_AddExistUserRpn, error)
	mustEmbedUnimplementedOrgMagServiceServer()
}

// UnimplementedOrgMagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrgMagServiceServer struct {
}

func (UnimplementedOrgMagServiceServer) CreateCompany(context.Context, *CompanyCreateRequest) (*CompanyCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedOrgMagServiceServer) DeleteCompany(context.Context, *CompanyDeleteRequest) (*CompanyDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedOrgMagServiceServer) UpdateCompany(context.Context, *CompanyUpdateRequest) (*CompanyUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedOrgMagServiceServer) QueryCompany(context.Context, *C2S_QueryCompanyReqT) (*C2S_QueryCompanyRpnT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCompany not implemented")
}
func (UnimplementedOrgMagServiceServer) GetCompanyListById(context.Context, *GetCompanyListByIdRequest) (*GetCompanyListByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyListById not implemented")
}
func (UnimplementedOrgMagServiceServer) GetCompanyListByUid(context.Context, *C2S_GetCompanyListByUidReqT) (*GetCompanyListByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyListByUid not implemented")
}
func (UnimplementedOrgMagServiceServer) ListUserByOpenId(context.Context, *C2S_ListUserByOpenIdReq) (*C2S_ListUserByOpenIdRpn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserByOpenId not implemented")
}
func (UnimplementedOrgMagServiceServer) ListUserByUid(context.Context, *C2S_ListUserByUIdReq) (*C2S_ListUserByOpenIdRpn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserByUid not implemented")
}
func (UnimplementedOrgMagServiceServer) AddCompanyUser(context.Context, *C2S_AddCompanyUserReq) (*C2S_AddCompanyUserRpn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCompanyUser not implemented")
}
func (UnimplementedOrgMagServiceServer) AddCompanyOutsideUser(context.Context, *C2S_AddCompanyOutsideUserReq) (*C2S_AddCompanyOutsideUserRpn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCompanyOutsideUser not implemented")
}
func (UnimplementedOrgMagServiceServer) DeleteUser(context.Context, *C2S_DeleteUserReq) (*C2S_DeleteUserRpn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedOrgMagServiceServer) QueryUsers(context.Context, *C2S_UserQueryReq) (*C2S_UserQueryRpn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUsers not implemented")
}
func (UnimplementedOrgMagServiceServer) GetCompanyById(context.Context, *C2S_GetCompanyByIdReq) (*CompanyMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyById not implemented")
}
func (UnimplementedOrgMagServiceServer) AddExistUser(context.Context, *C2S_AddExistUserReq) (*C2S_AddExistUserRpn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExistUser not implemented")
}
func (UnimplementedOrgMagServiceServer) mustEmbedUnimplementedOrgMagServiceServer() {}

// UnsafeOrgMagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgMagServiceServer will
// result in compilation errors.
type UnsafeOrgMagServiceServer interface {
	mustEmbedUnimplementedOrgMagServiceServer()
}

func RegisterOrgMagServiceServer(s grpc.ServiceRegistrar, srv OrgMagServiceServer) {
	s.RegisterService(&OrgMagService_ServiceDesc, srv)
}

func _OrgMagService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).CreateCompany(ctx, req.(*CompanyCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).DeleteCompany(ctx, req.(*CompanyDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).UpdateCompany(ctx, req.(*CompanyUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_QueryCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_QueryCompanyReqT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).QueryCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/QueryCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).QueryCompany(ctx, req.(*C2S_QueryCompanyReqT))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_GetCompanyListById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyListByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).GetCompanyListById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/GetCompanyListById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).GetCompanyListById(ctx, req.(*GetCompanyListByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_GetCompanyListByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_GetCompanyListByUidReqT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).GetCompanyListByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/GetCompanyListByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).GetCompanyListByUid(ctx, req.(*C2S_GetCompanyListByUidReqT))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_ListUserByOpenId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_ListUserByOpenIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).ListUserByOpenId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/listUserByOpenId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).ListUserByOpenId(ctx, req.(*C2S_ListUserByOpenIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_ListUserByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_ListUserByUIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).ListUserByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/listUserByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).ListUserByUid(ctx, req.(*C2S_ListUserByUIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_AddCompanyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_AddCompanyUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).AddCompanyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/addCompanyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).AddCompanyUser(ctx, req.(*C2S_AddCompanyUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_AddCompanyOutsideUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_AddCompanyOutsideUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).AddCompanyOutsideUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/addCompanyOutsideUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).AddCompanyOutsideUser(ctx, req.(*C2S_AddCompanyOutsideUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/deleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).DeleteUser(ctx, req.(*C2S_DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_QueryUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_UserQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).QueryUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/queryUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).QueryUsers(ctx, req.(*C2S_UserQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_GetCompanyById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_GetCompanyByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).GetCompanyById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/getCompanyById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).GetCompanyById(ctx, req.(*C2S_GetCompanyByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgMagService_AddExistUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_AddExistUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgMagServiceServer).AddExistUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structure.OrgMagService/addExistUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgMagServiceServer).AddExistUser(ctx, req.(*C2S_AddExistUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrgMagService_ServiceDesc is the grpc.ServiceDesc for OrgMagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrgMagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "structure.OrgMagService",
	HandlerType: (*OrgMagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _OrgMagService_CreateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _OrgMagService_DeleteCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _OrgMagService_UpdateCompany_Handler,
		},
		{
			MethodName: "QueryCompany",
			Handler:    _OrgMagService_QueryCompany_Handler,
		},
		{
			MethodName: "GetCompanyListById",
			Handler:    _OrgMagService_GetCompanyListById_Handler,
		},
		{
			MethodName: "GetCompanyListByUid",
			Handler:    _OrgMagService_GetCompanyListByUid_Handler,
		},
		{
			MethodName: "listUserByOpenId",
			Handler:    _OrgMagService_ListUserByOpenId_Handler,
		},
		{
			MethodName: "listUserByUid",
			Handler:    _OrgMagService_ListUserByUid_Handler,
		},
		{
			MethodName: "addCompanyUser",
			Handler:    _OrgMagService_AddCompanyUser_Handler,
		},
		{
			MethodName: "addCompanyOutsideUser",
			Handler:    _OrgMagService_AddCompanyOutsideUser_Handler,
		},
		{
			MethodName: "deleteUser",
			Handler:    _OrgMagService_DeleteUser_Handler,
		},
		{
			MethodName: "queryUsers",
			Handler:    _OrgMagService_QueryUsers_Handler,
		},
		{
			MethodName: "getCompanyById",
			Handler:    _OrgMagService_GetCompanyById_Handler,
		},
		{
			MethodName: "addExistUser",
			Handler:    _OrgMagService_AddExistUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OrgMagService.proto",
}

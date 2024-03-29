// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: AssetLibraryService.proto

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

// AssetLibrarySerivceClient is the client API for AssetLibrarySerivce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetLibrarySerivceClient interface {
	// 查询资产列表
	QueryAssetPage(ctx context.Context, in *QueryAssetPageRequest, opts ...grpc.CallOption) (*QueryAssetPageResponse, error)
	// 添加资产
	AddAsset(ctx context.Context, in *AddAssetRequest, opts ...grpc.CallOption) (*AddAssetResponse, error)
	// 上传资产
	AssetUploaded(ctx context.Context, in *AssetUploadedRequest, opts ...grpc.CallOption) (*AssetUploadedResponse, error)
	// 修改资产
	UpdateAsset(ctx context.Context, in *UpdateAssetRequest, opts ...grpc.CallOption) (*UpdateAssetResponse, error)
	// 删除资产
	DeleteAsset(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*DeleteAssetResponse, error)
	// 获取资产详情
	GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*GetAssetResponse, error)
	// 更新缩略图
	UpdateAssetThumbnail(ctx context.Context, in *UpdateAssetThumbnailRequest, opts ...grpc.CallOption) (*UpdateAssetThumbnailResponse, error)
	// 添加分类
	AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*AddClassResponse, error)
	// 修改分类
	UpdateClass(ctx context.Context, in *UpdateClassRequest, opts ...grpc.CallOption) (*UpdateClassResponse, error)
	// 删除分类
	DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...grpc.CallOption) (*DeleteClassResponse, error)
	// 查询分类
	QueryClass(ctx context.Context, in *QueryClassRequest, opts ...grpc.CallOption) (*QueryClassResponse, error)
	// 查询分类详情
	GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error)
	// 添加资产标签
	AddTag(ctx context.Context, in *AddTagRequest, opts ...grpc.CallOption) (*AddTagResponse, error)
	// 修改资产标签
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagResponse, error)
	// 删除资产标签
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error)
	// 查询资产标签
	QueryTag(ctx context.Context, in *QueryTagRequest, opts ...grpc.CallOption) (*QueryTagResponse, error)
	// 查询资产标签列表
	QueryTagList(ctx context.Context, in *QueryTagListRequest, opts ...grpc.CallOption) (*QueryTagListResponse, error)
	// 添加实例范围
	AddInstanceScopeId(ctx context.Context, in *AddInstanceScopeIdRequest, opts ...grpc.CallOption) (*AddInstanceScopeIdResponse, error)
	// 删除实例范围
	DeleteInstanceScopeId(ctx context.Context, in *DeleteInstanceScopeIdRequest, opts ...grpc.CallOption) (*DeleteInstanceScopeIdResponse, error)
	// 获取实例范围
	GetInstanceScopeId(ctx context.Context, in *GetInstanceScopeIdRequest, opts ...grpc.CallOption) (*GetInstanceScopeIdResponse, error)
	// 查询实例范围
	QueryInstanceScopeId(ctx context.Context, in *QueryInstanceScopeIdRequest, opts ...grpc.CallOption) (*QueryInstanceScopeIdResponse, error)
	// 更新实例范围
	UpdateInstanceScopeId(ctx context.Context, in *UpdateInstanceScopeIdRequest, opts ...grpc.CallOption) (*UpdateInstanceScopeIdResponse, error)
}

type assetLibrarySerivceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetLibrarySerivceClient(cc grpc.ClientConnInterface) AssetLibrarySerivceClient {
	return &assetLibrarySerivceClient{cc}
}

func (c *assetLibrarySerivceClient) QueryAssetPage(ctx context.Context, in *QueryAssetPageRequest, opts ...grpc.CallOption) (*QueryAssetPageResponse, error) {
	out := new(QueryAssetPageResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/QueryAssetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) AddAsset(ctx context.Context, in *AddAssetRequest, opts ...grpc.CallOption) (*AddAssetResponse, error) {
	out := new(AddAssetResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/AddAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) AssetUploaded(ctx context.Context, in *AssetUploadedRequest, opts ...grpc.CallOption) (*AssetUploadedResponse, error) {
	out := new(AssetUploadedResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/AssetUploaded", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) UpdateAsset(ctx context.Context, in *UpdateAssetRequest, opts ...grpc.CallOption) (*UpdateAssetResponse, error) {
	out := new(UpdateAssetResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/UpdateAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) DeleteAsset(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*DeleteAssetResponse, error) {
	out := new(DeleteAssetResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/DeleteAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*GetAssetResponse, error) {
	out := new(GetAssetResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/GetAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) UpdateAssetThumbnail(ctx context.Context, in *UpdateAssetThumbnailRequest, opts ...grpc.CallOption) (*UpdateAssetThumbnailResponse, error) {
	out := new(UpdateAssetThumbnailResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/UpdateAssetThumbnail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*AddClassResponse, error) {
	out := new(AddClassResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/AddClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) UpdateClass(ctx context.Context, in *UpdateClassRequest, opts ...grpc.CallOption) (*UpdateClassResponse, error) {
	out := new(UpdateClassResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/UpdateClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...grpc.CallOption) (*DeleteClassResponse, error) {
	out := new(DeleteClassResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/DeleteClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) QueryClass(ctx context.Context, in *QueryClassRequest, opts ...grpc.CallOption) (*QueryClassResponse, error) {
	out := new(QueryClassResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/QueryClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error) {
	out := new(GetClassResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/GetClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) AddTag(ctx context.Context, in *AddTagRequest, opts ...grpc.CallOption) (*AddTagResponse, error) {
	out := new(AddTagResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagResponse, error) {
	out := new(UpdateTagResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/UpdateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error) {
	out := new(DeleteTagResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) QueryTag(ctx context.Context, in *QueryTagRequest, opts ...grpc.CallOption) (*QueryTagResponse, error) {
	out := new(QueryTagResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/QueryTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) QueryTagList(ctx context.Context, in *QueryTagListRequest, opts ...grpc.CallOption) (*QueryTagListResponse, error) {
	out := new(QueryTagListResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/QueryTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) AddInstanceScopeId(ctx context.Context, in *AddInstanceScopeIdRequest, opts ...grpc.CallOption) (*AddInstanceScopeIdResponse, error) {
	out := new(AddInstanceScopeIdResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/AddInstanceScopeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) DeleteInstanceScopeId(ctx context.Context, in *DeleteInstanceScopeIdRequest, opts ...grpc.CallOption) (*DeleteInstanceScopeIdResponse, error) {
	out := new(DeleteInstanceScopeIdResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/DeleteInstanceScopeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) GetInstanceScopeId(ctx context.Context, in *GetInstanceScopeIdRequest, opts ...grpc.CallOption) (*GetInstanceScopeIdResponse, error) {
	out := new(GetInstanceScopeIdResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/GetInstanceScopeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) QueryInstanceScopeId(ctx context.Context, in *QueryInstanceScopeIdRequest, opts ...grpc.CallOption) (*QueryInstanceScopeIdResponse, error) {
	out := new(QueryInstanceScopeIdResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/QueryInstanceScopeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetLibrarySerivceClient) UpdateInstanceScopeId(ctx context.Context, in *UpdateInstanceScopeIdRequest, opts ...grpc.CallOption) (*UpdateInstanceScopeIdResponse, error) {
	out := new(UpdateInstanceScopeIdResponse)
	err := c.cc.Invoke(ctx, "/com.zixel.file.backend.AssetLibrarySerivce/UpdateInstanceScopeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetLibrarySerivceServer is the server API for AssetLibrarySerivce service.
// All implementations must embed UnimplementedAssetLibrarySerivceServer
// for forward compatibility
type AssetLibrarySerivceServer interface {
	// 查询资产列表
	QueryAssetPage(context.Context, *QueryAssetPageRequest) (*QueryAssetPageResponse, error)
	// 添加资产
	AddAsset(context.Context, *AddAssetRequest) (*AddAssetResponse, error)
	// 上传资产
	AssetUploaded(context.Context, *AssetUploadedRequest) (*AssetUploadedResponse, error)
	// 修改资产
	UpdateAsset(context.Context, *UpdateAssetRequest) (*UpdateAssetResponse, error)
	// 删除资产
	DeleteAsset(context.Context, *DeleteAssetRequest) (*DeleteAssetResponse, error)
	// 获取资产详情
	GetAsset(context.Context, *GetAssetRequest) (*GetAssetResponse, error)
	// 更新缩略图
	UpdateAssetThumbnail(context.Context, *UpdateAssetThumbnailRequest) (*UpdateAssetThumbnailResponse, error)
	// 添加分类
	AddClass(context.Context, *AddClassRequest) (*AddClassResponse, error)
	// 修改分类
	UpdateClass(context.Context, *UpdateClassRequest) (*UpdateClassResponse, error)
	// 删除分类
	DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error)
	// 查询分类
	QueryClass(context.Context, *QueryClassRequest) (*QueryClassResponse, error)
	// 查询分类详情
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)
	// 添加资产标签
	AddTag(context.Context, *AddTagRequest) (*AddTagResponse, error)
	// 修改资产标签
	UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagResponse, error)
	// 删除资产标签
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagResponse, error)
	// 查询资产标签
	QueryTag(context.Context, *QueryTagRequest) (*QueryTagResponse, error)
	// 查询资产标签列表
	QueryTagList(context.Context, *QueryTagListRequest) (*QueryTagListResponse, error)
	// 添加实例范围
	AddInstanceScopeId(context.Context, *AddInstanceScopeIdRequest) (*AddInstanceScopeIdResponse, error)
	// 删除实例范围
	DeleteInstanceScopeId(context.Context, *DeleteInstanceScopeIdRequest) (*DeleteInstanceScopeIdResponse, error)
	// 获取实例范围
	GetInstanceScopeId(context.Context, *GetInstanceScopeIdRequest) (*GetInstanceScopeIdResponse, error)
	// 查询实例范围
	QueryInstanceScopeId(context.Context, *QueryInstanceScopeIdRequest) (*QueryInstanceScopeIdResponse, error)
	// 更新实例范围
	UpdateInstanceScopeId(context.Context, *UpdateInstanceScopeIdRequest) (*UpdateInstanceScopeIdResponse, error)
	mustEmbedUnimplementedAssetLibrarySerivceServer()
}

// UnimplementedAssetLibrarySerivceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetLibrarySerivceServer struct {
}

func (UnimplementedAssetLibrarySerivceServer) QueryAssetPage(context.Context, *QueryAssetPageRequest) (*QueryAssetPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAssetPage not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) AddAsset(context.Context, *AddAssetRequest) (*AddAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAsset not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) AssetUploaded(context.Context, *AssetUploadedRequest) (*AssetUploadedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetUploaded not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) UpdateAsset(context.Context, *UpdateAssetRequest) (*UpdateAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAsset not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) DeleteAsset(context.Context, *DeleteAssetRequest) (*DeleteAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAsset not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) GetAsset(context.Context, *GetAssetRequest) (*GetAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAsset not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) UpdateAssetThumbnail(context.Context, *UpdateAssetThumbnailRequest) (*UpdateAssetThumbnailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAssetThumbnail not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) AddClass(context.Context, *AddClassRequest) (*AddClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClass not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) UpdateClass(context.Context, *UpdateClassRequest) (*UpdateClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClass not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClass not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) QueryClass(context.Context, *QueryClassRequest) (*QueryClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryClass not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClass not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) AddTag(context.Context, *AddTagRequest) (*AddTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) QueryTag(context.Context, *QueryTagRequest) (*QueryTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTag not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) QueryTagList(context.Context, *QueryTagListRequest) (*QueryTagListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTagList not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) AddInstanceScopeId(context.Context, *AddInstanceScopeIdRequest) (*AddInstanceScopeIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInstanceScopeId not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) DeleteInstanceScopeId(context.Context, *DeleteInstanceScopeIdRequest) (*DeleteInstanceScopeIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstanceScopeId not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) GetInstanceScopeId(context.Context, *GetInstanceScopeIdRequest) (*GetInstanceScopeIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstanceScopeId not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) QueryInstanceScopeId(context.Context, *QueryInstanceScopeIdRequest) (*QueryInstanceScopeIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryInstanceScopeId not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) UpdateInstanceScopeId(context.Context, *UpdateInstanceScopeIdRequest) (*UpdateInstanceScopeIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstanceScopeId not implemented")
}
func (UnimplementedAssetLibrarySerivceServer) mustEmbedUnimplementedAssetLibrarySerivceServer() {}

// UnsafeAssetLibrarySerivceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetLibrarySerivceServer will
// result in compilation errors.
type UnsafeAssetLibrarySerivceServer interface {
	mustEmbedUnimplementedAssetLibrarySerivceServer()
}

func RegisterAssetLibrarySerivceServer(s grpc.ServiceRegistrar, srv AssetLibrarySerivceServer) {
	s.RegisterService(&AssetLibrarySerivce_ServiceDesc, srv)
}

func _AssetLibrarySerivce_QueryAssetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAssetPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).QueryAssetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/QueryAssetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).QueryAssetPage(ctx, req.(*QueryAssetPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_AddAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).AddAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/AddAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).AddAsset(ctx, req.(*AddAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_AssetUploaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetUploadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).AssetUploaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/AssetUploaded",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).AssetUploaded(ctx, req.(*AssetUploadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_UpdateAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).UpdateAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/UpdateAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).UpdateAsset(ctx, req.(*UpdateAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_DeleteAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).DeleteAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/DeleteAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).DeleteAsset(ctx, req.(*DeleteAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_GetAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).GetAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/GetAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).GetAsset(ctx, req.(*GetAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_UpdateAssetThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAssetThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).UpdateAssetThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/UpdateAssetThumbnail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).UpdateAssetThumbnail(ctx, req.(*UpdateAssetThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_AddClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).AddClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/AddClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).AddClass(ctx, req.(*AddClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_UpdateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).UpdateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/UpdateClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).UpdateClass(ctx, req.(*UpdateClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_DeleteClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).DeleteClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/DeleteClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).DeleteClass(ctx, req.(*DeleteClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_QueryClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).QueryClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/QueryClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).QueryClass(ctx, req.(*QueryClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_GetClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).GetClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/GetClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).GetClass(ctx, req.(*GetClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).AddTag(ctx, req.(*AddTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/UpdateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/DeleteTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_QueryTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).QueryTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/QueryTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).QueryTag(ctx, req.(*QueryTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_QueryTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).QueryTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/QueryTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).QueryTagList(ctx, req.(*QueryTagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_AddInstanceScopeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInstanceScopeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).AddInstanceScopeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/AddInstanceScopeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).AddInstanceScopeId(ctx, req.(*AddInstanceScopeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_DeleteInstanceScopeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceScopeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).DeleteInstanceScopeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/DeleteInstanceScopeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).DeleteInstanceScopeId(ctx, req.(*DeleteInstanceScopeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_GetInstanceScopeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceScopeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).GetInstanceScopeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/GetInstanceScopeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).GetInstanceScopeId(ctx, req.(*GetInstanceScopeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_QueryInstanceScopeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInstanceScopeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).QueryInstanceScopeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/QueryInstanceScopeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).QueryInstanceScopeId(ctx, req.(*QueryInstanceScopeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetLibrarySerivce_UpdateInstanceScopeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceScopeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetLibrarySerivceServer).UpdateInstanceScopeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.zixel.file.backend.AssetLibrarySerivce/UpdateInstanceScopeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetLibrarySerivceServer).UpdateInstanceScopeId(ctx, req.(*UpdateInstanceScopeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetLibrarySerivce_ServiceDesc is the grpc.ServiceDesc for AssetLibrarySerivce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetLibrarySerivce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.zixel.file.backend.AssetLibrarySerivce",
	HandlerType: (*AssetLibrarySerivceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryAssetPage",
			Handler:    _AssetLibrarySerivce_QueryAssetPage_Handler,
		},
		{
			MethodName: "AddAsset",
			Handler:    _AssetLibrarySerivce_AddAsset_Handler,
		},
		{
			MethodName: "AssetUploaded",
			Handler:    _AssetLibrarySerivce_AssetUploaded_Handler,
		},
		{
			MethodName: "UpdateAsset",
			Handler:    _AssetLibrarySerivce_UpdateAsset_Handler,
		},
		{
			MethodName: "DeleteAsset",
			Handler:    _AssetLibrarySerivce_DeleteAsset_Handler,
		},
		{
			MethodName: "GetAsset",
			Handler:    _AssetLibrarySerivce_GetAsset_Handler,
		},
		{
			MethodName: "UpdateAssetThumbnail",
			Handler:    _AssetLibrarySerivce_UpdateAssetThumbnail_Handler,
		},
		{
			MethodName: "AddClass",
			Handler:    _AssetLibrarySerivce_AddClass_Handler,
		},
		{
			MethodName: "UpdateClass",
			Handler:    _AssetLibrarySerivce_UpdateClass_Handler,
		},
		{
			MethodName: "DeleteClass",
			Handler:    _AssetLibrarySerivce_DeleteClass_Handler,
		},
		{
			MethodName: "QueryClass",
			Handler:    _AssetLibrarySerivce_QueryClass_Handler,
		},
		{
			MethodName: "GetClass",
			Handler:    _AssetLibrarySerivce_GetClass_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _AssetLibrarySerivce_AddTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _AssetLibrarySerivce_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _AssetLibrarySerivce_DeleteTag_Handler,
		},
		{
			MethodName: "QueryTag",
			Handler:    _AssetLibrarySerivce_QueryTag_Handler,
		},
		{
			MethodName: "QueryTagList",
			Handler:    _AssetLibrarySerivce_QueryTagList_Handler,
		},
		{
			MethodName: "AddInstanceScopeId",
			Handler:    _AssetLibrarySerivce_AddInstanceScopeId_Handler,
		},
		{
			MethodName: "DeleteInstanceScopeId",
			Handler:    _AssetLibrarySerivce_DeleteInstanceScopeId_Handler,
		},
		{
			MethodName: "GetInstanceScopeId",
			Handler:    _AssetLibrarySerivce_GetInstanceScopeId_Handler,
		},
		{
			MethodName: "QueryInstanceScopeId",
			Handler:    _AssetLibrarySerivce_QueryInstanceScopeId_Handler,
		},
		{
			MethodName: "UpdateInstanceScopeId",
			Handler:    _AssetLibrarySerivce_UpdateInstanceScopeId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetLibraryService.proto",
}

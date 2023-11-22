// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: core.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreServiceClient interface {
	// creation
	CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*GenericCreateResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*GenericCreateResponse, error)
	CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*GenericCreateResponse, error)
	// editing
	EditField(ctx context.Context, in *EditFieldRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error)
	EditCategory(ctx context.Context, in *EditCategoryRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error)
	EditEntity(ctx context.Context, in *EditEntityRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error)
	// deleting
	DeleteField(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error)
	DeleteEntity(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error)
	DeleteCategory(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error)
	// fetching
	NewCategorySchema(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Schema, error)
	NewFieldSchema(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Schema, error)
	GetCategorySchema(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*CategorySchemaResponse, error)
	GetCategoryData(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GetCategoryDataResponse, error)
	GetFieldData(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GetFieldDataResponse, error)
	GetEntityData(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GetEntityDataResponse, error)
	// listing
	ListFields(ctx context.Context, in *PaginatedFieldFetchRequest, opts ...grpc.CallOption) (*PaginatedFieldsFetchResponse, error)
	ListCategories(ctx context.Context, in *PaginatedCategoriesFetchRequest, opts ...grpc.CallOption) (*PaginatedCategoriesFetchResponse, error)
	ListEntities(ctx context.Context, in *PaginatedEntitiesFetchRequest, opts ...grpc.CallOption) (*PaginatedEntititesFetchResponse, error)
	//misc
	FieldsInheritance(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*InheritanceResponse, error)
	CoreMiddleware(ctx context.Context, in *CoreMiddlewareRequest, opts ...grpc.CallOption) (*CoreMiddlewareResponse, error)
}

type coreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreServiceClient(cc grpc.ClientConnInterface) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*GenericCreateResponse, error) {
	out := new(GenericCreateResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/CreateField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*GenericCreateResponse, error) {
	out := new(GenericCreateResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*GenericCreateResponse, error) {
	out := new(GenericCreateResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) EditField(ctx context.Context, in *EditFieldRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error) {
	out := new(GenericEditDeleteResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/EditField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) EditCategory(ctx context.Context, in *EditCategoryRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error) {
	out := new(GenericEditDeleteResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/EditCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) EditEntity(ctx context.Context, in *EditEntityRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error) {
	out := new(GenericEditDeleteResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/EditEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteField(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error) {
	out := new(GenericEditDeleteResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/DeleteField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteEntity(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error) {
	out := new(GenericEditDeleteResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/DeleteEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) DeleteCategory(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GenericEditDeleteResponse, error) {
	out := new(GenericEditDeleteResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) NewCategorySchema(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, "/core.CoreService/NewCategorySchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) NewFieldSchema(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, "/core.CoreService/NewFieldSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetCategorySchema(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*CategorySchemaResponse, error) {
	out := new(CategorySchemaResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetCategorySchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetCategoryData(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GetCategoryDataResponse, error) {
	out := new(GetCategoryDataResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetCategoryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetFieldData(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GetFieldDataResponse, error) {
	out := new(GetFieldDataResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetFieldData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetEntityData(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*GetEntityDataResponse, error) {
	out := new(GetEntityDataResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/GetEntityData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ListFields(ctx context.Context, in *PaginatedFieldFetchRequest, opts ...grpc.CallOption) (*PaginatedFieldsFetchResponse, error) {
	out := new(PaginatedFieldsFetchResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/ListFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ListCategories(ctx context.Context, in *PaginatedCategoriesFetchRequest, opts ...grpc.CallOption) (*PaginatedCategoriesFetchResponse, error) {
	out := new(PaginatedCategoriesFetchResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/ListCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ListEntities(ctx context.Context, in *PaginatedEntitiesFetchRequest, opts ...grpc.CallOption) (*PaginatedEntititesFetchResponse, error) {
	out := new(PaginatedEntititesFetchResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/ListEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) FieldsInheritance(ctx context.Context, in *GenericFetchRequest, opts ...grpc.CallOption) (*InheritanceResponse, error) {
	out := new(InheritanceResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/FieldsInheritance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) CoreMiddleware(ctx context.Context, in *CoreMiddlewareRequest, opts ...grpc.CallOption) (*CoreMiddlewareResponse, error) {
	out := new(CoreMiddlewareResponse)
	err := c.cc.Invoke(ctx, "/core.CoreService/CoreMiddleware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
// All implementations must embed UnimplementedCoreServiceServer
// for forward compatibility
type CoreServiceServer interface {
	// creation
	CreateField(context.Context, *CreateFieldRequest) (*GenericCreateResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*GenericCreateResponse, error)
	CreateEntity(context.Context, *CreateEntityRequest) (*GenericCreateResponse, error)
	// editing
	EditField(context.Context, *EditFieldRequest) (*GenericEditDeleteResponse, error)
	EditCategory(context.Context, *EditCategoryRequest) (*GenericEditDeleteResponse, error)
	EditEntity(context.Context, *EditEntityRequest) (*GenericEditDeleteResponse, error)
	// deleting
	DeleteField(context.Context, *GenericFetchRequest) (*GenericEditDeleteResponse, error)
	DeleteEntity(context.Context, *GenericFetchRequest) (*GenericEditDeleteResponse, error)
	DeleteCategory(context.Context, *GenericFetchRequest) (*GenericEditDeleteResponse, error)
	// fetching
	NewCategorySchema(context.Context, *emptypb.Empty) (*Schema, error)
	NewFieldSchema(context.Context, *emptypb.Empty) (*Schema, error)
	GetCategorySchema(context.Context, *GenericFetchRequest) (*CategorySchemaResponse, error)
	GetCategoryData(context.Context, *GenericFetchRequest) (*GetCategoryDataResponse, error)
	GetFieldData(context.Context, *GenericFetchRequest) (*GetFieldDataResponse, error)
	GetEntityData(context.Context, *GenericFetchRequest) (*GetEntityDataResponse, error)
	// listing
	ListFields(context.Context, *PaginatedFieldFetchRequest) (*PaginatedFieldsFetchResponse, error)
	ListCategories(context.Context, *PaginatedCategoriesFetchRequest) (*PaginatedCategoriesFetchResponse, error)
	ListEntities(context.Context, *PaginatedEntitiesFetchRequest) (*PaginatedEntititesFetchResponse, error)
	//misc
	FieldsInheritance(context.Context, *GenericFetchRequest) (*InheritanceResponse, error)
	CoreMiddleware(context.Context, *CoreMiddlewareRequest) (*CoreMiddlewareResponse, error)
	mustEmbedUnimplementedCoreServiceServer()
}

// UnimplementedCoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (UnimplementedCoreServiceServer) CreateField(context.Context, *CreateFieldRequest) (*GenericCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateField not implemented")
}
func (UnimplementedCoreServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*GenericCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCoreServiceServer) CreateEntity(context.Context, *CreateEntityRequest) (*GenericCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntity not implemented")
}
func (UnimplementedCoreServiceServer) EditField(context.Context, *EditFieldRequest) (*GenericEditDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditField not implemented")
}
func (UnimplementedCoreServiceServer) EditCategory(context.Context, *EditCategoryRequest) (*GenericEditDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategory not implemented")
}
func (UnimplementedCoreServiceServer) EditEntity(context.Context, *EditEntityRequest) (*GenericEditDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEntity not implemented")
}
func (UnimplementedCoreServiceServer) DeleteField(context.Context, *GenericFetchRequest) (*GenericEditDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteField not implemented")
}
func (UnimplementedCoreServiceServer) DeleteEntity(context.Context, *GenericFetchRequest) (*GenericEditDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntity not implemented")
}
func (UnimplementedCoreServiceServer) DeleteCategory(context.Context, *GenericFetchRequest) (*GenericEditDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCoreServiceServer) NewCategorySchema(context.Context, *emptypb.Empty) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewCategorySchema not implemented")
}
func (UnimplementedCoreServiceServer) NewFieldSchema(context.Context, *emptypb.Empty) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewFieldSchema not implemented")
}
func (UnimplementedCoreServiceServer) GetCategorySchema(context.Context, *GenericFetchRequest) (*CategorySchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategorySchema not implemented")
}
func (UnimplementedCoreServiceServer) GetCategoryData(context.Context, *GenericFetchRequest) (*GetCategoryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryData not implemented")
}
func (UnimplementedCoreServiceServer) GetFieldData(context.Context, *GenericFetchRequest) (*GetFieldDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFieldData not implemented")
}
func (UnimplementedCoreServiceServer) GetEntityData(context.Context, *GenericFetchRequest) (*GetEntityDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntityData not implemented")
}
func (UnimplementedCoreServiceServer) ListFields(context.Context, *PaginatedFieldFetchRequest) (*PaginatedFieldsFetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFields not implemented")
}
func (UnimplementedCoreServiceServer) ListCategories(context.Context, *PaginatedCategoriesFetchRequest) (*PaginatedCategoriesFetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedCoreServiceServer) ListEntities(context.Context, *PaginatedEntitiesFetchRequest) (*PaginatedEntititesFetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntities not implemented")
}
func (UnimplementedCoreServiceServer) FieldsInheritance(context.Context, *GenericFetchRequest) (*InheritanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FieldsInheritance not implemented")
}
func (UnimplementedCoreServiceServer) CoreMiddleware(context.Context, *CoreMiddlewareRequest) (*CoreMiddlewareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoreMiddleware not implemented")
}
func (UnimplementedCoreServiceServer) mustEmbedUnimplementedCoreServiceServer() {}

// UnsafeCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServiceServer will
// result in compilation errors.
type UnsafeCoreServiceServer interface {
	mustEmbedUnimplementedCoreServiceServer()
}

func RegisterCoreServiceServer(s grpc.ServiceRegistrar, srv CoreServiceServer) {
	s.RegisterService(&CoreService_ServiceDesc, srv)
}

func _CoreService_CreateField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/CreateField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateField(ctx, req.(*CreateFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CreateEntity(ctx, req.(*CreateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_EditField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).EditField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/EditField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).EditField(ctx, req.(*EditFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_EditCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).EditCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/EditCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).EditCategory(ctx, req.(*EditCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_EditEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).EditEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/EditEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).EditEntity(ctx, req.(*EditEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/DeleteField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteField(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/DeleteEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteEntity(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).DeleteCategory(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_NewCategorySchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).NewCategorySchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/NewCategorySchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).NewCategorySchema(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_NewFieldSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).NewFieldSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/NewFieldSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).NewFieldSchema(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetCategorySchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetCategorySchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetCategorySchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetCategorySchema(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetCategoryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetCategoryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetCategoryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetCategoryData(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetFieldData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetFieldData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetFieldData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetFieldData(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetEntityData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetEntityData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/GetEntityData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetEntityData(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ListFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginatedFieldFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/ListFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListFields(ctx, req.(*PaginatedFieldFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginatedCategoriesFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/ListCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListCategories(ctx, req.(*PaginatedCategoriesFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ListEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginatedEntitiesFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ListEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/ListEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ListEntities(ctx, req.(*PaginatedEntitiesFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_FieldsInheritance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).FieldsInheritance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/FieldsInheritance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).FieldsInheritance(ctx, req.(*GenericFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_CoreMiddleware_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreMiddlewareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).CoreMiddleware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreService/CoreMiddleware",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).CoreMiddleware(ctx, req.(*CoreMiddlewareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoreService_ServiceDesc is the grpc.ServiceDesc for CoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateField",
			Handler:    _CoreService_CreateField_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CoreService_CreateCategory_Handler,
		},
		{
			MethodName: "CreateEntity",
			Handler:    _CoreService_CreateEntity_Handler,
		},
		{
			MethodName: "EditField",
			Handler:    _CoreService_EditField_Handler,
		},
		{
			MethodName: "EditCategory",
			Handler:    _CoreService_EditCategory_Handler,
		},
		{
			MethodName: "EditEntity",
			Handler:    _CoreService_EditEntity_Handler,
		},
		{
			MethodName: "DeleteField",
			Handler:    _CoreService_DeleteField_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _CoreService_DeleteEntity_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CoreService_DeleteCategory_Handler,
		},
		{
			MethodName: "NewCategorySchema",
			Handler:    _CoreService_NewCategorySchema_Handler,
		},
		{
			MethodName: "NewFieldSchema",
			Handler:    _CoreService_NewFieldSchema_Handler,
		},
		{
			MethodName: "GetCategorySchema",
			Handler:    _CoreService_GetCategorySchema_Handler,
		},
		{
			MethodName: "GetCategoryData",
			Handler:    _CoreService_GetCategoryData_Handler,
		},
		{
			MethodName: "GetFieldData",
			Handler:    _CoreService_GetFieldData_Handler,
		},
		{
			MethodName: "GetEntityData",
			Handler:    _CoreService_GetEntityData_Handler,
		},
		{
			MethodName: "ListFields",
			Handler:    _CoreService_ListFields_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _CoreService_ListCategories_Handler,
		},
		{
			MethodName: "ListEntities",
			Handler:    _CoreService_ListEntities_Handler,
		},
		{
			MethodName: "FieldsInheritance",
			Handler:    _CoreService_FieldsInheritance_Handler,
		},
		{
			MethodName: "CoreMiddleware",
			Handler:    _CoreService_CoreMiddleware_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

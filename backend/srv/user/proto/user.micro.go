// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Create(ctx context.Context, in *UserCreateRequest, opts ...client.CallOption) (*UserCreateResponse, error)
	Query(ctx context.Context, in *UserQueryRequest, opts ...client.CallOption) (*UserInfo, error)
	Find(ctx context.Context, in *UserFindRequest, opts ...client.CallOption) (*UserFindResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *UserCreateRequest, opts ...client.CallOption) (*UserCreateResponse, error) {
	req := c.c.NewRequest(c.name, "User.Create", in)
	out := new(UserCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Query(ctx context.Context, in *UserQueryRequest, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "User.Query", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Find(ctx context.Context, in *UserFindRequest, opts ...client.CallOption) (*UserFindResponse, error) {
	req := c.c.NewRequest(c.name, "User.Find", in)
	out := new(UserFindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Create(context.Context, *UserCreateRequest, *UserCreateResponse) error
	Query(context.Context, *UserQueryRequest, *UserInfo) error
	Find(context.Context, *UserFindRequest, *UserFindResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Create(ctx context.Context, in *UserCreateRequest, out *UserCreateResponse) error
		Query(ctx context.Context, in *UserQueryRequest, out *UserInfo) error
		Find(ctx context.Context, in *UserFindRequest, out *UserFindResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Create(ctx context.Context, in *UserCreateRequest, out *UserCreateResponse) error {
	return h.UserHandler.Create(ctx, in, out)
}

func (h *userHandler) Query(ctx context.Context, in *UserQueryRequest, out *UserInfo) error {
	return h.UserHandler.Query(ctx, in, out)
}

func (h *userHandler) Find(ctx context.Context, in *UserFindRequest, out *UserFindResponse) error {
	return h.UserHandler.Find(ctx, in, out)
}

// Client API for AdminUser service

type AdminUserService interface {
	Create(ctx context.Context, in *AdminUserRequest, opts ...client.CallOption) (*AdminUserResponse, error)
	Find(ctx context.Context, in *AdminUserRequest, opts ...client.CallOption) (*AdminUserResponse, error)
}

type adminUserService struct {
	c    client.Client
	name string
}

func NewAdminUserService(name string, c client.Client) AdminUserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "adminuser"
	}
	return &adminUserService{
		c:    c,
		name: name,
	}
}

func (c *adminUserService) Create(ctx context.Context, in *AdminUserRequest, opts ...client.CallOption) (*AdminUserResponse, error) {
	req := c.c.NewRequest(c.name, "AdminUser.Create", in)
	out := new(AdminUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminUserService) Find(ctx context.Context, in *AdminUserRequest, opts ...client.CallOption) (*AdminUserResponse, error) {
	req := c.c.NewRequest(c.name, "AdminUser.Find", in)
	out := new(AdminUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdminUser service

type AdminUserHandler interface {
	Create(context.Context, *AdminUserRequest, *AdminUserResponse) error
	Find(context.Context, *AdminUserRequest, *AdminUserResponse) error
}

func RegisterAdminUserHandler(s server.Server, hdlr AdminUserHandler, opts ...server.HandlerOption) error {
	type adminUser interface {
		Create(ctx context.Context, in *AdminUserRequest, out *AdminUserResponse) error
		Find(ctx context.Context, in *AdminUserRequest, out *AdminUserResponse) error
	}
	type AdminUser struct {
		adminUser
	}
	h := &adminUserHandler{hdlr}
	return s.Handle(s.NewHandler(&AdminUser{h}, opts...))
}

type adminUserHandler struct {
	AdminUserHandler
}

func (h *adminUserHandler) Create(ctx context.Context, in *AdminUserRequest, out *AdminUserResponse) error {
	return h.AdminUserHandler.Create(ctx, in, out)
}

func (h *adminUserHandler) Find(ctx context.Context, in *AdminUserRequest, out *AdminUserResponse) error {
	return h.AdminUserHandler.Find(ctx, in, out)
}
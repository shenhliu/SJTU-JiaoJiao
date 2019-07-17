package mock

import (
	"context"
	"errors"
	"github.com/micro/go-micro/client"
	sellinfo "jiaojiao/srv/sellinfo/proto"
)

type mockSellInfoSrv struct{}
type mockContentSrv struct{}

func (a *mockSellInfoSrv) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, opts ...client.CallOption) (*sellinfo.SellInfoCreateResponse, error) {
	var rsp sellinfo.SellInfoCreateResponse
	return &rsp, nil
}

func (a *mockSellInfoSrv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, opts ...client.CallOption) (*sellinfo.SellInfoMsg, error) {
	var rsp sellinfo.SellInfoMsg
	if req.SellInfoId != 0 {
		if req.SellInfoId == 1000 {
			rsp.SellInfoId = 1000
			rsp.GoodName = "good"
			rsp.ValidTime = 1234567890
			rsp.Description = "very good!"
			rsp.ContentId = "123456789abc123456789abc"
			rsp.UserId = 1000
		} else if req.SellInfoId == 2000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

func (a *mockSellInfoSrv) Find(ctx context.Context, req *sellinfo.SellInfoFindRequest, opts ...client.CallOption) (*sellinfo.SellInfoFindResponse, error) {
	var rsp sellinfo.SellInfoFindResponse
	return &rsp, nil
}

func (a *mockContentSrv) Create(ctx context.Context, req *sellinfo.ContentCreateRequest, opts ...client.CallOption) (*sellinfo.ContentCreateResponse, error) {
	var rsp sellinfo.ContentCreateResponse
	return &rsp, nil
}

func (a *mockContentSrv) Delete(ctx context.Context, req *sellinfo.ContentDeleteRequest, opts ...client.CallOption) (*sellinfo.ContentDeleteResponse, error) {
	var rsp sellinfo.ContentDeleteResponse
	if req.ContentId == "" || req.ContentToken == "" {
		rsp.Status = sellinfo.ContentDeleteResponse_INVALID_PARAM
	} else {
		if req.ContentId == "1000" {
			if req.ContentToken == "valid_token" {
				rsp.Status = sellinfo.ContentDeleteResponse_SUCCESS
			} else {
				rsp.Status = sellinfo.ContentDeleteResponse_INVALID_TOKEN
			}
		} else if req.ContentId == "2000" {
			return nil, errors.New("")
		} else {
			rsp.Status = sellinfo.ContentDeleteResponse_INVALID_TOKEN
		}
	}
	return &rsp, nil
}

func NewSellInfoService() sellinfo.SellInfoService {
	return new(mockSellInfoSrv)
}

func NewContentService() sellinfo.ContentService {
	return new(mockContentSrv)
}

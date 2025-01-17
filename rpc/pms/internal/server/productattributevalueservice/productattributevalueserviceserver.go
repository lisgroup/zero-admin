// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"zero-admin/rpc/pms/internal/logic/productattributevalueservice"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"
)

type ProductAttributeValueServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductAttributeValueServiceServer
}

func NewProductAttributeValueServiceServer(svcCtx *svc.ServiceContext) *ProductAttributeValueServiceServer {
	return &ProductAttributeValueServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductAttributeValueServiceServer) ProductAttributeValueAdd(ctx context.Context, in *pmsclient.ProductAttributeValueAddReq) (*pmsclient.ProductAttributeValueAddResp, error) {
	l := productattributevalueservicelogic.NewProductAttributeValueAddLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueAdd(in)
}

func (s *ProductAttributeValueServiceServer) ProductAttributeValueList(ctx context.Context, in *pmsclient.ProductAttributeValueListReq) (*pmsclient.ProductAttributeValueListResp, error) {
	l := productattributevalueservicelogic.NewProductAttributeValueListLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueList(in)
}

func (s *ProductAttributeValueServiceServer) ProductAttributeValueUpdate(ctx context.Context, in *pmsclient.ProductAttributeValueUpdateReq) (*pmsclient.ProductAttributeValueUpdateResp, error) {
	l := productattributevalueservicelogic.NewProductAttributeValueUpdateLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueUpdate(in)
}

func (s *ProductAttributeValueServiceServer) ProductAttributeValueDelete(ctx context.Context, in *pmsclient.ProductAttributeValueDeleteReq) (*pmsclient.ProductAttributeValueDeleteResp, error) {
	l := productattributevalueservicelogic.NewProductAttributeValueDeleteLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueDelete(in)
}

// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"zero-admin/rpc/sms/internal/logic/homerecommendproductservice"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"
)

type HomeRecommendProductServiceServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedHomeRecommendProductServiceServer
}

func NewHomeRecommendProductServiceServer(svcCtx *svc.ServiceContext) *HomeRecommendProductServiceServer {
	return &HomeRecommendProductServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *HomeRecommendProductServiceServer) HomeRecommendProductAdd(ctx context.Context, in *smsclient.HomeRecommendProductAddReq) (*smsclient.HomeRecommendProductAddResp, error) {
	l := homerecommendproductservicelogic.NewHomeRecommendProductAddLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductAdd(in)
}

func (s *HomeRecommendProductServiceServer) HomeRecommendProductList(ctx context.Context, in *smsclient.HomeRecommendProductListReq) (*smsclient.HomeRecommendProductListResp, error) {
	l := homerecommendproductservicelogic.NewHomeRecommendProductListLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductList(in)
}

func (s *HomeRecommendProductServiceServer) HomeRecommendProductUpdate(ctx context.Context, in *smsclient.HomeRecommendProductUpdateReq) (*smsclient.HomeRecommendProductUpdateResp, error) {
	l := homerecommendproductservicelogic.NewHomeRecommendProductUpdateLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductUpdate(in)
}

func (s *HomeRecommendProductServiceServer) HomeRecommendProductDelete(ctx context.Context, in *smsclient.HomeRecommendProductDeleteReq) (*smsclient.HomeRecommendProductDeleteResp, error) {
	l := homerecommendproductservicelogic.NewHomeRecommendProductDeleteLogic(ctx, s.svcCtx)
	return l.HomeRecommendProductDelete(in)
}

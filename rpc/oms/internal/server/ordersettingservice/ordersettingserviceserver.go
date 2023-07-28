// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package server

import (
	"context"

	"zero-admin/rpc/oms/internal/logic/ordersettingservice"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"
)

type OrderSettingServiceServer struct {
	svcCtx *svc.ServiceContext
	omsclient.UnimplementedOrderSettingServiceServer
}

func NewOrderSettingServiceServer(svcCtx *svc.ServiceContext) *OrderSettingServiceServer {
	return &OrderSettingServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderSettingServiceServer) OrderSettingAdd(ctx context.Context, in *omsclient.OrderSettingAddReq) (*omsclient.OrderSettingAddResp, error) {
	l := ordersettingservicelogic.NewOrderSettingAddLogic(ctx, s.svcCtx)
	return l.OrderSettingAdd(in)
}

func (s *OrderSettingServiceServer) OrderSettingList(ctx context.Context, in *omsclient.OrderSettingListReq) (*omsclient.OrderSettingListResp, error) {
	l := ordersettingservicelogic.NewOrderSettingListLogic(ctx, s.svcCtx)
	return l.OrderSettingList(in)
}

func (s *OrderSettingServiceServer) OrderSettingUpdate(ctx context.Context, in *omsclient.OrderSettingUpdateReq) (*omsclient.OrderSettingUpdateResp, error) {
	l := ordersettingservicelogic.NewOrderSettingUpdateLogic(ctx, s.svcCtx)
	return l.OrderSettingUpdate(in)
}

func (s *OrderSettingServiceServer) OrderSettingDelete(ctx context.Context, in *omsclient.OrderSettingDeleteReq) (*omsclient.OrderSettingDeleteResp, error) {
	l := ordersettingservicelogic.NewOrderSettingDeleteLogic(ctx, s.svcCtx)
	return l.OrderSettingDelete(in)
}
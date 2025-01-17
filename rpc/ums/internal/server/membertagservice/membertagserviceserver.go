// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"zero-admin/rpc/ums/internal/logic/membertagservice"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"
)

type MemberTagServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberTagServiceServer
}

func NewMemberTagServiceServer(svcCtx *svc.ServiceContext) *MemberTagServiceServer {
	return &MemberTagServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MemberTagServiceServer) MemberTagAdd(ctx context.Context, in *umsclient.MemberTagAddReq) (*umsclient.MemberTagAddResp, error) {
	l := membertagservicelogic.NewMemberTagAddLogic(ctx, s.svcCtx)
	return l.MemberTagAdd(in)
}

func (s *MemberTagServiceServer) MemberTagList(ctx context.Context, in *umsclient.MemberTagListReq) (*umsclient.MemberTagListResp, error) {
	l := membertagservicelogic.NewMemberTagListLogic(ctx, s.svcCtx)
	return l.MemberTagList(in)
}

func (s *MemberTagServiceServer) MemberTagUpdate(ctx context.Context, in *umsclient.MemberTagUpdateReq) (*umsclient.MemberTagUpdateResp, error) {
	l := membertagservicelogic.NewMemberTagUpdateLogic(ctx, s.svcCtx)
	return l.MemberTagUpdate(in)
}

func (s *MemberTagServiceServer) MemberTagDelete(ctx context.Context, in *umsclient.MemberTagDeleteReq) (*umsclient.MemberTagDeleteResp, error) {
	l := membertagservicelogic.NewMemberTagDeleteLogic(ctx, s.svcCtx)
	return l.MemberTagDelete(in)
}

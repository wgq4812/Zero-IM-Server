// Code generated by goctl. DO NOT EDIT!
// Source: msg-gateway.proto

package rpcserver

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/internal/rpclogic"
	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"

	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/pb"
)

type OnlineMessageRelayServiceServer struct {
	svcCtx *rpcsvc.ServiceContext
	pb.UnimplementedOnlineMessageRelayServiceServer
}

func NewOnlineMessageRelayServiceServer(svcCtx *rpcsvc.ServiceContext) *OnlineMessageRelayServiceServer {
	return &OnlineMessageRelayServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OnlineMessageRelayServiceServer) OnlinePushMsg(ctx context.Context, in *pb.OnlinePushMsgReq) (*pb.OnlinePushMsgResp, error) {
	l := rpclogic.NewOnlinePushMsgLogic(ctx, s.svcCtx)
	return l.OnlinePushMsg(in)
}

func (s *OnlineMessageRelayServiceServer) GetUsersOnlineStatus(ctx context.Context, in *pb.GetUsersOnlineStatusReq) (*pb.GetUsersOnlineStatusResp, error) {
	l := rpclogic.NewGetUsersOnlineStatusLogic(ctx, s.svcCtx)
	return l.GetUsersOnlineStatus(in)
}

func (s *OnlineMessageRelayServiceServer) KickUserConns(ctx context.Context, in *pb.KickUserConnsReq) (*pb.KickUserConnsResp, error) {
	l := rpclogic.NewKickUserConnsLogic(ctx, s.svcCtx)
	return l.KickUserConns(in)
}

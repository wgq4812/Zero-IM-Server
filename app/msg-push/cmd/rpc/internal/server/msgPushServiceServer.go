// Code generated by goctl. DO NOT EDIT!
// Source: msgpush.proto

package server

import (
	"context"

	"github.com/Path-IM/Path-IM-Server/app/msg-push/cmd/rpc/internal/logic"
	"github.com/Path-IM/Path-IM-Server/app/msg-push/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server/app/msg-push/cmd/rpc/pb"
)

type MsgPushServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMsgPushServiceServer
}

func NewMsgPushServiceServer(svcCtx *svc.ServiceContext) *MsgPushServiceServer {
	return &MsgPushServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MsgPushServiceServer) PushMsg(ctx context.Context, in *pb.PushMsgReq) (*pb.PushMsgResp, error) {
	l := logic.NewPushMsgLogic(ctx, s.svcCtx)
	return l.OfflinePushMsg(in)
}

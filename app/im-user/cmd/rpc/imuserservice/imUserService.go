// Code generated by goctl. DO NOT EDIT!
// Source: im-user.proto

package imuserservice

import (
	"context"

	"github.com/Path-IM/Path-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResp                           = pb.CommonResp
	GetGroupMemberIDListFromCacheReq     = pb.GetGroupMemberIDListFromCacheReq
	GetGroupMemberIDListFromCacheResp    = pb.GetGroupMemberIDListFromCacheResp
	GetSingleConversationRecvMsgOptsReq  = pb.GetSingleConversationRecvMsgOptsReq
	GetSingleConversationRecvMsgOptsResp = pb.GetSingleConversationRecvMsgOptsResp
	GetUserListFromSuperGroupWithOptReq  = pb.GetUserListFromSuperGroupWithOptReq
	GetUserListFromSuperGroupWithOptResp = pb.GetUserListFromSuperGroupWithOptResp
	IfAInBBlacklistReq                   = pb.IfAInBBlacklistReq
	IfAInBBlacklistResp                  = pb.IfAInBBlacklistResp
	IfAInBFriendListReq                  = pb.IfAInBFriendListReq
	IfAInBFriendListResp                 = pb.IfAInBFriendListResp
	UserIDOpt                            = pb.UserIDOpt
	UserInfo                             = pb.UserInfo
	VerifyTokenReq                       = pb.VerifyTokenReq
	VerifyTokenResp                      = pb.VerifyTokenResp

	ImUserService interface {
		//  获取群组成员列表
		GetGroupMemberIDListFromCache(ctx context.Context, in *GetGroupMemberIDListFromCacheReq, opts ...grpc.CallOption) (*GetGroupMemberIDListFromCacheResp, error)
		//  判断用户A是否在B黑名单中
		IfAInBBlacklist(ctx context.Context, in *IfAInBBlacklistReq, opts ...grpc.CallOption) (*IfAInBBlacklistResp, error)
		//  判断用户A是否在B好友列表中
		IfAInBFriendList(ctx context.Context, in *IfAInBFriendListReq, opts ...grpc.CallOption) (*IfAInBFriendListResp, error)
		//  获取单聊会话的消息接收选项
		GetSingleConversationRecvMsgOpts(ctx context.Context, in *GetSingleConversationRecvMsgOptsReq, opts ...grpc.CallOption) (*GetSingleConversationRecvMsgOptsResp, error)
		//  获取超级群成员列表 通过消息接收选项
		GetUserListFromSuperGroupWithOpt(ctx context.Context, in *GetUserListFromSuperGroupWithOptReq, opts ...grpc.CallOption) (*GetUserListFromSuperGroupWithOptResp, error)
		//  检查token
		VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenResp, error)
	}

	defaultImUserService struct {
		cli zrpc.Client
	}
)

func NewImUserService(cli zrpc.Client) ImUserService {
	return &defaultImUserService{
		cli: cli,
	}
}

//  获取群组成员列表
func (m *defaultImUserService) GetGroupMemberIDListFromCache(ctx context.Context, in *GetGroupMemberIDListFromCacheReq, opts ...grpc.CallOption) (*GetGroupMemberIDListFromCacheResp, error) {
	client := pb.NewImUserServiceClient(m.cli.Conn())
	return client.GetGroupMemberIDListFromCache(ctx, in, opts...)
}

//  判断用户A是否在B黑名单中
func (m *defaultImUserService) IfAInBBlacklist(ctx context.Context, in *IfAInBBlacklistReq, opts ...grpc.CallOption) (*IfAInBBlacklistResp, error) {
	client := pb.NewImUserServiceClient(m.cli.Conn())
	return client.IfAInBBlacklist(ctx, in, opts...)
}

//  判断用户A是否在B好友列表中
func (m *defaultImUserService) IfAInBFriendList(ctx context.Context, in *IfAInBFriendListReq, opts ...grpc.CallOption) (*IfAInBFriendListResp, error) {
	client := pb.NewImUserServiceClient(m.cli.Conn())
	return client.IfAInBFriendList(ctx, in, opts...)
}

//  获取单聊会话的消息接收选项
func (m *defaultImUserService) GetSingleConversationRecvMsgOpts(ctx context.Context, in *GetSingleConversationRecvMsgOptsReq, opts ...grpc.CallOption) (*GetSingleConversationRecvMsgOptsResp, error) {
	client := pb.NewImUserServiceClient(m.cli.Conn())
	return client.GetSingleConversationRecvMsgOpts(ctx, in, opts...)
}

//  获取超级群成员列表 通过消息接收选项
func (m *defaultImUserService) GetUserListFromSuperGroupWithOpt(ctx context.Context, in *GetUserListFromSuperGroupWithOptReq, opts ...grpc.CallOption) (*GetUserListFromSuperGroupWithOptResp, error) {
	client := pb.NewImUserServiceClient(m.cli.Conn())
	return client.GetUserListFromSuperGroupWithOpt(ctx, in, opts...)
}

//  检查token
func (m *defaultImUserService) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenResp, error) {
	client := pb.NewImUserServiceClient(m.cli.Conn())
	return client.VerifyToken(ctx, in, opts...)
}

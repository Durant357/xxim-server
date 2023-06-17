// Code generated by goctl. DO NOT EDIT.
// Source: conversation.proto

package conversationservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConversationSettingUpdateReq  = pb.ConversationSettingUpdateReq
	ConversationSettingUpdateResp = pb.ConversationSettingUpdateResp
	FriendApplyReq                = pb.FriendApplyReq
	FriendApplyResp               = pb.FriendApplyResp
	GroupCreateReq                = pb.GroupCreateReq
	GroupCreateResp               = pb.GroupCreateResp

	ConversationService interface {
		// ConversationSettingUpdate 更新会话设置
		ConversationSettingUpdate(ctx context.Context, in *ConversationSettingUpdateReq, opts ...grpc.CallOption) (*ConversationSettingUpdateResp, error)
	}

	defaultConversationService struct {
		cli zrpc.Client
	}
)

func NewConversationService(cli zrpc.Client) ConversationService {
	return &defaultConversationService{
		cli: cli,
	}
}

// ConversationSettingUpdate 更新会话设置
func (m *defaultConversationService) ConversationSettingUpdate(ctx context.Context, in *ConversationSettingUpdateReq, opts ...grpc.CallOption) (*ConversationSettingUpdateResp, error) {
	client := pb.NewConversationServiceClient(m.cli.Conn())
	return client.ConversationSettingUpdate(ctx, in, opts...)
}
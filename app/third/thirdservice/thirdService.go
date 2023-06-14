// Code generated by goctl. DO NOT EDIT!
// Source: third.proto

package thirdservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CaptchaVerifyReq    = pb.CaptchaVerifyReq
	CaptchaVerifyResp   = pb.CaptchaVerifyResp
	EmailCodeSendReq    = pb.EmailCodeSendReq
	EmailCodeSendResp   = pb.EmailCodeSendResp
	EmailCodeVerifyReq  = pb.EmailCodeVerifyReq
	EmailCodeVerifyResp = pb.EmailCodeVerifyResp
	GetCaptchaReq       = pb.GetCaptchaReq
	GetCaptchaResp      = pb.GetCaptchaResp
	SmsCodeSendReq      = pb.SmsCodeSendReq
	SmsCodeSendResp     = pb.SmsCodeSendResp
	SmsCodeVerifyReq    = pb.SmsCodeVerifyReq
	SmsCodeVerifyResp   = pb.SmsCodeVerifyResp

	ThirdService interface {
		// SmsCodeSend 发送短信
		SmsCodeSend(ctx context.Context, in *SmsCodeSendReq, opts ...grpc.CallOption) (*SmsCodeSendResp, error)
		// SmsCodeVerify 验证短信
		SmsCodeVerify(ctx context.Context, in *SmsCodeVerifyReq, opts ...grpc.CallOption) (*SmsCodeVerifyResp, error)
		// EmailCodeSend 发送邮件
		EmailCodeSend(ctx context.Context, in *EmailCodeSendReq, opts ...grpc.CallOption) (*EmailCodeSendResp, error)
		// EmailCodeVerify 验证邮件
		EmailCodeVerify(ctx context.Context, in *EmailCodeVerifyReq, opts ...grpc.CallOption) (*EmailCodeVerifyResp, error)
		// GetCaptcha 获取图形验证码
		GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error)
		// CaptchaVerify 验证图形验证码
		CaptchaVerify(ctx context.Context, in *CaptchaVerifyReq, opts ...grpc.CallOption) (*CaptchaVerifyResp, error)
	}

	defaultThirdService struct {
		cli zrpc.Client
	}
)

func NewThirdService(cli zrpc.Client) ThirdService {
	return &defaultThirdService{
		cli: cli,
	}
}

// SmsCodeSend 发送短信
func (m *defaultThirdService) SmsCodeSend(ctx context.Context, in *SmsCodeSendReq, opts ...grpc.CallOption) (*SmsCodeSendResp, error) {
	client := pb.NewThirdServiceClient(m.cli.Conn())
	return client.SmsCodeSend(ctx, in, opts...)
}

// SmsCodeVerify 验证短信
func (m *defaultThirdService) SmsCodeVerify(ctx context.Context, in *SmsCodeVerifyReq, opts ...grpc.CallOption) (*SmsCodeVerifyResp, error) {
	client := pb.NewThirdServiceClient(m.cli.Conn())
	return client.SmsCodeVerify(ctx, in, opts...)
}

// EmailCodeSend 发送邮件
func (m *defaultThirdService) EmailCodeSend(ctx context.Context, in *EmailCodeSendReq, opts ...grpc.CallOption) (*EmailCodeSendResp, error) {
	client := pb.NewThirdServiceClient(m.cli.Conn())
	return client.EmailCodeSend(ctx, in, opts...)
}

// EmailCodeVerify 验证邮件
func (m *defaultThirdService) EmailCodeVerify(ctx context.Context, in *EmailCodeVerifyReq, opts ...grpc.CallOption) (*EmailCodeVerifyResp, error) {
	client := pb.NewThirdServiceClient(m.cli.Conn())
	return client.EmailCodeVerify(ctx, in, opts...)
}

// GetCaptcha 获取图形验证码
func (m *defaultThirdService) GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error) {
	client := pb.NewThirdServiceClient(m.cli.Conn())
	return client.GetCaptcha(ctx, in, opts...)
}

// CaptchaVerify 验证图形验证码
func (m *defaultThirdService) CaptchaVerify(ctx context.Context, in *CaptchaVerifyReq, opts ...grpc.CallOption) (*CaptchaVerifyResp, error) {
	client := pb.NewThirdServiceClient(m.cli.Conn())
	return client.CaptchaVerify(ctx, in, opts...)
}

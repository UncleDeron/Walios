// 消息处理模块
package process

import (
	"context"
	"fmt"
	"walios/events"
	"walios/pb"
	"walios/protocol"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"google.golang.org/protobuf/proto"
)

var (
	Ctx *context.Context
)

// Processor 消息处理模块
type Processor struct {
	Handlers map[uint32]HandlerFunc
}

type HandlerFunc func([]byte)

// NewProcessor 创建一个Processor实例
func NewProcessor(ctx *context.Context) *Processor {
	Ctx = ctx
	return &Processor{
		Handlers: make(map[uint32]HandlerFunc),
	}
}

// AddHandler 添加消息的处理方法
func (p *Processor) AddHandler(msgID uint32, handler HandlerFunc) {
	p.Handlers[msgID] = handler
}

// 登录响应处理
func loginResHandler(msg []byte) {
	msgData := &pb.ResData{}
	err := proto.Unmarshal(msg, msgData)
	if err != nil {
		runtime.EventsEmit(*Ctx, "notify", events.NotifyData{
			Msg:       "登录失败，未知错误",
			Type:      "danger",
			AutoClose: true,
			Duration:  3000,
		})
		fmt.Println(err)
	}
	var notifyType string
	if msgData.Code == pb.ResponseCode_SUCCESS {
		notifyType = "success"
	} else {
		notifyType = "danger"
	}
	runtime.EventsEmit(*Ctx, "notify", events.NotifyData{
		Msg:       msgData.Msg,
		Type:      notifyType,
		AutoClose: true,
		Duration:  3000,
	})
}

func registerResHandler(msg []byte) {
	msgData := &pb.ResData{}
	err := proto.Unmarshal(msg, msgData)
	if err != nil {
		runtime.EventsEmit(*Ctx, "notify", events.NotifyData{
			Msg:       "注册失败，未知错误",
			Type:      "error",
			AutoClose: true,
			Duration:  3000,
		})
		fmt.Println(err)
	}
	var notifyType string
	if msgData.Code == pb.ResponseCode_SUCCESS {
		notifyType = "success"
	} else {
		notifyType = "danger"
	}
	runtime.EventsEmit(*Ctx, "notify", events.NotifyData{
		Msg:       msgData.Msg,
		Type:      notifyType,
		AutoClose: true,
		Duration:  3000,
	})
}

func (p *Processor) BindHandler() {
	p.AddHandler(protocol.LoginResponse, loginResHandler)
	p.AddHandler(protocol.RegisterResponse, registerResHandler)
}

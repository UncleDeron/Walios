package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"walios/events"
	"walios/handler"
	"walios/pb"
	"walios/protocol"

	"google.golang.org/protobuf/proto"
)

// App struct
type App struct {
	ctx     context.Context
	client  *Client
	handler *handler.Handler
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
	// 建立客户端连接
	runtime.EventsEmit(a.ctx, "notify", &events.NotifyData{
		Msg:       "服务器连接中",
		Type:      "danger",
		AutoClose: false,
	})
	client, err := NewClient(&a.ctx)
	if err != nil {
		// runtime.(err)
		fmt.Println(err)
		_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.WarningDialog,
			Message: "服务器连接失败",
			Buttons: []string{"关闭"},
		})
		if err != nil {
			runtime.Quit(a.ctx)
		}
	}
	runtime.EventsEmit(a.ctx, "closeNotify", &events.NotifyData{
		Msg:       "连接成功",
		Type:      "info",
		AutoClose: true,
		Duration:  500,
	})
	go client.Read()
	a.client = client
	a.handler = handler.NewHandler(a.ctx)
	a.handler.InitAppHandler()
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	// 关闭连接
	a.client.Stop()
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
}

// 登录
func (a *App) Login(username string, password string) error {
	fmt.Println("Login:", username, password)
	loginData := &pb.ClientMsg{
		Action: pb.ClientActionType_LOGIN,
		Data: &pb.ClientMsg_LoginMsgData{
			LoginMsgData: &pb.LoginMsgData{
				UserAccount: username,
				UserPwd:     password,
			},
		},
	}
	msgData, err := proto.Marshal(loginData)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return err
	}
	if err = a.client.SendMsg(0, msgData); err != nil {
		fmt.Println("SendMsg error:", err)
		return err
	}
	return nil
}

// 注册
func (a *App) Register(regData protocol.RegisterData) error {
	fmt.Println("Register: ", regData)
	registerData := &pb.ClientMsg{
		Action: pb.ClientActionType_REGISTER,
		Data: &pb.ClientMsg_RegisterMsgData{
			RegisterMsgData: &pb.RegisterMsgData{
				UserAccount: regData.UserAccount,
				UserPwd:     regData.UserPwd,
				NickName:    regData.NickName,
				Email:       regData.Email,
			},
		},
	}
	msgData, err := proto.Marshal(registerData)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return err
	}
	if err = a.client.SendMsg(0, msgData); err != nil {
		fmt.Println("SendMsg error:", err)
		return err
	}
	return nil
}

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
	"walios/events"
	"walios/process"
	"walios/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/aceld/zinx/znet"
	"github.com/garyburd/redigo/redis"
)

type Client struct {
	Conn       net.Conn
	Processor  *process.Processor
	Config     *utils.ServerConfig
	RetryCount uint8
	Pool       *redis.Pool
	ctx        context.Context
}

// Client.Read() 客户端读取服务器发送的消息
func (c *Client) Read() {
	for {
		dp := znet.NewDataPack()
		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.Conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务器连接中断...")
				c.ReConnect()
				return
			}
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(c.Conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			if handler, ok := c.Processor.Handlers[msg.ID]; ok {
				handler(msg.Data)
			} else {
				fmt.Println("msg id not found:", msg.ID)
			}
		}
	}
}

func (c *Client) SendMsg(msgID uint32, msgData []byte) error {
	dp := znet.NewDataPack()
	// if zMsg.GetDataLen() > c.Config.MaxContentLength {
	// 	return errors.New("msg data too long")
	// }
	msg, _ := dp.Pack(znet.NewMsgPackage(msgID, msgData))
	_, err := c.Conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return err
	}

	return nil
}

func (c *Client) Stop() {
	c.Conn.Close()
}

// NewClient 创建一个客户端实例
func NewClient(ctx *context.Context) (client *Client, err error) {
	// 创建消息处理器并绑定处理事件
	processor := process.NewProcessor(ctx)
	processor.BindHandler()
	// 读取配置文件
	config := utils.ServerConfig{}
	err = utils.LoadConfig("./conf/server.json", &config)
	if err != nil {
		panic(err)
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		return nil, err
	}
	pool := initPool(&config.Redis)
	client = &Client{
		Conn:       conn,
		Processor:  processor,
		Config:     &config,
		RetryCount: 0,
		Pool:       pool,
		ctx:        *ctx,
	}
	return client, nil
}

// 重新连接
func (c *Client) ReConnect() (err error) {
	runtime.EventsEmit(c.ctx, "notify", &events.NotifyData{
		Msg:       "服务器连接中断，正在重新连接...",
		Type:      "danger",
		AutoClose: false,
	})
	for {
		if c.RetryCount > 3 {
			_, err := runtime.MessageDialog(c.ctx, runtime.MessageDialogOptions{
				Type:    runtime.WarningDialog,
				Message: "服务器连接失败",
				Buttons: []string{"关闭"},
			})
			if err != nil {
				runtime.Quit(c.ctx)
			}
			return errors.New("retry count more than 3")
		}
		c.Conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.Config.Host, c.Config.Port))
		if err != nil {
			c.RetryCount++
			fmt.Println("reconnect failed, retry count:", c.RetryCount)
			// 10s后重试
			time.Sleep(time.Second * 10)
			continue
		}
		c.RetryCount = 0
		runtime.EventsEmit(c.ctx, "closeNotify", &events.NotifyData{
			Msg:       "连接成功",
			Type:      "info",
			AutoClose: true,
			Duration:  500,
		})
		fmt.Println("reconnect success")
		break
	}
	return nil
}

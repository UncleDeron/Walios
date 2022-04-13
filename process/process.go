// 消息处理模块
package process

// Processor 消息处理模块
type Processor struct {
	Handlers map[uint32]HandlerFunc
}

type HandlerFunc func(msgData interface{})

// NewProcessor 创建一个Processor实例
func NewProcessor() *Processor {
	return &Processor{
		Handlers: make(map[uint32]HandlerFunc),
	}
}

// AddHandler 添加消息的处理方法
func (p *Processor) AddHandler(msgID uint32, handler HandlerFunc) {

}

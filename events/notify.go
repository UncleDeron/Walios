package events

type NotifyData struct {
	Msg       string `json:"msg"`
	Type      string `json:"type"`
	AutoClose bool   `json:"autoClose"`
	Duration  int    `json:"duration"`
}

func NewAutoCloseNotify(msg, notifyType string) *NotifyData {
	return &NotifyData{
		Msg:       msg,
		Type:      notifyType,
		AutoClose: true,
		Duration:  3000,
	}
}

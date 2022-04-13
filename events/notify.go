package events

type NotifyData struct {
	Msg       string `json:"msg"`
	Type      string `json:"type"`
	AutoClose bool   `json:"autoClose"`
	Duration  int    `json:"duration"`
}

type CloseNotifyData struct {
	NotifyData
}

package protocol

const (
	Client uint32 = iota
	User
	Sms
)

type RegisterData struct {
	UserAccount string `json:"userAccount"`
	UserPwd     string `json:"userPwd"`
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
}

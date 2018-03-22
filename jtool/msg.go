package jtool

const (
	//Session 常量
	S_UserName  = "UserName" //用户名
	S_UserId    = "UserId"
	S_CompanyId = "CompanyId"
	S_Template  = "Template"
)

/// 返回用的信息结构体
type Message struct {
	Result int         //0为失败 1为成功
	Data   interface{} //需要返回的信息结构体
	Msg    string      //提示信息
}

func NewMessage() *Message {
	m := new(Message)
	m.Result = 0
	m.Msg = "系统错误，请联系网站客服"
	return m

}

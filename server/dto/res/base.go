package res

const (
	SUCCESS = 200
	ERROR   = 500
)

type BaseRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok() BaseRes {
	return BaseRes{
		Code: SUCCESS,
		Msg:  "success",
		Data: nil,
	}
}
func Fail(code int, msg string) BaseRes {
	return BaseRes{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
func Err(msg string) BaseRes {
	return BaseRes{
		Code: ERROR,
		Msg:  msg,
		Data: nil,
	}
}
func Success(data interface{}) BaseRes {
	return BaseRes{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	}
}

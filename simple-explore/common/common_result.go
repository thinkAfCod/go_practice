package common

type ResultMsg string

const (
	SUCCESS ResultMsg = "成功"
	FAILED  ResultMsg = "失败"
	ERROR   ResultMsg = "服务器错误"
	INVALID ResultMsg = "请求无效"
)

func (rm *ResultMsg) status() int32 {
	switch *rm {
	case "失败":
		return 404
	case "服务器错误":
		return 500
	case "请求无效":
		return 202
	case "成功":
	default:
		return 200
	}
	return 0
}

//func (rm *ResultMsg) code() string {
//
//}

type CommonResult struct {
	Msg  string       `json:"msg" form:"msg"`
	Code int32        `json:"code" form:"code"`
	Data *interface{} `json:"data" form:"data"`
}

func Success(arg interface{}) CommonResult {
	msg := SUCCESS
	return createResult(msg, arg)
}

func Failed(arg interface{}) CommonResult {
	msg := FAILED
	return createResult(msg, arg)
}

func Error(arg interface{}) CommonResult {
	msg := ERROR
	return createResult(msg, arg)
}

func Invalid(arg interface{}) CommonResult {
	msg := INVALID
	return createResult(msg, arg)
}

func createResult(msg ResultMsg, arg interface{}) CommonResult {
	return CommonResult{
		Msg:  string(msg),
		Code: msg.status(),
		Data: &arg,
	}
}

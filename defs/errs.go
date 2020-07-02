package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorNotFound = ErrResponse{HttpSC: 404, Error: Err{Error: "很抱歉您访问的地址或者方法不存在", ErrorCode: -1}}
)

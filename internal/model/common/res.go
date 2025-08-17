package common

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

const (
	DEFAULT_SUCCESS_CODE = 200
	DEFAULT_FAIL_CODE    = 500
)

func Success[T any](data T, msg ...string) Result[T] {
	message := DEFAULT_SUCCESS_MESSAGE
	if len(msg) > 0 {
		message = msg[0]
	}
	return Result[T]{
		Code: DEFAULT_SUCCESS_CODE,
		Msg:  message,
		Data: data,
	}
}

func Fail[T any](msg ...string) Result[T] {
	message := DEFAULT_FAIL_MESSAGE
	if len(msg) > 0 {
		message = msg[0]
	}
	return Result[T]{
		Code: DEFAULT_FAIL_CODE,
		Msg:  message,
	}
}

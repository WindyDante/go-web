package common

type ServerErr struct {
	Msg string
	Err error
}

const (
	DEFAULT_FAIL_MESSAGE = "fail"
)

const (
	DATABASE_CONNECT_ERR = "数据库连接失败"
	CONFIG_LOAD_ERR      = "配置加载失败"
)

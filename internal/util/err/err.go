package err

import (
	"oa/internal/log"
	"oa/internal/model/common"

	"go.uber.org/zap"
)

func HandleError(s *common.ServerErr) string {
	if s.Err != nil {
		if s.Msg == "" || len(s.Msg) == 0 {
			s.Msg = s.Err.Error()
		}
		log.Error(s.Msg, zap.Error(s.Err))
	}
	return s.Msg
}

func HandlePanicError(s *common.ServerErr) {
	if s.Err != nil {
		if s.Msg == "" || len(s.Msg) == 0 {
			s.Msg = s.Err.Error()
		}
		log.Panic(s.Msg, zap.Error(s.Err))
	}
	panic(s.Msg)
}

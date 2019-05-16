package zk

import logger "github.com/sereiner/log"

type zkLogger struct {
	logger *logger.Logger
}

func (l *zkLogger) Printf(f string, c ...interface{}) {
	if l.logger != nil {
		l.logger.Printf(f, c...)
	}
}

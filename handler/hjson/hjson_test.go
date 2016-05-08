package hjson

import (
	"testing"

	"github.com/someonegg/golog"
)

func TestJson_all(t *testing.T) {
	log := golog.SubLoggerWithHandler(golog.RootLogger, Default)
	log.Info("json handler")
}

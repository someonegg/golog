package hmulti

import (
	"os"
	"testing"

	"github.com/someonegg/golog"
	"github.com/someonegg/golog/handler/hjson"
)

func TestMulti_all(t *testing.T) {
	log := golog.SubLoggerWithHandler(golog.RootLogger,
		New(hjson.Default, golog.NewHandler(os.Stderr)))
	log.Info("multi handler")
}

package golog

import (
	"errors"
	"testing"
)

func TestLogger_level(t *testing.T) {
	log := SubLoggerWithLevel(RootLogger, LevelInfo)
	log.Debug("debug level")
	log.Info("info level")
	log.Warn("warn level")
	log.Error("error level")
	//log.Panic("panic")
	//log.Fatal("fatal")
}

func TestLogger_predefs(t *testing.T) {
	log := SubLoggerWithFields(RootLogger, "pkg", "main")
	log.Info("msg")
	log1 := SubLoggerWithFields(log, "file", "a.go")
	log1.Info("msg 1")
	log.AddPredef("config", "dev")
	log1.Info("msg 2")
	log1.AddPredef("pkg", "mainex")
	log1.Info("msg 3")
	log1.DelPredef("pkg")
	log1.Info("msg 4")
	log1.DelPredef("pkg")
	log1.Info("msg 5")
	log.DelPredef("pkg")
	log1.Info("msg 6")
}

func TestLogger_special(t *testing.T) {
	log := SubLoggerWithFields(RootLogger, "error", errors.New("special_error"))
	log.Info("special")
}

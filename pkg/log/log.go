package log

import (
	"flag"

	"github.com/golang/glog"
)

func init() {
	flag.Lookup("logtostderr").Value.Set("true")
}

func Info(args ...interface{}) {
	glog.Info(args)
}

func Warn(args ...interface{}) {
	glog.Warning(args)
}

func Error(args ...interface{}) {
	glog.Error(args)
}

func Fatal(args ...interface{}) {
	glog.Fatal(args)
}

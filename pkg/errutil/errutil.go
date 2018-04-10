package errutil

import (
	"os"

	"github.com/golang/glog"
)

func exitWithCode(code int) {
	os.Exit(code)
}

var fatalErrHandler = func() {
	exitWithCode(1)
}

func CheckErrInfo(e error) {
	checkErr(e, nil)
}
func CheckErrFatal(e error) {
	checkErr(e, fatalErrHandler)
}

func checkErr(e error, handleErr func()) {
	if e != nil {
		glog.Error(e)
		if handleErr != nil {
			handleErr()
		}
	}
}

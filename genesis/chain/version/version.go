package version

import "strings"

const Maj = "0"
const Min = "1"
const Fix = "0"

const Version = Maj + "." + Min + "." + Fix

var commitVer string

func GetVersion() string {
	return Version
}

func GetCommitVersion() string {
	if len(commitVer) < 6 {
		return Version + "-"
	}
	return Version + "-" + commitVer[:6]
}

/*=======================unholy separator===========================*/

var (
	app_name string
)

func InitNodeInfo(app string) {
	if len(app_name) > 0 {
		return
	}
	if slc := strings.Split(app, "-"); len(slc) > 1 {
		app_name = slc[1]
	} else {
		app_name = app
	}
}

func AppName() string {
	return app_name
}

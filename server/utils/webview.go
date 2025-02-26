package utils

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"path/filepath"
)

func WebviewQuery() string {
	webview2Path := "lib" + string(filepath.Separator) + "webview2"
	if fileutil.IsExist(webview2Path) {
		return webview2Path
	} else {
		return ""
	}
}

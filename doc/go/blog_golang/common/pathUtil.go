package common

import (
	"fmt"
	"net/url"
)

type pathUtil struct{}

func (pathUtil) UrlDecode(urlString string) string {
	decodeUrl, err := url.QueryUnescape(urlString)
	if err != nil {
		fmt.Println(err)
	}
	return decodeUrl
}

// 获取文件扩展名
func getFileExpandedName(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}

var PathUtil = &pathUtil{}

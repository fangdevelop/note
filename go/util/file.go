package util

import (
	"os"
	"strings"
)

//判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return os.IsExist(err)
}

//根据文件路径创建文件夹
func MkdirWithFilePath(filePath string) error {
	paths := strings.Split(filePath, "/")
	paths[len(paths)-1] = ""
	for i, v := range paths {
		if i == len(paths)-1 {
			break
		}
		if i != 0 {
			paths[len(paths)-1] += "/"
		}
		paths[len(paths)-1] += v
	}
	return os.MkdirAll(paths[len(paths)-1], 0775)
}

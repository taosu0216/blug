package pkg

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetPostsLocation() string {
	// 获取当前文件的绝对路径
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	currentDir = filepath.Dir(currentDir)
	currentDir = filepath.Dir(currentDir)

	return fmt.Sprintf("%s/_posts/", currentDir)
}

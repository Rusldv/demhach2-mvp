package lib

import (
	"strings"
)

// IsRootDir усли это корневая дирректория сайта то true
func IsRootDir(path string) bool {
	//fmt.Println("path:", path)
	sl := strings.Split(path, "/")
	//fmt.Println(sl[len(sl)-1])
	if sl[len(sl)-1] == "" {
		return true
	}
	return false
}

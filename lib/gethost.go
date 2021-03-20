package lib

import "io/ioutil"

// GetHostDir ищет папку с контентом сайта
func GetHostDir(dir string) bool {
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		return false
	}
	return true
}

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 创建文件夹
func MkdirP(dir string) error {
	var err error
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, os.ModePerm)
			if err != nil {
				fmt.Printf("fail to create dir! err=%+v", err)
			}
			return err
		}
	}
	return err
}

func MoveFile(dir string, fileName string, destDir string) error {
	origin := filepath.Join(dir, fileName)
	err := os.Rename(origin, filepath.Join(destDir, fileName))
	if err != nil {
		fmt.Printf("fail to move file:%s", origin)
	}
	return err
}

func ReadFile(dir string, fileName string) []byte {
	data, err := ioutil.ReadFile(filepath.Join(dir, fileName))
	if err != nil {
		panic("read fail:" + filepath.Join(dir, fileName) + err.Error())
	}
	return data
}

func WriteFile(dir string, fileName string, data []byte) error {
	err := os.WriteFile(filepath.Join(dir, fileName), data, 0644)
	return err
}

func BaseFileName(fName string) string {
	return strings.TrimSuffix(fName, filepath.Ext(fName))
}

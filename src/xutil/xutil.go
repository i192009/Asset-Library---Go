package xutil

import (
	"github.com/go-playground/validator"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var Validate = validator.New()

// 切片去重
func RemoveDuplicateElement(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func DownloadFile(url string, path string) error {
	Logger.Info("DownloadFile url:", url, " path:", path)
	// 发起 GET 请求获取文件内容
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 创建本地文件
	dirPath := filepath.Dir(path)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		Logger.Errorf("DownloadFile MkdirAll is error.err:%s", err)
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		Logger.Errorf("DownloadFile Create is error.err:%s", err)
		return err
	}
	defer file.Close()

	// 将文件内容写入本地文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		Logger.Errorf("DownloadFile Copy is error.err:%s", err)
		return err
	}

	return nil
}

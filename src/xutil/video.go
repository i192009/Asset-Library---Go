package xutil

import (
	"gitlab.zixel.cn/go/framework/logger"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var Logger = logger.Get()

func VideoGenSnapshot(video string, output string, frame int) (err error) {
	Logger.Info("VideoGenSnapshot video:", video, " output:", output, " frame:", frame)
	if _, err = os.Stat(video); err != nil {
		Logger.Errorf("VideoGenSnapshot is error.err:%s", err)
		return
	}
	// 创建本地文件
	dirPath := filepath.Dir(output)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		Logger.Errorf("VideoGenSnapshot MkdirAll is error.err:%s", err)
		return
	}
	args := []string{"-i", video, "-f", "image2", "-vframes", strconv.Itoa(frame), output}
	err = exec.Command("ffmpeg", args...).Run()
	if err != nil {
		Logger.Errorf("VideoGenSnapshot is error.err:%s", err)
		return
	}

	Logger.Info("--snapshotPath--", output)

	return
}

func VideoMovFaststart(video string, output string) (err error) {
	var fi os.FileInfo
	if fi, err = os.Stat(video); err != nil {
		return err
	} else if fi.IsDir() {
		os.MkdirAll(output, 0755)
		err = filepath.WalkDir(video, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				err = VideoMovFaststart(path, filepath.Join(output, d.Name()))
			}
			Logger.Errorf("VideoMovFaststart is error.err:%s", err)
			return err
		})
	} else {
		args := []string{"-i", video, "-acodec", "copy", "-vcodec", "copy", "-movflags", "faststart", output}
		err = exec.Command("ffmpeg", args...).Run()
	}

	return
}

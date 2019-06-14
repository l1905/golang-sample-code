package log_simple

import (
	"fmt"
	"os"
	"time"
)

type Qwriter interface {
	Write(p []byte) (n int, err error)
}

type QlogData struct {
	LogChan chan[]byte	// 日志队列
	FileHandler *os.File
	realPath []byte
}

func NewQlogData() (*QlogData) {
	return &QlogData{LogChan : make(chan []byte, 100000), FileHandler:nil}
}

func (qlog *QlogData) Write(b []byte) (n int,err error)  {
	qlog.LogChan <- []byte(string(b))
	//因为[]byte是同一个地址， 要先转化string， 在[]byte
	return len(b), nil
}

func (qlog *QlogData) Consume() {

	for true {
		select {
		case log := <- qlog.LogChan:

			//fmt.Println("output:", string(log))
			//fmt.Println(string(log))
			//获取当前路径， 全部文件
			currentTime := time.Now()
			//通过配置文件，获取log前缀
			prefix_path := "/data/log/golang/v2out"
			fileName := currentTime.Format("2006-01-02-15")

			dirPath := fmt.Sprintf("%s/%d/%d", prefix_path, int(currentTime.Year()), int(currentTime.Month()))

			//文件路径
			realPath := fmt.Sprintf("%s/%s.log", dirPath, fileName)
			checkDirExists, _ := DirExists(dirPath)
			if checkDirExists == false {
				os.MkdirAll(dirPath, os.ModePerm)
			}
			if qlog.FileHandler == nil || string(qlog.realPath) != realPath {
				qlog.realPath = []byte(realPath)
				//有变更时， 才会去做修改， 不能每次都打开文件
				qlog.FileHandler, _ = os.OpenFile(realPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
			}

			qlog.FileHandler.Write(log)
		}
	}
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

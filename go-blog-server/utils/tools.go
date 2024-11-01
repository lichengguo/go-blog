package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 生成随机数字符串
func RandomStr() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(1024)
	return string(num)
}

// md5文件名
func Md5Str(fileName string) (newFileName string) {
	// md5加密
	h := md5.New()
	io.WriteString(h, RandomStr())
	io.WriteString(h, fileName)
	// 获取文件名后缀
	suffix := ""
	suffixSlice := strings.Split(fileName, ".")
	if len(suffixSlice) == 1 {
		// 文件名没有后缀，直接返回
		newFileName = hex.EncodeToString(h.Sum(nil))
		return
	}
	suffix = suffixSlice[len(suffixSlice)-1]
	newFileName = fmt.Sprintf("%s.%s", hex.EncodeToString(h.Sum(nil)), suffix)
	return
}

// 初始化常用目录
func InitDir() {
	// 当前程序执行的目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取当前程序执行目录失败")
		return
	}

	// 日志文件目录拼接
	logDirPath := fmt.Sprintf("%s/%s", dir, LogPath)
	// 判断目录是否存在
	_, err = os.Stat(logDirPath)
	if err != nil {
		// 不存在则创建目录
		//os.MkdirAll(logDirPath, 0755)
		os.MkdirAll(logDirPath, 0755)
	}

	// 图片目录
	imgDirPath := fmt.Sprintf("%s/%s", dir, ImgPath)
	_, e := os.Stat(imgDirPath)
	if e != nil {
		os.MkdirAll(imgDirPath, 0755)
	}
}

package imgutil

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"gin-vue-admin/global"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	p "path"
	"path/filepath"
	"strconv"
	"strings"
)

/* 参考博客 https://github.com/Aszzo/tinker */

// ResizeJpg 压缩jpg  quality 设置为1时,图片最小
func ResizeJpg(path string, quality int) {
	defer func() {
		if r := recover(); r != nil {
			global.GVA_LOG.Error(fmt.Sprint(r))
		}
	}()
	// open file
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Printf("open file error:%s \n", err)
	}
	resized := IsResize(file, strconv.Itoa(quality))

	if !resized {
		size := GetFileSize(path)
		fmt.Printf(">正在压缩图片：%v，图片大小： %.2f kb \n", p.Base(path), size)
		// decode png into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Printf("decode jpg fail: %s \n", err)
		}
		defer file.Close()

		m := resize.Resize(0, 0, img, resize.NearestNeighbor)

		out, err := os.Create(path)
		if err != nil {
			log.Println(err)
		}
		defer out.Close()

		_ = jpeg.Encode(out, m, &jpeg.Options{Quality: quality})
		defer func() {
			size := GetFileSize(path)
			WriteHash(path, strconv.Itoa(quality))
			fmt.Printf("压缩成功，压缩后文件大小%.2f kb \n", size)
		}()
	}
}

// ResizePng 压缩png
func ResizePng(path string, quality int, filename string) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	resized := IsResize(file, strconv.Itoa(quality))

	if !resized {
		size := GetFileSize(path)
		fmt.Printf(">正在压缩图片：%v，图片大小： %.2f kb \n ", p.Base(path), size)

		str, _ := os.Executable()
		dir := filepath.Dir(str)
		qualityString := strconv.Itoa(quality)
		cmd := exec.Command(dir+"/pngquant", path, "--ext=.png", "--force", "--quality", qualityString)
		//cmd := exec.Command("./pngquant", path, "--ext=.png", "--force", "--quality", qualityString)
		defer func() {
			size := GetFileSize(path)
			WriteHash(path, strconv.Itoa(quality))
			fmt.Printf("压缩成功，压缩后图片大小%.2f kb \n", size)
		}()
		if err := cmd.Run(); err != nil { // 运行命令
			log.Println(err)
		}
	}

}

// GetFileSize 读取文件大小
func GetFileSize(path string) float64 {
	fileInfo, _ := os.Stat(path)
	size := float64(fileInfo.Size()) / float64(1000)
	return size
}

// 判断文件|文件夹是否存在
func isExits(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// GetFileHash 计算文件哈希值
func GetFileHash(file *os.File) string {
	// 计算文件哈希值
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Println(err)
		return ""
	}
	sum := hash.Sum([]byte(""))
	// file.seek(0,0)方法，将文件指针移到开头,
	// https://blog.csdn.net/weixin_40306555/article/details/102779246
	file.Seek(0, 0)
	fileHash := new(bytes.Buffer)
	_, hashErr := fmt.Fprintf(fileHash, "%x", sum)
	if hashErr != nil {
		return ""
	}
	return fileHash.String()
}

// IsResize 图片是否被压缩过
func IsResize(file *os.File, quality string) bool {
	fileHash := GetFileHash(file)
	// 读取.tinker目录下的resize.log信息
	// 里面会存储压缩的历史记录
	isExitFile := isExits(".tinker")
	if !isExitFile {
		os.Mkdir(".tinker", 0777)
	}

	logName := ".tinker/resize.log"
	f, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return false
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return false
	}
	containsHash := strings.Contains(string(content), fileHash+"-"+quality)
	// 如果有压缩日志，直接返回
	if containsHash {
		return containsHash
	}
	return false
}

// WriteHash 写入hash
func WriteHash(path string, quality string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	fileHash := GetFileHash(file)
	logName := ".tinker/resize.log"
	f, err := os.OpenFile(logName, os.O_WRONLY|os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()
	// 没有压缩日志，返回false,并且创建新的日志
	write := bufio.NewWriter(f)
	write.WriteString(fileHash + "-" + quality + "\n")
	writeErr := write.Flush()
	if writeErr != nil {
		log.Println(writeErr)
		return writeErr
	}
	return nil
}

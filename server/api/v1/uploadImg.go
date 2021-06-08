package v1

// go 代码
import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils/imgutil"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

func FileUpload(c *gin.Context) {
	_, file, err := c.Request.FormFile("file")
	if err != nil {
		resErr(c)
		log.Error(err)
		return
	}
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	if ext != ".jpg" {
		response.FailWithMessage("请上传.jpg文件类型的格式", c)
	}
	rootPath := global.GVA_CONFIG.File.RootPath
	imgPath := global.GVA_CONFIG.File.ImgPath
	tempPath := time.Now().Format("2006/01/")
	fileName := fmt.Sprint(rand.Int()) + path.Ext(file.Filename)
	if err = os.MkdirAll(rootPath+imgPath+tempPath, os.ModePerm); err != nil {
		resErr(c)
		log.Error(err)
		return
	}
	f, err := file.Open()
	if err != nil {
		resErr(c)
		log.Error(err)
	}
	fullPath := rootPath + imgPath + tempPath + fileName
	global.GVA_LOG.Info(fmt.Sprint("文件路径为:", fullPath))
	out, err := os.Create(fullPath)
	defer out.Close() //创建文件的 defer 关闭了
	if err != nil {
		log.Error(err)
		resErr(c)
		return
	}
	_, err = io.Copy(out, f)
	if err != nil {
		log.Error(err)
		resErr(c)
		return
	}
	if err = f.Close(); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	go imgutil.ResizeJpg(fullPath, 7) //文件压缩
	resOkCodeData(0, "ok", imgPath+tempPath+fileName, c)
}

// FileUploadFull 不压缩的上传
func FileUploadFull(c *gin.Context) {
	_, file, err := c.Request.FormFile("file")
	if err != nil {
		resErr(c)
		log.Error(err)
		return
	}
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	if ext != ".jpg" {
		response.FailWithMessage("请上传.jpg文件类型的格式", c)
	}
	rootPath := global.GVA_CONFIG.File.RootPath
	imgPath := global.GVA_CONFIG.File.ImgPath
	tempPath := time.Now().Format("2006/01/")
	fileName := fmt.Sprint(rand.Int()) + path.Ext(file.Filename)
	if err = os.MkdirAll(rootPath+imgPath+tempPath, os.ModePerm); err != nil {
		resErr(c)
		log.Error(err)
		return
	}
	f, err := file.Open()
	if err != nil {
		resErr(c)
		log.Error(err)
	}
	fullPath := rootPath + imgPath + tempPath + fileName
	global.GVA_LOG.Info(fmt.Sprint("文件路径为:", fullPath))
	out, err := os.Create(fullPath)
	defer out.Close() //创建文件的 defer 关闭了
	if err != nil {
		log.Error(err)
		resErr(c)
		return
	}
	_, err = io.Copy(out, f)
	if err != nil {
		log.Error(err)
		resErr(c)
		return
	}
	if err = f.Close(); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	resOkCodeData(0, "ok", imgPath+tempPath+fileName, c)
}

func resErr(c *gin.Context) {
	resErrCodeMsg(-1, "操作失败", c)
}

func resErrCodeMsg(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  msg,
	})
}
func resOk(c *gin.Context) {
	resOkCodeData(1, "success", nil, c)
}
func resOkCodeData(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

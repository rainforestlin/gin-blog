package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/pkg/utils"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 检查目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt 检查后缀是否是配置中允许的后缀
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowsExis {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// CheckMaxSize 检查文件大小是否超出限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*2014 {
			return true
		}
	}
	return false
}

// CheckPermission 检查目录是否有权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// CreateSavePath 创建保存上传文件的目录。若涉及到的目录已存在，则不进行任何操作
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.Mkdir(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// SaveFile 保存上传的文件。该方法通过调用os.Create方法创建目标地址文件，
// 在通过file.Open方法打开源地址的文件，结合io.Copy方法实现两者之间的文件内容拷贝
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()

	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)

	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

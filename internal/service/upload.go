package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadeSavePath := upload.GetSavePath()
	dst := uploadeSavePath + "/" + fileName

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckSavePath(dst) {
		err := upload.CreateSavePath(dst, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}

	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	if upload.CheckPermission(dst) {
		return nil, errors.New("insurfficient file permission")
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

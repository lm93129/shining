package model

import (
	"github.com/jinzhu/gorm"
)

type AppManage struct {
	gorm.Model
	Name              string `gorm:"not null"`
	BundleId          string
	Version           string `gorm:"not null"`
	Build             string
	Size              int64
	AppType           string `gorm:"not null"`
	OssUrl            string `gorm:"not null"`
	ProjectId         string `gorm:"not null"`
	UpdateDescription string
}

package appfile

import (
	"appdlserver/model"
	"appdlserver/util/appfileparser"
	"fmt"
	"log"
	"os"
)

type Plist struct {
}

func (server *Plist) Show(id string) []byte {
	var appshow model.AppManage

	err := model.DB.First(&appshow, id).Error
	if err != nil {
		log.Panicln(err)
	}
	singUrl := fmt.Sprintf("https://%s.%s/%s",
		os.Getenv("OSS_BUCKET"),
		os.Getenv("OSS_END_POINT"),
		appshow.OssUrl,
	)
	plistdata, _ := appfileparser.EnCodePlist(appshow, singUrl)
	return plistdata
}

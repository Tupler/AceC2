package Task

import (
	"ReAce/Logger"
	"os"
)

type DownLoader struct {
	FileName string
}

func (downloader *DownLoader) SaveFile(fileBytes []byte) {
	err := os.Mkdir("downloads", 0750)
	if err != nil && !os.IsExist(err) {
		Logger.ALogger.Error(err.Error())
		return
	}
	err = os.WriteFile("downloads/"+downloader.FileName, fileBytes, 0655)
	if err != nil {
		Logger.ALogger.Error(err.Error())
		return
	}
	Logger.ALogger.Success("保存文件到:", "downloads/"+downloader.FileName)

}

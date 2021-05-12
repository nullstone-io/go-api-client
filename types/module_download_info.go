package types

import "net/url"

type ModuleDownloadInfo struct {
	FileExtension string
	DataLength    int
	DownloadUrl   url.URL
}

func (i ModuleDownloadInfo) GetterUrl() string {
	getterUrl := i.DownloadUrl
	q := getterUrl.Query()
	q.Set("archive", i.FileExtension)
	getterUrl.RawQuery = q.Encode()
	return getterUrl.String()
}

package main

type downloadFromNetDisk struct {
	secret   DynamicSecret
	filePath string
}

func (d *downloadFromNetDisk) DownloadFile() {
	if err := d.loginCheck(); err != nil {
		//todo重新登录
	}
	d.downloadFileFromCloud(d.filePath)
}

func (d *downloadFromNetDisk) loginCheck() error {
	d.checkSecret(d.secret.GetSecret())
	return nil
}

func (d *downloadFromNetDisk) downloadFileFromCloud(file string) {

}

func (d *downloadFromNetDisk) checkSecret(file string) {
	//todo 调用阿里云的验证接口去验证密码是否有效
}

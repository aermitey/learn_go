package main

func main() {
	d := &downloadFromNetDisk{
		secret: &mobileTokenDynamic{
			mobileNumber: "12312344321",
		},
		filePath: "面向对象编程.ppt",
	}
	d.DownloadFile()
}

type DynamicSecret interface {
	GetSecret() string
}

type mobileTokenDynamic struct {
	mobileNumber string
}

func (m mobileTokenDynamic) GetSecret() string {
	return "123321"
}

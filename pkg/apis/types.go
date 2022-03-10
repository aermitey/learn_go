package apis

type PersonalInformation struct {
	Name string `json:"name"` //注解会将Name更换为name(注解的值)
	Sex  string `json:"sex"`
	//sex    string  `json:"sex"` //私有(不可导出)成员变量在序列化反序列化时会被忽略
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
}

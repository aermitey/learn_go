package calc

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall <= 0 {
		panic("身高不能为0或负数")
	}
	//todo 验证体重和发型
	return weight / (tall * tall)
}

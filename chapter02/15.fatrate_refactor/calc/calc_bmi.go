package calc

import "fmt"

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		err = fmt.Errorf("身高不能为0或负数")
	} else if weight <= 0 {
		err = fmt.Errorf("体重不能为0或负数")
	} else if tall <= 0 && weight <= 0 {
		err = fmt.Errorf("身高体重不能为0或负数")
	}
	bmi = weight / (tall * tall)
	return
}

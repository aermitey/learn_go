package main

import "fmt"

func main() {
	//xqInfo := []string{"小强", "男", "在职"}
	//xlInfo := []string{"小李", "男", "在职"}
	//xsInfo := []string{"小苏", "女", "在职"}
	//多维数组
	//newPersonInfos := [3][3]string{
	//	[3]string{"小强", "男", "在职"},
	//	[3]string{"小李", "男", "在职"},
	//	[3]string{"小苏", "女", "在职"},
	//}
	//for _, val := range newPersonInfos {
	//	fmt.Println(val)
	//}
	//
	//newPersonInfos2 := [...][3]string{
	//	[3]string{"小强", "男", "在职"},
	//	[3]string{"小李", "男", "在职"},
	//	[3]string{"小苏", "女", "在职"},
	//	[3]string{"小苏", "女", "在职"},
	//	[3]string{"小苏", "女", "在职"},
	//	[3]string{"小苏", "女", "在职"},
	//}
	//for _, val := range newPersonInfos2 {
	//	fmt.Println(val)
	//}
	//fmt.Println(newPersonInfos2[1][1])

	newPersonInfo3 := [3][][3]string{
		[][3]string{
			[3]string{"小强", "男", "在职"},
			[3]string{"小李", "男", "在职"},
			[3]string{"小苏", "女", "在职"},
		},
		[][3]string{
			[3]string{"小强", "男", "在职"},
			[3]string{"小李", "男", "在职"},
			[3]string{"小苏", "女", "在职"},
		},
		[][3]string{
			[3]string{"小强", "男", "在职"},
			[3]string{"小李", "男", "在职"},
			[3]string{"小苏", "女", "在职"},
		},
	}
	newPersonInfo3[2] = append(newPersonInfo3[2],
		[3]string{"小强", "男", "在职"},
	)
	fmt.Println(newPersonInfo3)
	//for d1, val := range newPersonInfo3 {
	//	a := val
	//	for d2, val := range a {
	//		fmt.Println(d1, d2, val)
	//	}
	//	fmt.Println(val)
	//}
}

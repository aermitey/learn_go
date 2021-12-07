package main

import "fmt"

type line struct {
	cod1 [2]float64
	cod2 [2]float64
}

func main() {
	for {
		var l1 line

		fmt.Println("输入l1的坐标(x1,y1):")
		fmt.Scan(&l1.cod1[0])
		fmt.Scan(&l1.cod1[1])

		fmt.Println("输入l1的坐标(x2,y2):")
		fmt.Scan(&l1.cod2[0])
		fmt.Scan(&l1.cod2[1])

		fmt.Println(l1)

		var l2 line

		fmt.Println("输入l2的坐标(x1,y1):")
		fmt.Scan(&l2.cod1[0])
		fmt.Scan(&l2.cod1[1])

		fmt.Println("输入l2的坐标(x2,y2):")
		fmt.Scan(&l2.cod2[0])
		fmt.Scan(&l2.cod2[1])

		fmt.Println(l1, l2)

		var x float64
		var y float64
		//x = (l1.cod1[0]*l1.cod2[1]-l1.cod2[0]*l1.cod1[1])/(l1.cod2[0]-l1.cod1[0]) + (l2.cod2[0]*l2.cod1[1]-l2.cod1[0]*l2.cod2[1])/(l2.cod2[0]-l2.cod1[0])
		//y = (l1.cod2[1]-l1.cod1[0])/(l1.cod2[0]-l1.cod1[0]) - (l2.cod2[1]-l2.cod1[1])/(l2.cod2[0]-l2.cod1[0])
		//z := x / y
		//if ((l1.cod1[0]-z)*(z-l1.cod2[0])) >= 0 && (l2.cod1[0]-z)*(z-l2.cod2[0]) >= 0 {
		//	fmt.Println("平行")
		//} else {
		//	fmt.Println("不平行")
		//}
		x = (l1.cod2[1] - l1.cod1[1]) / (l1.cod2[0] - l1.cod1[0])
		y = (l2.cod2[1] - l2.cod1[1]) / (l2.cod2[0] - l2.cod1[0])
		if x == y {
			fmt.Println("平行")
		} else {
			fmt.Println("不平行")
		}
		var whetherContinue string
		fmt.Print("是否计算下一个(y/n):")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}

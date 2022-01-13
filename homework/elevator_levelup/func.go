package main

//func inputFloor() []int { //输入目标楼层
//	var tg []int
//	fmt.Println("获取楼层信息") //用户在电梯外面点击使用电梯或在电梯里面点击目标楼层
//	fmt.Scanln(tg)
//	return tg
//}

func GetMinNum(ary []int) int { //获取切片中的最小值
	if len(ary) == 0 {
		return 0
	}

	maxVal := ary[0]
	for i := 1; i < len(ary); i++ {
		if maxVal > ary[i] {
			maxVal = ary[i]
		}
	}

	return maxVal
}

func GetMaxNum(ary []int) int { //获取切片中的最大值
	if len(ary) == 0 {
		return 0
	}

	maxVal := ary[0]
	for i := 1; i < len(ary); i++ {
		if maxVal < ary[i] {
			maxVal = ary[i]
		}
	}

	return maxVal
}

func delItem(slice []int, item int) []int { //移除切片中一个元素
	for i := 0; i < len(slice); i++ {
		if item == slice[i] {
			slice = append(slice[:i], slice[i+1:]...)
			i = i - 1
		}
	}
	return slice
}

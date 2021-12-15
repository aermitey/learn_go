package main

import "fmt"

func guess(left, right uint) (guessed uint) {
	guessed = (left + right) / 2
	var getFromInput string
	//fmt.Scanln(&guessed)
	fmt.Println("猜测结果:", guessed)
	fmt.Print("left:", left)
	fmt.Println("right:", right)
	fmt.Println("如果高了，输入1，低了输入0，对了输入9")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case "1":
		if left == right {
			fmt.Println("赖皮")
			return
		}
		guess(left, guessed-1)

	case "0":
		if left == right {
			fmt.Println("赖皮")
			return
		}
		guess(guessed+1, right)
	case "9":
		fmt.Println(guessed)
	}
	return
}

//func main() {
//	guess(1, 100)
//}

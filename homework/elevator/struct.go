package main

import (
	"fmt"
	"time"
)

type elevator struct {
	floor    int   //"电梯所在层数"
	status   int   //"电梯楼层数" 0表示静止，1表示向上，2表示向下
	target   []int //"电梯目标层数"
	maxFloor int   //"目标楼层最大值"
	minFloor int   //"目标楼层最小值"
}

//用户输入目标楼层或在电梯外点击使用电梯调用此函数
func (e *elevator) add(tg ...int) {
	//tg = inputFloor()
	e.target = append(e.target, tg...)
	e.minFloor = GetMinNum(e.target)
	e.maxFloor = GetMaxNum(e.target)
	if e.target[0] > e.floor {
		e.status = 1 //电梯向上
	} else if e.target[0] == e.floor {
		e.status = 0 //电梯静止
	} else {
		e.status = 2 //电梯向下
	}
}

//当运行到一个指定楼层时调用此函数
func (e *elevator) arrive(f int) (err error) {
	fmt.Printf("已到达%d楼层，打开电梯门\n", f)
	e.target = delItem(e.target, f)
	e.minFloor = GetMinNum(e.target)
	e.maxFloor = GetMaxNum(e.target)
	time.Sleep(3000 * time.Millisecond)
	if len(e.target) == 0 {
		e.floor = f
		e.status = 0
		err = fmt.Errorf("电梯结束运行，关闭电梯门，停在%d层", f)
		return
	}
	fmt.Println("关闭电梯门，继续运行")
	return nil
}

//电梯运行时函数
func (e *elevator) up() (err error) {
	for i := e.floor; i < e.maxFloor; i++ {
		time.Sleep(1000 * time.Millisecond)
		e.floor++
		fmt.Printf("电梯正在上升，当前处于%d层\n", e.floor)
		for _, item := range e.target {
			if e.floor == item {
				err = e.arrive(item)
				if err != nil {
					return err
				}
			}
		}
	}
	err = e.down()
	if err != nil {
		return err
	}
	return nil
}

func (e *elevator) down() (err error) {
	for i := e.floor; i > e.minFloor; i-- {
		time.Sleep(1000 * time.Millisecond)
		e.floor--
		fmt.Printf("电梯正在下降，当前处于%d层\n", e.floor)
		for _, item := range e.target {
			if e.floor == item {
				err = e.arrive(item)
				if err != nil {
					return err
				}
			}
		}
	}
	err = e.up()
	if err != nil {
		return err
	}
	return nil
}

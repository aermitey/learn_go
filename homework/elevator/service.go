package main

type service struct {
	ele elevatorOperation
}

func (s service) ElevatorOperation(e *elevator) (err error) {
	if e.status == 0 {
		e.target = delItem(e.target, e.target[0])
		if e.target[0] > e.floor {
			e.status = 1 //电梯向上
			err = s.ele.up()
		} else {
			e.status = 2 //电梯向下
			err = s.ele.down()
		}
		if err != nil {
			return err
		}
	}
	if e.status == 1 {
		err = s.ele.up()
		if err != nil {
			return err
		}
	}
	if e.status == 2 {
		err = s.ele.down()
		if err != nil {
			return err
		}
	}
	return nil
}

package main

import (
	"fmt"
	"math/rand"
)

type Chinese = int
type Math = int
type English = int

func main() {
	var mathScore1 int = 100
	var mathScore2 Math = 100
	mathScore2 = mathScore1
	fmt.Println(mathScore2)
	vt := &voteGame{
		students: []*student{
			&student{name: "小强"},
			&student{name: "小李"},
			&student{name: "小张"},
			&student{name: "小王"},
			&student{name: "小陈"},
			&student{name: "小周"},
			&student{name: "小胡"},
			&student{name: "小冯"},
		},
	}
	m := vt.goRun()
	m.Distribute()
	fmt.Println(m)
	fmt.Println(vt.students)
	var stdXQ = &student{name: "小强"}
	var monXQ monitor = monitor(*stdXQ)
	fmt.Println(monXQ)

}

func getScoresOfStudent(name string) (Chinese, Math, English) {
	chineseScore := 100
	mathScore := 100
	englishScore := 100
	return chineseScore, mathScore, englishScore
}

type student struct {
	name     string
	agree    int
	disagree int
}

func (s *student) voteAgree(target *student) {
	target.agree++
}

func (s *student) voteDisagree(target *student) {
	target.disagree++
}

type voteGame struct {
	students []*student
}

func (v voteGame) goRun() *monitor {
	for _, s := range v.students {
		randInt := rand.Int()
		if randInt%2 == 0 {
			s.voteAgree(v.students[randInt%len(v.students)])
		} else {
			s.voteDisagree(v.students[randInt%len(v.students)])
		}
	}
	maxScore := -1
	maxScoreIndex := -1
	for i, s := range v.students {
		if s.agree > maxScore {
			maxScore = s.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 {
		return (*monitor)(v.students[maxScoreIndex])
	}
	return nil
}

type monitor student

func (m *monitor) Distribute() {
	fmt.Println("发作业啦")
}

type foo func() //能定义指针之外的任意类型

func (f *foo) test() {

}

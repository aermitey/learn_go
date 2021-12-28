package main

import (
	"testing"
)

func TestCalc_BMI(t *testing.T) {
	inputPerson := &Person{}
	inputPerson.tall = 1.7
	inputPerson.weight = 70
	expectedOutput := 24.221453287197235
	t.Logf("开始计算，输入：tall:%f, weight:%f, 期望结果：%f", inputPerson.tall, inputPerson.weight, expectedOutput)
	err := Calc{}.BMI(inputPerson)
	actualOutput := inputPerson.bmi
	t.Logf("实际得到：%f， error：%v", actualOutput, err)
	if err != nil {
		t.Fatalf("得到错误%v，与预期不符", err)
	}
	if actualOutput != expectedOutput {
		t.Errorf("预期值为%f，得到值为%f，与预期不符", expectedOutput, actualOutput)
	}
}

func TestCalc_FatRate(t *testing.T) {
	inputPerson := &Person{}
	inputPerson.bmi = 23.0
	inputPerson.age = 23
	inputPerson.sex = "男"
	expectedOutput := 0.16690000000000002
	t.Logf("开始计算，输入：bmi:%f, age:%d, sex:%s,期望结果：%f", inputPerson.bmi, inputPerson.age, inputPerson.sex, expectedOutput)
	err := Calc{}.FatRate(inputPerson)
	if err != nil {
		t.Fatalf("得到错误%v，与预期不符", err)
	}
	actualOutput := inputPerson.fatRate
	if actualOutput != expectedOutput {
		t.Errorf("预期值为%f，得到值为%f，与预期不符", expectedOutput, actualOutput)
	}
}

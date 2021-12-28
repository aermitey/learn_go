package calc

import (
	"testing"
)

func TestCalcFatRate(t *testing.T) {
	inputBmi, inputAge, inputSex := 23.0, 23, "男"
	expectedOutput := 0.16690000000000002
	t.Logf("开始计算，输入：bmi:%f, age:%d, sex:%s 期望结果：%f", inputBmi, inputAge, inputSex, expectedOutput)
	actualOutput, err := CalcFatRate(inputBmi, inputAge, inputSex)
	t.Logf("实际得到：%f， error：%v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got %v", err)
	}
	if expectedOutput != actualOutput {
		t.Fatalf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}

package calc

import "testing"

func TestCalcBMI1(t *testing.T) {
	inputTall, inputWeight := 1.0, 1.0
	expectedOutput := 1.0
	t.Logf("开始计算，输入：tall:%f, weight:%f, 期望结果：%f", inputTall, inputWeight, expectedOutput)
	actualOutput, err := CalcBMI(inputTall, inputWeight)
	t.Logf("实际得到：%f， error：%v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}

func TestCalcBMI2(t *testing.T) {
	inputTall, inputWeight := 1.0, 1.0
	expectedOutput := 1.0
	t.Logf("开始计算，输入：tall:%f, weight:%f, 期望结果：%f", inputTall, inputWeight, expectedOutput)
	actualOutput, err := CalcBMI(inputTall, inputWeight)
	t.Logf("实际得到：%f， error：%v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}

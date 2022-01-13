package main

import "testing"

func TestCase1(t *testing.T) {
	f := &FatRateRank{}
	f.inputRecord("王强", 0.38)
	f.inputRecord("王强", 0.32)
	{
		rankOfWQ, fatRateOfWQ := f.getRank("王强")
		if rankOfWQ != 1 {
			t.Fatalf("预期王强第一，但是得到的是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是：%f", fatRateOfWQ)
		}
	}
	f.inputRecord("李静", 0.28)
	{
		rankOfWQ, fatRateOfWQ := f.getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("预期王强第二，但是得到的是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfLJ, fatRateOfLJ := f.getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是：%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是：%f", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	f := &FatRateRank{}
	f.inputRecord("王强", 0.38)
	f.inputRecord("张伟", 0.38)
	f.inputRecord("李静", 0.28)
	{
		rankOfLJ, fatRateOfLJ := f.getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是：%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是：%f", fatRateOfLJ)
		}
	}
	{
		rankOfWQ, fatRateOfWQ := f.getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("预期王强第二，但是得到的是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂率为0.38，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfZW, fatRateOfZW := f.getRank("张伟")
		if rankOfZW != 2 {
			t.Fatalf("预期张伟第二，但是得到的是：%d", rankOfZW)
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("预期张伟的体脂率为0.38，但是得到的是：%f", fatRateOfZW)
		}
	}
}

func TestCase3(t *testing.T) {
	f := &FatRateRank{}
	f.inputRecord("王强", 0.38)
	f.inputRecord("李静", 0.28)
	f.inputRecord("张伟")

	{
		rankOfLJ, fatRateOfLJ := f.getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是：%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是：%f", fatRateOfLJ)
		}
	}
	{
		rankOfWQ, fatRateOfWQ := f.getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("预期王强第二，但是得到的是：%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂率为0.38，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfZW, _ := f.getRank("张伟")
		if rankOfZW != 3 {
			t.Fatalf("预期张伟第三，但是得到的是：%d", rankOfZW)
		}
	}
}

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	learngo_tools "learngo.tools"
	"learngo/chapter03/01.fatrate_refactor/calc"
)

func main() {
	//录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)
	//argument := os.Args
	//fmt.Println(argument)
	//
	//name = argument[1]
	//sex = argument[2]
	//tall = argument[3]
	//weight = argument[4]
	//age = argument[5]
	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体质计算器，根据身高、体重、性别、年龄计算体质，并给出建议",
		Long:  "该体质计算器是基于BMI的体质计算器。。。。。",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name", name)
			fmt.Println("sex", sex)
			fmt.Println("tall", tall)
			fmt.Println("weight", weight)
			fmt.Println("age", age)
			//计算
			bmi := calc.CalcBMI(tall, weight)
			fatRAte := calc.CalcFatRate(bmi, age, sex)
			fmt.Println(fatRAte)
			//评估结果
			fmt.Println(learngo_tools.MaxNum(1, 2))
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.ExecuteC()

}

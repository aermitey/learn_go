package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
	"log"
	"reflect"
	"testing"
)

func TestMarshalJson(t *testing.T) { //json只适用于大写(可导出)的变量，改为小写之后序列化之后的json文件中会不会出现相关数据
	personalInformation := PersonalInformation{
		Name:   "小强...'",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%v\n", personalInformation)
	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal 的结果是(原生)：", data)
	fmt.Println("marshal 的结果是(string)：", string(data))
}

func TestUnmarshalJson(t *testing.T) {
	//同理反序列化时，非可导出变量的数据会丢失变为空
	data := `{"name":"小强","sex":"男","tall":1.7,"weight":71,"age":35}`
	personalInformation := &PersonalInformation{}
	err := json.Unmarshal([]byte(data), personalInformation)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", personalInformation)
	fmt.Println("数据", personalInformation)
	fmt.Println("数据", reflect.TypeOf(personalInformation))
}

func TestMarshalYaml(t *testing.T) { //支持json注解，不同的数据不能放在同一行，可以注释
	personalInformation := PersonalInformation{
		Name:   "小强...'",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%v\n", personalInformation)
	data, err := yaml.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("yaml marshal 的结果是(原生)：", data)
	fmt.Println("yaml marshal 的结果是(string)：", string(data))
}

func TestUnmarshalYaml(t *testing.T) {
	//同理反序列化时，非可导出变量的数据会丢失变为空
	data := `name: 小强...'
			sex: 男
			tall: 1.7
			weight: 71
			age: 35`
	personalInformation := PersonalInformation{}
	yaml.Unmarshal([]byte(data), &personalInformation)
	fmt.Println(personalInformation)
}

func TestMarshalProtobuf(t *testing.T) {
	personalInformation := PersonalInformation{
		Name:   "小强...'",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}
	data, err := proto.Marshal(&personalInformation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
	fmt.Printf("%+v\n", string(data))
	//通常在非程序交互过程中，要保留原生protobuf，可以直接写入文件。如果想要单行保存，必须转码。
	//选择的通用转码是base64
	outPut64 := base64.StdEncoding.EncodeToString(data)
	fmt.Println(outPut64)
}

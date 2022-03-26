package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
	proto2 "learngo/pkg/apis"
	"log"
	"os"
)

func newRecord(filepath string) *record {
	return &record{
		filepath:      filepath,
		yamlFilePath:  filepath + ".yaml",
		protoFilePath: filepath + ".proto.base64",
	}
}

type record struct {
	filepath      string
	yamlFilePath  string
	protoFilePath string
}

func (r *record) savePersonalInformation(personInfo *proto2.PersonalInformation) error {
	{
		data, err := json.Marshal(personInfo)
		if err != nil {
			log.Println("marshal出错:", err)
			return err
		}
		if err = r.writeFileWithAppendJson(data); err != nil {
			log.Println("写入json时出错:", err)
			return err
		}
	}
	{
		data, err := yaml.Marshal(personInfo)
		if err != nil {
			log.Println("marshal出错:", err)
			return err
		}
		if err = r.writeFileWithAppendYaml(data); err != nil {
			log.Println("写入yaml时出错:", err)
			return err
		}
	}
	{
		data, err := proto.Marshal(personInfo)
		if err != nil {
			log.Println("marshal出错:", err)
			return err
		}
		if err = r.writeFileWithAppendProto(data); err != nil {
			log.Println("写入proto时出错:", err)
			return err
		}
	}
	return nil
}

func (r *record) writeFileWithAppendJson(data []byte) error {
	file, err := os.OpenFile(r.filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //0777最大权限,此方法为追加，无法修改原来的数据
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func (r *record) writeFileWithAppendYaml(data []byte) error {
	file, err := os.OpenFile(r.yamlFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //0777最大权限,此方法为追加，无法修改原来的数据
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	newData := append([]byte("---\n"), data...)
	_, err = file.Write(append(newData, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func (r *record) writeFileWithAppendProto(data []byte) error {
	file, err := os.OpenFile(r.protoFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //0777最大权限,此方法为追加，无法修改原来的数据
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	_, err = file.Write([]byte(base64.StdEncoding.EncodeToString(data)))
	if err != nil {
		return err
	}
	return nil
}

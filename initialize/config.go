package initialize

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"nav/config"
)

/*
拓展类读取yaml文件，构造获取配置文件 初始化 函数
**/

type ServerConf config.Server

func (c *ServerConf) getServerYamlConf() *ServerConf {
	fmt.Print("开始读取配置文件 \n")
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("读取 yaml 文件出错: %v \n", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("读取 yaml 文件内容出错: %v \n", err)
	}
	fmt.Print("读取配置文件成功 \n")
	return c
}

func InitConf() ServerConf {
	var c ServerConf
	c.getServerYamlConf()
	return c
}

// 函数读取yaml文件
//func getDbYamlConf() config.PostgresYaml{
//	c :=config.PostgresYaml{}
//	yamlFile, err := ioutil.ReadFile("config.yaml")
//	if err != nil {
//		log.Printf("yamlFile.Get err   #%v ", err)
//	}
//
//	err = yaml.Unmarshal(yamlFile, &c)
//	if err != nil {
//		log.Fatalf("Unmarshal: %v", err)
//	}
//	return c
//}
//

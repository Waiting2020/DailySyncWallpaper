package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type conf struct {
	LogPath  string `yaml:"LogPath"`
	ImgPath  string `yaml:"ImgPath"`
	UserName string `yaml:"UserName"`
}

// 配置
var config conf

// 获取本地配置
func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile(GetCurrentDirectory() + `\conf.yaml`)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

//持久化配置到本地
func (c *conf) setConf() *conf {
	yamlText, err := yaml.Marshal(c)
	if err != nil {
		log.Printf("Marshal conf.Get err   #%v ", err)
	}
	err = ioutil.WriteFile(GetCurrentDirectory()+`\conf.yaml`, yamlText, 666)
	if err != nil {
		log.Printf("Write conf into file.Get err   #%v ", err)
	}
	return c
}

//获取当前路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

package main

import (
	"os"
	"strings"

	"github.com/clbanning/mxj/v2"
	toml "github.com/pelletier/go-toml"
	yaml "gopkg.in/yaml.v3"
)

func main() {

	filePath := os.Args[1]

	toType := os.Args[2]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Stdout.WriteString("error")
	}

	fileType := strings.Split(filePath, ".")[1]

	var newMap mxj.Maps
	var err error
	switch fileType {
	case "json":
		newMap, err = mxj.NewMapsFromJsonFile(filePath)
	case "xml":
		newMap, err = mxj.NewMapsFromXmlFile(filePath)
	case "yaml":
		yamlFile, err := os.ReadFile(filePath)
		if err != nil {
			os.Stdout.WriteString("error")
		}
		obj := make(map[string]interface{})
		err = yaml.Unmarshal(yamlFile, obj)
		if err != nil {
			os.Stdout.WriteString("error")
		}
		theMap := mxj.Map(obj)
		newMap = append(newMap, theMap)
	case "toml":
		tomlFile, err := os.ReadFile(filePath)
		if err != nil {
			os.Stdout.WriteString("error")
		}
		obj := make(map[string]interface{})
		err = toml.Unmarshal(tomlFile, &obj)
		if err != nil {
			os.Stdout.WriteString("error")
		}
		theMap := mxj.Map(obj)
		newMap = append(newMap, theMap)
	default:
		os.Stdout.WriteString("error")
	}

	if err != nil {
		os.Stdout.WriteString("error")
	}

	var outText string
	var outErr error
	switch toType {
	case "json":
		outText, outErr = newMap.JsonString()
	case "xml":
		outText, outErr = newMap.XmlString()
	case "yaml":
		outByte, err := yaml.Marshal(newMap)
		if err != nil {
			os.Stdout.WriteString("error")
		}
		outText = string(outByte)
	case "toml":
		outByte, err := toml.Marshal(newMap[0])
		if err != nil {
			os.Stdout.WriteString("error")
		}
		outText = string(outByte)
	default:
		os.Stdout.WriteString("error")
	}

	if outErr != nil {
		os.Stdout.WriteString("error")
	}

	os.Stdout.WriteString(outText)

}

package main

import "fmt"

type MethodConfig struct {
	ProtoType int
}

type ServiceConfig struct {
	Methods map[string]MethodConfig
}

var CCPServiceConfig = ServiceConfig{
	Methods: map[string]MethodConfig{
		"GetAuth": {ProtoType: 20},
	},
}

func getMethodProtoType(config ServiceConfig, method string) (uint32, error) {
	if vaule, isPresent := config.Methods[method]; isPresent {
		fmt.Println(config.Methods[method])
		fmt.Println(vaule.ProtoType)
		return uint32(vaule.ProtoType), nil
	}
	return 1, nil
}

func main() {
	getMethodProtoType(CCPServiceConfig, "GetAuth")
	fmt.Println("Test git")
}

package core

import "github.com/pkg/errors"

type MethodConfig struct {
	ProtoType int
}

type ServiceConfig struct {
	Methods map[string]MethodConfig
}

var CCPServiceConfig = ServiceConfig{
	Methods: map[string]MethodConfig{
		"GetAuth": MethodConfig{20},
	},
}

func getMethodProtoType(config ServiceConfig, method string) (uint32, error) {
	if vaule, isPresent := config.Methods[method]; isPresent {
		return uint32(vaule.ProtoType), nil
	}
	logger.Warnf("method not found: %s", method)
	return 99999, errors.New("method not found")
}

func method2ProtoType(method string) (uint32, error) {
	return getMethodProtoType(CCPServiceConfig, method)
}

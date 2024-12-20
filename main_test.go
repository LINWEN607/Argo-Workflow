package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"reflect"
	"testing"
)

func TestNacosSetup(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NacosSetup()
		})
	}
}

func TestServerSetup(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ServerSetup()
		})
	}
}

func Test_getHostIp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHostIp(); got != tt.want {
				t.Errorf("getHostIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSpecifiedConfig(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getSpecifiedConfig(tt.args.c)
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		data    map[string]interface{}
		keyPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValue(tt.args.data, tt.args.keyPath); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_health(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			health(tt.args.c)
		})
	}
}

func Test_hello(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hello(tt.args.c)
		})
	}
}

func Test_registerServiceInstance(t *testing.T) {
	type args struct {
		nacosClient naming_client.INamingClient
		param       vo.RegisterInstanceParam
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerServiceInstance(tt.args.nacosClient, tt.args.param)
		})
	}
}

func Test_selectOneHealthyInstance(t *testing.T) {
	type args struct {
		client      naming_client.INamingClient
		serviceName string
	}
	tests := []struct {
		name         string
		args         args
		wantInstance *model.Instance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInstance := selectOneHealthyInstance(tt.args.client, tt.args.serviceName); !reflect.DeepEqual(gotInstance, tt.wantInstance) {
				t.Errorf("selectOneHealthyInstance() = %v, want %v", gotInstance, tt.wantInstance)
			}
		})
	}
}

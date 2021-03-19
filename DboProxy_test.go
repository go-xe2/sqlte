package sqlte

import (
	"testing"
)

type TestArg struct {
	Name string `json:"name"`
	Sex int8 `json:"sex"`
}

type ProxyTestDbo struct {
	QueryAge func() error
	QueryUsers func(userId string, userAge int) (string, error) `args:"userId,userArg"`
	QuerySex func(argParam TestArg) (string, error)
}

func TestDboProxy(t *testing.T)  {
	var proxyTestDbo = ProxyTestDbo{}
	DboProxy(&proxyTestDbo, TestProxyBuildFunc)
	if e := proxyTestDbo.QueryAge(); e != nil {
		t.Error(e)
	}
	if s, e := proxyTestDbo.QueryUsers("张三", 33); e != nil {
		t.Error(e)
	} else {
		t.Log("QueryUsers res:", s)
	}
	if s, e := proxyTestDbo.QuerySex(TestArg{
		Name: "李四",
		Sex:  2,
	}); e != nil {
		t.Error(e)
	} else {
		t.Log("QuerySex:", s)
	}
}


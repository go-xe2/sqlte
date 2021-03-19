package sqlte

import (
	"github.com/go-xe2/x/os/xlog"
	"testing"
)

type TestCaller struct {
}

func (it *TestCaller) Call (funcName string, params map[string]interface{}, outNum int) (result []interface{},err error) {
	xlog.Info("call funcName:", funcName, "params:", params, ", outNum:", outNum)
	if outNum > 0 {
		result = make([]interface{}, outNum - 1)
	} else {
		result = make([]interface{}, outNum)
	}
	return result, nil
}

func TestSqlCaller(t *testing.T)  {
	var proxyTestDbo = ProxyTestDbo{}
	DboProxy(&proxyTestDbo, NewSqlCallerProxy(&TestCaller{}))
	proxyTestDbo.QueryAge()
	proxyTestDbo.QuerySex(TestArg{
		Name: "3333333",
		Sex:  22,
	})
	proxyTestDbo.QueryUsers("ddtestd", 2)
}

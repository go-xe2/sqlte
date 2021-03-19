package sqlte

import (
	"fmt"
	"reflect"
)

type SqlCaller interface {
	Call (funcName string, params map[string]interface{}, outNum int) ([]interface{}, error)
}

func NewSqlCallerProxy(caller SqlCaller) func(funcField reflect.StructField, field reflect.Value) func(arg FunProxyArg) []reflect.Value  {
	return func (funcField reflect.StructField, field reflect.Value) func(arg FunProxyArg) []reflect.Value  {
		return func(arg FunProxyArg) []reflect.Value {
			funcName := funcField.Name
			params := make(map[string]interface{})
			for i, k := range arg.FunArgs {
				v := arg.Args[i]
				if v.Kind() == reflect.Ptr && v.IsNil() {
					continue
				}
				params[k.Name] = v.Interface()
			}
			var returns = make([]reflect.Value, 0)
			res, err := caller.Call(funcName, params, field.Type().NumOut())
			fmt.Println("call func:", funcName, ", params:", params)
			for _, r := range res {
				//if r == nil {
				//	returns = append(returns, reflect.Zero(reflect.TypeOf(&r)).Elem())
				//} else {
					returns = append(returns, reflect.ValueOf(r))
				//}
			}
			var e error
			if err != nil {
				returns = append(returns, reflect.ValueOf(&err).Elem())
			} else {
				returns = append(returns, reflect.Zero(reflect.TypeOf(&e).Elem()))
			}
			return returns
		}
	}
}

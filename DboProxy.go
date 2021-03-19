package sqlte

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/os/xlog"
	"reflect"
	"strings"
)

type FunArg struct {
	Name  string
	Index int
}


//代理数据
type FunProxyArg struct {
	FunArgs []FunArg
	FunArgsLen int
	Args []reflect.Value
	ArgsLen int
}

func (it FunProxyArg)New(tagArgs []FunArg,args []reflect.Value) FunProxyArg {
	return FunProxyArg{
		FunArgs:tagArgs,
		Args:args,
		FunArgsLen: len(tagArgs),
		ArgsLen: len(args),
	}
}

func DboProxy(dbo interface{}, buildFunc func(funcField reflect.StructField, field reflect.Value) func(arg FunProxyArg) []reflect.Value) {
	v := reflect.ValueOf(dbo)
	if v.Kind() != reflect.Ptr {
		panic("DboProxy代理对象参数dbo只能传入指针类型")
	}
	buildProxy(v, buildFunc)
}

func buildProxy(v reflect.Value, buildFunc func(funcField reflect.StructField, field reflect.Value) func(arg FunProxyArg) []reflect.Value) {
	for {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			break
		}
	}
	t := v.Type()
	et := t
	if et.Kind() == reflect.Ptr {
		et = et.Elem()
	}
	ptr := v
	var obj reflect.Value
	if ptr.Kind() == reflect.Ptr {
		obj = ptr.Elem()
	} else {
		obj = ptr
	}
	count := obj.NumField()
	for i := 0; i < count; i++ {
		f := obj.Field(i)
		ft := f.Type()
		sf := et.Field(i)
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
		if f.CanSet() {
			switch ft.Kind() {
			case reflect.Struct:
				if buildFunc != nil {
					buildProxy(f, buildFunc) //循环扫描字段
				}
			case reflect.Func:
				if buildFunc != nil {
					bindProxyFunc(v, f, ft, sf, buildFunc(sf, f))
				}
			}
		}
	}
	if t.Kind() == reflect.Ptr {
		v.Set(ptr)
	} else {
		v.Set(obj)
	}
}

// bindProxyFunc 绑定代理方法
func bindProxyFunc(source reflect.Value, f reflect.Value, ft reflect.Type, sf reflect.StructField, proxyFunc func(arg FunProxyArg) []reflect.Value) {
	// 获取方法参数，创建代理参数列表
	var tagParams []string
	var args = sf.Tag.Get(`args`)
	argLen := ft.NumIn()
	var tagArgs = make([]FunArg, 0, argLen)
	if args != `` {
		tagParams = strings.Split(args, `,`)
		var tagParamsLen = len(tagParams)
		if tagParamsLen != ft.NumIn() {
			panic(fmt.Errorf("函当字段%s tags设置与输入参数不一致，期望%d实际长度%d",sf.Name, argLen, tagParamsLen))
		}
		if tagParamsLen != 0 {
			for index, v := range tagParams {
				var tagArg = FunArg{
					Index: index,
					Name:  v,
				}
				tagArgs = append(tagArgs, tagArg)
			}
		}
	} else {
		for i := 0; i < argLen; i++ {
			tagArgs = append(tagArgs, FunArg{
				Name:  fmt.Sprintf("arg%d", i),
				Index: i,
			})
		}
	}
	var tagArgsLen = len(tagArgs)
	if tagArgsLen > 0 && argLen != tagArgsLen {
		panic(fmt.Errorf("函当字段%s tags设置与输入参数不一致，期望%d实际长度%d",sf.Name, argLen, tagArgsLen))
	}
	var fn = func(args []reflect.Value) (results []reflect.Value) {
		// 调用代理函数处理返回结果
		proxyResults := proxyFunc(FunProxyArg{}.New(tagArgs, args))
		for _, returnV := range proxyResults {
			results = append(results, returnV)
		}
		return results
	}
	if f.Kind() == reflect.Ptr {
		fp := reflect.New(ft)
		fp.Elem().Set(reflect.MakeFunc(ft, fn))
		f.Set(fp)
	} else {
		f.Set(reflect.MakeFunc(ft, fn))
	}
}

func TestProxyBuildFunc(funcField reflect.StructField, field reflect.Value) func(arg FunProxyArg) []reflect.Value {
	return func(arg FunProxyArg) []reflect.Value {
		fmt.Println("====>> do func name:", funcField.Name)
		xlog.Info("传入参数名:", arg.FunArgs, "参数个数:", arg.FunArgsLen)
		for i := 0; i < arg.ArgsLen; i++ {
			fmt.Println(arg.FunArgs[i].Name, "=", arg.Args[i].Interface())
		}
		var e error
		fmt.Println("funcField.Type.NumOut:", field.Type().NumOut())
		var returns = make([]reflect.Value, 0)
		if len(arg.Args) == 0 {
			e = errors.New("参数个数不能为0")
		} else {
			returns = append(returns, reflect.ValueOf(fmt.Sprintf("返回第1个参数值=%v",arg.Args[0].Interface())))
		}
		if e == nil {
			returns = append(returns, reflect.Zero(reflect.TypeOf(&e).Elem()))
		} else {
			returns = append(returns, reflect.ValueOf(&e).Elem())
		}
		return returns
	}
}
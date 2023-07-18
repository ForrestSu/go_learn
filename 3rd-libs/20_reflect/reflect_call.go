package main

import (
	"context"
	"fmt"
	"reflect"

	"bou.ke/monkey"
	"github.com/ForrestSu/go_learn/3rd-libs/20_reflect/dao"
)

// just fake a UnExported struct: *dao.student
func WrapperInvoke(target reflect.Type, methodName string, innerFunc interface{}) (i interface{}) {
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic(fmt.Sprintf("unknown method %s", methodName))
	}
	wrapper := reflect.MakeFunc(m.Func.Type(), func(in []reflect.Value) (results []reflect.Value) {
		// fmt.Printf("wrapper fn ....")
		return reflect.ValueOf(innerFunc).Call(in[1:])

		// if len(in) != m.Func.Type().NumIn() {
		//	panic(fmt.Sprintf("Arguments len(in) = %d is invalid!", len(in)))
		// }
		// ctx, _ := in[1].Interface().(context.Context)
		// err := innerFn(ctx, in[2].Interface(), in[3].Interface())
		// //
		// return []reflect.Value{reflect.ValueOf(err)}
	})
	return wrapper.Interface()
}

func mai1n() {

	var mockFunc = func(ctx context.Context, req interface{}, rsp interface{}, _ []int) error {
		fmt.Printf("monkey patch ok...\n")
		return nil
	}
	methodName := "Invoke"
	clientType := reflect.TypeOf(dao.New())
	wrapFunc := WrapperInvoke(clientType, methodName, mockFunc)
	monkey.PatchInstanceMethod(clientType, methodName, wrapFunc)

	req := NewRequest()
	rsp := NewResponse()
	fmt.Println("invoke begin...")
	err := dao.New().Invoke(context.Background(), req, rsp, 1)
	if err == nil {
		fmt.Println("return nil.")
	} else {
		fmt.Printf("error! ret err = %+v", err)
	}
}

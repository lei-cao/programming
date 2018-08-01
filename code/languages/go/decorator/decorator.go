package decorator

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type FooFunc func(int, int, int) int
type BarFunc func(a, b string) string

func foo(a, b, c int) int {
	return a + b*c
}

func bar(a, b string) string {
	return "hello " + a + " and " + b
}

func MyDoubleUpDecorator(f func(int, int, int) int) func(int, int, int) int {
	return func(a, b, c int) int {
		fmt.Println("Decoracting...")
		return f(a, b, c) * 2
	}
}

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func LogTimeDecorator(f FooFunc) FooFunc {
	return func(a, b, c int) int {
		defer func(start time.Time) {
			fmt.Printf("it takes %v for doing %s \n", time.Since(start), getFuncName(f))
		}(time.Now())
		return f(a, b, c)
	}
}

func GenricDecorator(decoPtr, fn interface{}) error {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()

	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(), func(in []reflect.Value) (out []reflect.Value) {
		defer func(start time.Time) {
			fmt.Printf("it takes %v for doing %s \n", time.Since(start), getFuncName(fn))
		}(time.Now())
		out = targetFunc.Call(in)
		return
	})
	decoratedFunc.Set(v)
	return nil
}

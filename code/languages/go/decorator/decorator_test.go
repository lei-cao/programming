package decorator

import (
	"testing"
	"fmt"
)

func TestGenricDecorator(t *testing.T) {
	var barFun BarFunc
	GenricDecorator(&barFun, bar)
	fmt.Println(barFun("lion", "tiger"))

	var fooFun FooFunc
	GenricDecorator(&fooFun, foo)
	fmt.Println(fooFun(1, 3, 4))
}

func TestMyDoubleUpDecorator(t *testing.T) {
	fmt.Println(MyDoubleUpDecorator(foo)(1, 2, 3))
}
func TestLogTimeDecorator(t *testing.T) {
	logTimeDeco := LogTimeDecorator
	fmt.Println(logTimeDeco(foo)(1, 2, 3))
}

package main

import "fmt"
import "errors"

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为０")
	} else {
		result = a / b
	}
	return

}

func main() {
	//var err1 error = fmt.Errorf("%s ", "this is normal error")
	err1 := fmt.Errorf("%s ", "this is normal error")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2 = ", err2)

	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("result = ", result)
	}
}

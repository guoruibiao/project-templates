package main

import (
	"fmt"
	"os"
)

func main() {
	description := `
	    代码中不可避免会出现一些账号密码之类的数据，不便于放到代码库。
		因此最好是可以放到ENV中，方式：
		export variable-name=variable-value
	`
	variable1 := os.Getenv("VARIABLE1")
	variable2 := os.Getenv("VARIABLE2")
	fmt.Println("variable1=", variable1, ", variable2=", variable2)
}

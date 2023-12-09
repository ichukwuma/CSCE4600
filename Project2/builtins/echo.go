package builtins

import (
	"fmt"
	"strings"
//	"os"
)

func Echo(args ...string) {

//	if len(args) == 1 && args[0] == "$PATH" {
//		path := os.Getenv("PATH")
//		fmt.Println(path)
//	} else {
//		echoOut := strings.Join(args, " ")
//		fmt.Println(echoOut)
//	}
//
//}
	echoOut := strings.Join(args, " ")
	fmt.Println(echoOut)
//	fmt.Println(args)

	//for path
//	path := os.Getenv("PATH")
	//fmt.Println(path)
}

package builtins

import (
	"fmt"
	"os"
)

func Bye() {

	fmt.Println("Session Ended, good bye")
	os.Exit(0)


}

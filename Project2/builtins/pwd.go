package builtins
import (
	"fmt"
	"os"
)


func Pwd() {

	mydirectory, check4error := os.Getwd()

	if check4error != nil {
		fmt.Println(check4error)
		return

	}
	fmt.Println(mydirectory)

}

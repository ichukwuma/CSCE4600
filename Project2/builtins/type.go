package builtins

import(
	"os"
	"fmt"
	"os/exec"
)

	

func Type() {


	typecmd := os.Args[1]

	if _, typeerror := exec.LookPath(typecmd); typeerror == nil {
		fmt.Printf("%s is an external command\n", typecmd)
	} else {
		fmt.Printf("%s is not a shell built-in command\n", typecmd)
	}

}


package builtins

import (
	"time"
	"fmt"
	"os/exec"
	"os"
)

func Time() {
	timecmd := os.Args[1:]

	begin := time.Now()

	commandtracked := exec.Command(timecmd[0], timecmd[1:]...)
	//commandtracked.Stderr = os.Stderr
	//commandtracked.Stdout = os.Stdout

	lsoutput, _  := commandtracked.CombinedOutput()

	end := time.Now()


	timespent := end.Sub(begin)
	fmt.Printf("Total: %s\n", timespent)

	fmt.Printf(string(lsoutput))

}


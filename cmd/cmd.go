package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func CheckNTPService() bool {
	out, err := exec.Command("ps", "-ef", "| grep", "firefox").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HHHHHHHH:%s", out)
	return true
}
func DateCommand() string {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", out)
}

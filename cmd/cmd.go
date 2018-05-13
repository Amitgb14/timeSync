package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func DateCommand() string {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", out)
}

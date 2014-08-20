package git

import (
	"log"
	"os/exec"
)

func Exec(dir string, args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	err := cmd.Run()

	if err != nil {
		log.Fatal(args, err)
	}
}

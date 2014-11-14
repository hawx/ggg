package git

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadFile(path, file string) (string, error) {
	masterRef := run(path, "rev-parse", "master")
	if masterRef == "" {
		return "", errors.New("No such branch: master")
	}

	ref := ""
	for _, row := range strings.Split(run(path, "ls-tree", masterRef), "\n") {
		cols := strings.Fields(row)

		if cols[3] == file && cols[1] == "blob" {
			ref = cols[2]
			break
		}
	}

	if ref == "" {
		return "", errors.New("No such file: " + file)
	}

	return run(path, "cat-file", "blob", ref), nil
}

func CreateRepo(path string) {
	os.Mkdir(path, 0755)
	run(path, "init", "--bare")

	sampleHook := filepath.Join(path, "hooks", "post-update.sample")
	hook := filepath.Join(path, "hooks", "post-update") // need to replace with hook that calls ggg!
	os.Rename(sampleHook, hook)
	run(path, "update-server-info")
}

func run(dir string, args ...string) string {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}

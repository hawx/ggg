package git

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Branches(path string) []string {
	branches := []string{}
	for _, branch := range strings.Split(run(path, "branch"), "\n") {
		if len(branch) > 0 {
			branches = append(branches, branch[2:])
		}
	}
	return branches
}

func GetDefaultBranch(path string) string {
	for _, branch := range strings.Split(run(path, "branch"), "\n") {
		if len(branch) > 0 && branch[0] == 42 {
			return branch[2:]
		}
	}

	return ""
}

func GetBranch(path, branch string) string {
	return run(path, "rev-parse", branch)
}

func ReadFile(path, branch, file string) (string, error) {
	branchRef := GetBranch(path, branch)

	ref := ""
	for _, row := range strings.Split(run(path, "ls-tree", branchRef), "\n") {
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
	os.Mkdir(path, 0775)
	run(path, "init", "--bare", "--shared=group")

	sampleHook := filepath.Join(path, "hooks", "post-update.sample")
	hook := filepath.Join(path, "hooks", "post-update") // need to replace with hook that calls ggg!
	os.Rename(sampleHook, hook)
	run(path, "update-server-info")
}

func CopyRepo(path, remoteUrl string) {
	log.Println("git clone --bare", remoteUrl, path)
	run("", "clone", "--bare", remoteUrl, path)

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

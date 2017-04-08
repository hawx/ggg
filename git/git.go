package git

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	goGit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/filemode"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func Branches(path string) (branches []string) {
	r, err := goGit.PlainOpen(path)
	if err != nil {
		return branches
	}

	refs, err := r.References()
	if err != nil {
		return branches
	}

	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.IsBranch() {
			branches = append(branches, ref.Name().Short())
		}

		return nil
	})

	return branches
}

type File struct {
	Name        string
	Path        string
	IsDir       bool
	IsSubmodule bool
}

func Files(path, branch, root string) (files []File, err error) {
	r, err := goGit.PlainOpen(path)
	if err != nil {
		return files, err
	}

	branchRef, _ := r.ResolveRevision(plumbing.Revision("refs/heads/" + branch))

	c, err := r.Commit(*branchRef)
	if err != nil {
		return files, err
	}

	tree, err := c.Tree()
	if err != nil {
		return files, err
	}

	if root != "" {
		tree, err = tree.Tree(root)
		if err != nil {
			return files, err
		}
	}

	walker := object.NewTreeWalker(tree, false)
	defer walker.Close()

	for {
		name, entry, err := walker.Next()
		if err != nil {
			return files, err
		}

		path := name
		if root != "" {
			path = root + "/" + name
		}

		switch entry.Mode {
		case filemode.Submodule:
			files = append(files, File{Name: name, Path: path, IsSubmodule: true})
		case filemode.Dir:
			files = append(files, File{Name: name, Path: path, IsDir: true})
		default:
			files = append(files, File{Name: name, Path: path})
		}
	}

	return files, nil
}

func GetDefaultBranch(path string) string {
	for _, branch := range strings.Split(run(path, "branch"), "\n") {
		if len(branch) > 0 && branch[0] == 42 {
			return branch[2:]
		}
	}

	return ""
}

func ReadFile(path, branch, file string) (string, error) {
	r, err := goGit.PlainOpen(path)
	if err != nil {
		return "", err
	}

	branchRef, _ := r.ResolveRevision(plumbing.Revision("refs/heads/" + branch))

	c, err := r.Commit(*branchRef)
	if err != nil {
		return "", err
	}

	f, err := c.File(file)
	if err != nil {
		return "", err
	}

	return f.Contents()
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

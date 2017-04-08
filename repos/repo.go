package repos

import (
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/shurcooL/github_flavored_markdown"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/filemode"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func Branches(path string) (branches []string) {
	r, err := git.PlainOpen(path)
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
	r, err := git.PlainOpen(path)
	if err != nil {
		return files, err
	}

	branchRef, err := r.ResolveRevision(plumbing.Revision("refs/heads/" + branch))
	if err != nil {
		return files, err
	}

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
	r, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	branchRef, err := r.ResolveRevision(plumbing.Revision("refs/heads/" + branch))
	if err != nil {
		return "", err
	}

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

type Repo struct {
	Name        string
	Web         string
	Description string
	Path        string
	Branch      string
	LastUpdate  time.Time
	IsPrivate   bool
}

func (r *Repo) CloneUrl() string {
	return r.Name + ".git"
}

func (r *Repo) Branches() []string {
	return Branches(r.Path)
}

func (r *Repo) Files(tree string) []File {
	files, _ := Files(r.Path, r.DefaultBranch(), tree)

	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir && !files[j].IsDir {
			return true
		}
		if !files[i].IsDir && files[j].IsDir {
			return false
		}

		return strings.ToLower(files[i].Name) < strings.ToLower(files[j].Name)
	})

	return files
}

func (r *Repo) IsEmpty() bool {
	return len(r.Branches()) == 0
}

func (r *Repo) Contents(file string) (string, error) {
	return ReadFile(r.Path, r.DefaultBranch(), file)
}

func (r *Repo) Readme() (name string, contents template.HTML) {
	branch := r.DefaultBranch()

	for _, file := range []string{"README.md", "Readme.md", "README.markdown", "readme.markdown"} {
		text, err := ReadFile(r.Path, branch, file)
		if err != nil {
			continue
		}

		return file, template.HTML(github_flavored_markdown.Markdown([]byte(text)))
	}

	for _, file := range []string{"README"} {
		text, err := ReadFile(r.Path, branch, file)
		if err != nil {
			continue
		}

		return file, template.HTML("<pre class='full'>" + text + "</pre>")
	}

	return "README", template.HTML("&hellip;")
}

func (r *Repo) DefaultBranch() string {
	if r.Branch == "" {
		return GetDefaultBranch(r.Path)
	}

	return r.Branch
}

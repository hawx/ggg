package repos

import (
	"html/template"
	"sort"
	"strings"
	"time"

	"github.com/shurcooL/github_flavored_markdown"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/filemode"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type GitRepo struct {
	repo *git.Repository
}

func OpenRepo(path string) (*GitRepo, error) {
	r, err := git.PlainOpen(path)

	return &GitRepo{repo: r}, err
}

func (g *GitRepo) Branches() (branches []string) {
	refs, err := g.repo.References()
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

func (g *GitRepo) Files(branch, root string) (files []File, err error) {
	branchRef, err := g.repo.ResolveRevision(plumbing.Revision("refs/heads/" + branch))
	if err != nil {
		return files, err
	}

	c, err := g.repo.Commit(*branchRef)
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

func (g *GitRepo) ReadFile(branch, file string) (string, error) {
	branchRef, err := g.repo.ResolveRevision(plumbing.Revision("refs/heads/" + branch))
	if err != nil {
		return "", err
	}

	c, err := g.repo.Commit(*branchRef)
	if err != nil {
		return "", err
	}

	f, err := c.File(file)
	if err != nil {
		return "", err
	}

	return f.Contents()
}

type File struct {
	Name        string
	Path        string
	IsDir       bool
	IsSubmodule bool
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

func (r *Repo) Branches() (branches []string) {
	g, err := OpenRepo(r.Path)
	if err != nil {
		return branches
	}

	return g.Branches()
}

func (r *Repo) Files(tree string) (files []File) {
	g, err := OpenRepo(r.Path)
	if err != nil {
		return files
	}

	files, _ = g.Files(r.DefaultBranch(), tree)

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
	g, err := OpenRepo(r.Path)
	if err != nil {
		return "", err
	}

	return g.ReadFile(r.DefaultBranch(), file)
}

func (r *Repo) Readme() (name string, contents template.HTML) {
	g, err := OpenRepo(r.Path)
	if err != nil {
		return "", template.HTML("no such repo")
	}

	branch := r.DefaultBranch()

	for _, file := range []string{"README.md", "Readme.md", "README.markdown", "readme.markdown"} {
		text, err := g.ReadFile(branch, file)
		if err != nil {
			continue
		}

		return file, template.HTML(github_flavored_markdown.Markdown([]byte(text)))
	}

	for _, file := range []string{"README"} {
		text, err := g.ReadFile(branch, file)
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

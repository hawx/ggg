package repos

import (
	"github.com/hawx/ggg/git"

	"github.com/boltdb/bolt"
	"github.com/shurcooL/go/github_flavored_markdown"

	"encoding/json"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"
	"sort"
)

type ByName []Repo

func (s ByName) Len() int {
	return len(s)
}

func (s ByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s ByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Repo struct {
	Name        string
	Web         string
	Description string
	Path        string
	LastUpdate  time.Time
	IsPrivate   bool
}

func (r Repo) CloneUrl() string {
	return r.Name + ".git"
}

func (r Repo) Readme() template.HTML {
	for _, file := range []string{"README.md", "Readme.md", "README.markdown"} {
		text, err := git.ReadFile(r.Path, file)
		if err != nil {
			log.Println(r.Name, " Readme(): ", err)
			continue
		}

		return template.HTML(github_flavored_markdown.Markdown([]byte(text)))
	}

	return "&hellip;"
}

type Db interface {
	GetAll() []Repo
	Get(name string) Repo
	Create(name, web, description string, isPrivate bool)
	Save(Repo)
	Delete(Repo)
	Close()
}

type BoltDb struct {
	me     *bolt.DB
	gitDir string
}

const bucketName = "repos"

func Open(path, gitDir string) Db {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		return err
	})

	return BoltDb{db, gitDir}
}

func (db BoltDb) Create(name, web, description string, isPrivate bool) {
	path := filepath.Join(db.gitDir, name) + ".git"
	repo := Repo{name, web, description, path, time.Now(), isPrivate}

	db.Save(repo)
}

func (db BoltDb) GetAll() []Repo {
	list := []Repo{}

	db.me.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var repo Repo
			json.Unmarshal(v, &repo)
			list = append(list, repo)
		}

		return nil
	})

	sort.Sort(ByName(list))
	return list
}

func (db BoltDb) Get(name string) Repo {
	var repo Repo

	db.me.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(name))
		if v != nil {
			json.Unmarshal(v, &repo)
		}
		return nil
	})

	return repo
}

func (db BoltDb) Save(repo Repo) {
	key := repo.Name
	serialised, _ := json.Marshal(repo)

	db.me.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Put([]byte(key), serialised)
	})
}

func (db BoltDb) Delete(repo Repo) {
	db.me.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete([]byte(repo.Name))
	})

	os.RemoveAll(repo.Path)
}

func (db BoltDb) Close() {
	db.me.Close()
}

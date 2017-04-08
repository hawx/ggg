package repos

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/boltdb/bolt"
)

type Db interface {
	GetAll() []*Repo
	Get(name string) *Repo
	Create(name, web, description, branch string, isPrivate bool)
	Save(*Repo)
	Delete(*Repo)
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

func (db BoltDb) Create(name, web, description, branch string, isPrivate bool) {
	path := filepath.Join(db.gitDir, name) + ".git"
	repo := &Repo{
		Name:        name,
		Web:         web,
		Description: description,
		Path:        path,
		Branch:      branch,
		LastUpdate:  time.Now(),
		IsPrivate:   isPrivate,
	}

	db.Save(repo)
}

func (db BoltDb) GetAll() []*Repo {
	list := []*Repo{}

	db.me.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var repo Repo
			json.Unmarshal(v, &repo)
			list = append(list, &repo)
		}

		return nil
	})

	sort.Sort(ByName(list))
	return list
}

func (db BoltDb) Get(name string) *Repo {
	var repo Repo

	db.me.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(name))
		if v != nil {
			json.Unmarshal(v, &repo)
		}
		return nil
	})

	return &repo
}

func (db BoltDb) Save(repo *Repo) {
	key := repo.Name
	serialised, _ := json.Marshal(repo)

	db.me.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Put([]byte(key), serialised)
	})
}

func (db BoltDb) Delete(repo *Repo) {
	db.me.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete([]byte(repo.Name))
	})

	os.RemoveAll(repo.Path)
}

func (db BoltDb) Close() {
	db.me.Close()
}

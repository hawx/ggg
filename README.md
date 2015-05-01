# ggg

App for serving a list of git repos.

``` bash
$ go get hawx.me/code/ggg
$ echo "user = me@domain.com" > settings.toml
$ ggg
Running on port :8080
...
```

Features:

- Public repos over http
- Create / edit / delete repos
- Private repos (not over http)
- Import from GitHub (`go get hawx.me/code/ggg/cmd/ggg-import-github`)

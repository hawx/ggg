# ggg

App for serving a list of git repos.

``` bash
$ go get github.com/hawx/ggg
$ echo "user = me@domain.com" > settings.toml
$ ggg
Running on port :8080
...
```

Features:

- Public repos over http
- Create / edit / delete repos
- Private repos (not over http)

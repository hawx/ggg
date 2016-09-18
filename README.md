# ggg

App for serving a list of git repos.

``` bash
$ go get hawx.me/code/ggg
$ cat settings.toml
title = "my example git server"
url = "https://example.com"
secret = "output of `head -c32 /dev/urandom | openssl base64`"
gitDir = "/path/to/git/dir"
dbPath = "/path/to/db"

[uberich]
appName = "ggg"
appURL = "https://example.com"
uberichURL = "https://login.example.com"
secret = "our shared secret"
$ ggg
Running on port :8080
...
```

Features:

- Public repos over http
- Create / edit / delete repos
- Private repos (not over http)
- Import from GitHub (`go get hawx.me/code/ggg/cmd/ggg-import-github`)

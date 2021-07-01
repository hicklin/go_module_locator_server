# Go Module Locator

This is a very simple web server used to provide `go mod` with the location of your private Go module. This
is useful when your private Go modules are accessed via non-standard ports.

## How to use

1.  Build `main.go` with `go build -o locatorServer main.go`.
2.  Execute `locatorServer` with the site `hostname`, access `port` and `repoLocation`.
    *  The `hostname` needs to be the same as in your go module.
    *  If you are running the server locally, you can point the hostname to you localhost ip by editing `/etc/hosts`.
    *  The `repoLocation` is where you will be pulling your Go module from via git. Example, if all your Go module repos 
       are named `go_module_<module_name>`, you can set this to `ssh://git@company.com/team/go_module_`.
3.  If you use port 80, i.e. `http`, you will need to set the `GOINSECURE` environment variable for your domain. Ex `GOINSECURE=mycooldomain.com`.
4.  Add this domain to the `GOPRIVATE` environment variable to avoid sum checks.

## How it works

Once running, when go wants to access one of your private Go modules, example

```go
import (
	myModule "mycooldomain.com/my_module"
)
```

go will send a GET request to `https://mycooldomain.com/my_module` and the server will tell go mod to git pull the repo 
from `<repoLocation><module_name>`, i.e. `ssh://git@company.com/team/go_module_my_module`.

[further info](https://golang.org/ref/mod#serving-from-proxy)

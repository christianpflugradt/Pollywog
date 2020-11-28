# Pollywog

**work in progress - this readme will be filled once pollywog reaches MVP status**

## dependencies ##

Pollywog depends on the following libraries (install them via go get)
* github.com/go-sql-driver/mysql
* github.com/mattn/go-sqlite3

## run from source ##

 * add codebase to your gopath, you can do this via a symlink (or a shortcut, depending on your OS)
    * `cd ~/go/src`
    * `ln -s /path-to-pollywog-project/src pollywog`
 * `go run pollywog.go`

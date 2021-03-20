cd src # workdir
go mod tidy # download/update dependencies
go build pollywog.go # build pollywog binary
chmod +x pollywog # make the binary executable

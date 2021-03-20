git checkout $1
echo "package model; const Version = \"`git describe --tags | cut -c1-5`\"" > src/domain/model/version.go # write git tag version into source code

#!/bin/bash
#
# usage:
#
# $ sudo ./golang-crosscompile-setup.bash


goroot=$(go env GOROOT)
platforms="darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm windows/386 windows/amd64"
programs=$(ls -l cmd/ | grep ^d | awk '{print $9}')

for program in ${programs}
do
    for platform in ${platforms}
    do
        split=(${platform//\// })
        goos=${split[0]}
        goarch=${split[1]}
        name=$program'-'$goos'-'$goarch
        if [ $goos = "windows" ]; then
            name+='.exe'
        fi
        env GOOS=$goos GOARCH=$goarch go build -o $name ./cmd/$program
    done
done
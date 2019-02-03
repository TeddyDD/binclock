#!/bin/sh
# Generates cross builds for all supported platforms.
#
# This script is used to build binaries for all supported platforms. Cgo is
# disabled to make sure binaries are statically linked. Appropriate flags are
# given to the go compiler to strip binaries. Current git tag is passed to the
# compiler by default to be used as the version in binaries. These are then
# compressed in an archive form (`.zip` for windows and `.tar.gz` for the rest)
# within a folder named `dist`.

set -o verbose

[ -z $version ] && version=$(git describe --tags)

mkdir -p dist

CGO_ENABLED=0 GOOS=darwin    GOARCH=386      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-darwin-386.tar.gz      binclock --remove-files
CGO_ENABLED=0 GOOS=darwin    GOARCH=amd64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-darwin-amd64.tar.gz    binclock --remove-files
CGO_ENABLED=0 GOOS=dragonfly GOARCH=amd64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-dragonfly-amd64.tar.gz binclock --remove-files
CGO_ENABLED=0 GOOS=freebsd   GOARCH=386      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-freebsd-386.tar.gz     binclock --remove-files
CGO_ENABLED=0 GOOS=freebsd   GOARCH=amd64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-freebsd-amd64.tar.gz   binclock --remove-files
CGO_ENABLED=0 GOOS=freebsd   GOARCH=arm      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-freebsd-arm.tar.gz     binclock --remove-files
CGO_ENABLED=0 GOOS=linux     GOARCH=386      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-linux-386.tar.gz       binclock --remove-files
CGO_ENABLED=0 GOOS=linux     GOARCH=amd64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-linux-amd64.tar.gz     binclock --remove-files
CGO_ENABLED=0 GOOS=linux     GOARCH=arm      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-linux-arm.tar.gz       binclock --remove-files
CGO_ENABLED=0 GOOS=linux     GOARCH=arm64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-linux-arm64.tar.gz     binclock --remove-files
CGO_ENABLED=0 GOOS=linux     GOARCH=ppc64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-linux-ppc64.tar.gz     binclock --remove-files
CGO_ENABLED=0 GOOS=linux     GOARCH=ppc64le  go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-linux-ppc64le.tar.gz   binclock --remove-files
CGO_ENABLED=0 GOOS=netbsd    GOARCH=386      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-netbsd-386.tar.gz      binclock --remove-files
CGO_ENABLED=0 GOOS=netbsd    GOARCH=amd64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-netbsd-amd64.tar.gz    binclock --remove-files
CGO_ENABLED=0 GOOS=netbsd    GOARCH=arm      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-netbsd-arm.tar.gz      binclock --remove-files
CGO_ENABLED=0 GOOS=openbsd   GOARCH=386      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-openbsd-386.tar.gz     binclock --remove-files
CGO_ENABLED=0 GOOS=openbsd   GOARCH=amd64    go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-openbsd-amd64.tar.gz   binclock --remove-files
CGO_ENABLED=0 GOOS=openbsd   GOARCH=arm      go build -ldflags="-s -w" && sync && tar czf dist/binclock-$version-openbsd-arm.tar.gz     binclock --remove-files

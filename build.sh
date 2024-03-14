#! /bin/bash

binName=latexresume
builds=dist

# Init values
arch=amd64
os=linux

# Get the software version from tag
version=$(git describe --tags --abbrev=0)

# Build LINUX_AMD64
GOARCH=$arch GOOS=$os go build -ldflags "-X main.version=$version"  -o $binName .

tarName=$binName\_$os\_$arch.tar.gz

tar -czvf $tarName $binName

mv $tarName $builds/$tarName && rm $binName

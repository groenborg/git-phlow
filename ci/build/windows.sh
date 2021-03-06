#!/bin/sh
set -e -u -x

apt-get update
apt-get install zip unzip

#GET VERSION AND PATH
#STORE VERSION AND PATH FOR BUILD
VERSION=$(cat gp-version/version)
BUILDPATH=$(pwd)

#CREATE GO DIRECTORY STRUCTURE
#THE STRUCTURE IS NECESSARY FOR GO TOOLS OTHERWISE
# 'build' AND 'get' WILL FAIL
mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma

#NAVIGATE TO FOLDER AND GET DEPS
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -t -v ./...
go get -v github.com/inconshreveable/mousetrap

#BUILD WE ONLY BUILD FOR amd64
export GOARCH=amd64

#WINDOWS
#BUILD AND COMPRESS
export GOOS=windows
go build -ldflags "-X   github.com/praqma/git-phlow/options.Version=`echo $VERSION` -X  github.com/praqma/git-phlow/options.Sha1=`git rev-parse HEAD` -X  github.com/praqma/git-phlow/options.Date=`date +'%d-%m-%Y'`"
zip $BUILDPATH/build-artifacts/git-phlow-$VERSION-windows-$GOARCH.zip git-phlow.exe
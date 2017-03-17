#!/bin/sh
set -e -u -x

#LIST RESOURCES IN ENTRY FOLDER
#FOR DEBUGGING PURPOSES
ls -lah
ls phlow-artifact-darwin-s3/
#GET VERSION AND PATH
#STORE VERSION AND PATH FOR BUILD
VERSION=$(cat resource-version/version)
SHA=$(sha256sum phlow-artifact-darwin-s3/git-phlow-$VERSION-darwin-amd64.tar.gz)
URL=https://github.com/Praqma/git-phlow/releases/download/v$VERSION/git-phlow-$VERSION-darwin-amd64.tar.gz
BUILDPATH=$(pwd)

#get stuffe
wget https://github.com/gliderlabs/sigil/releases/download/v0.4.0/sigil_0.4.0_Darwin_x86_64.tgz | tar -zxC /usr/local/bin


#Install sigil first

sigil -f git-phlow/ci/scripts/formula.tp sha$SHA= version=$VERSION url=$URL

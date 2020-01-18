#!/bin/bash

set -e

if [ "x$DEBUG" != "x" ] ; then
  set -x
fi

function fatal() {
  echo $* >&2
  exit 1
}

function ident() {
  sed -e 's/^/\t /'
}

if [ $# -ne 1 ]; then
  fatal "Usage: $0 [VERSION]" >&2
fi

VERSION=$1

BASE_DIR=$(pwd)
mkdir dist

echo "[+] Releasing version $VERSION"

build_dir=$(mktemp -d "/tmp/acp-install.XXXXXXXXXX")
pushd $build_dir > /dev/null

echo "[+] Downloading front"
curl --fail -L https://github.com/a-contre-plongee/acp-front/archive/v$VERSION.tar.gz -s -o front.tar.gz || fatal "Fail to download front version $VERSION"

echo "[+] Downloading back"
curl --fail -L https://github.com/johnsudaar/acp/archive/v$VERSION.tar.gz -s -o back.tar.gz || fatal "Fail to download back version $VERSION"

echo "[+] Unpacking front"
tar -xvf front.tar.gz >/dev/null || fatal "Fail to unpack front"

echo "[+] Unpacking back"
tar -xvf back.tar.gz >/dev/null || fatal "Fail to unpack back"

pushd acp-front-$VERSION > /dev/null
echo "[+] Front: Installing dependencies"
yarn install 2>&1 | ident
echo "[+] Front: Building"
yarn build 2>&1 | ident

popd
pushd acp-$VERSION > /dev/null

echo "[+] Back: Installing gox"
go get -u github.com/mitchellh/gox
echo "[+] Back: Building"
gox -os="linux" -arch="amd64 386 arm" -output="dist/acp-$VERSION-{{.OS}}-{{.Arch}}/acp" -ldflags="-X main.Version=$VERSION" .
pushd dist > /dev/null

for dir in * ; do
  echo "[+] Packaging front for $dir"
  cp -R $build_dir/acp-front-$VERSION/dist $dir/public
  zip -r $dir.zip $dir > /dev/null
  cp $dir.zip $BASE_DIR/dist/
done
popd
popd

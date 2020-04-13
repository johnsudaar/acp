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

VERSION=$(echo $1 | sed s/v//)

export VERSION=$VERSION
make pkg VERSION=$VERSION

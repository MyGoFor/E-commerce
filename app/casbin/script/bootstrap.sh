#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/casbin"
exec "$CURDIR/bin/casbin"

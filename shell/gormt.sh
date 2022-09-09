#!/bin/sh

# gormt方式
SHELL_FOLDER=$(
  cd "$(dirname $(dirname "$0"))"
  pwd
)
cd $SHELL_FOLDER

gormt \
  --host="127.0.0.1" \
  --user="root" \
  --password="123456" \
  --database="letter" \
  --table_names="user" \
  --outdir="./db/models/"

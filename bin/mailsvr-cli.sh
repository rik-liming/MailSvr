#!/bin/bash
if [[ ! $1 ]] || [[ ! $2 ]] || [[ ! $3 ]]; then
    echo "missing param, please check!"
    exit
fi

title=$1
content=$2
game=$3

curl http://mailsvr.xxx.com/mail/send -X POST -d "title=${title}&content=${content}&game=${game}"

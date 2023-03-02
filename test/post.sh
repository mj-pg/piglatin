#!/bin/bash

. $(dirname -- "$0")/config.sh


input="$@"
if [[ -z $input ]]; then
	echo "usage: $0 <text input>"
	exit 1
fi

template=$(dirname -- "$0")/req.json
body=`cat $template | sed "s/<INPUT>/$input/g"`

curl -X POST http://localhost:$APP_PORT/piglatins \
   -H 'Content-Type: application/json' \
   -d "$body"

echo

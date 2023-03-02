#!/bin/bash

. $(dirname -- "$0")/config.sh

curl http://localhost:$APP_PORT/piglatins

echo

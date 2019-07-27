#!/usr/bin/env bash

 set -eo pipefail

 #  A custom shell script to error whenever formatting is incorrect. We had to do a script becuase
#  `gofmt` does not return anything but `0` even with invalid formatting. This script iterates through each .go
#  file and runs gofmt, captures the output, and emits it if an error occurs, then after all files are
#  completed, it returns an error.

 declare -i HAD_ERROR
HAD_ERROR=0
for i in $(find . -path "./vendor" -prune -o -name '*.go' ); do
    \[ -f "$i" ] || continue

     # We want to ignore any auto generated matchers and auto generated mocks (they dont get committed)
    if [[ $i =~ .*\/matchers\/.* ]]; then
        echo "SKP - $i"
        continue
    elif [[ $i =~ mock_.*_test ]]; then
        echo "SKP - $i"
        continue
    fi

     OUTPUT=$(gofmt -l -e -d -s "$i")
    if [ -n "$OUTPUT" ]; then
        HAD_ERROR=1
        echo "===================================================================================="
        echo "ERR - $OUTPUT"
        echo "===================================================================================="
    else
        echo "OK  - $i"
    fi
done
exit $HAD_ERROR
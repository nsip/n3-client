#!/bin/bash

rm -rf ./build/

rm -f ./*/errlog.txt
rm -f ./*/temp.*
rm -f ./temp.*
rm -f ./gql/*/*.gql

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f

rm -f ./preprocess/util/jq*
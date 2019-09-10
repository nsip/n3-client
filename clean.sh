#!/bin/bash

rm -f ./build/debug*/*.json
rm -f ./build/*.json
rm -f ./build/n3-client
rm -f ./build/privacy
rm -f ./*/errlog.txt
rm -f ./*/temp.*
rm -f ./temp.*
rm -f ./gql/*/*.gql

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
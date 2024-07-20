#!/bin/bash

cat <<EOF > ./java/java.go
// TODO(adam): doc header
package java

import (
       "github.com/moov-io/ach"
)

// Types copied from moov-io/ach
EOF


# Copy every exported struct
names=($(grep -E "type [A-Z]{1,}[a-zA-Z0-9]* struct" *.go | grep -v Aux | cut -d' ' -f2 | sort -u))
for name in "${names[@]}"
do
    if [[ "$name" != "" ]];
    then
        echo "type $name ach.$name" >> ./java/java.go
    fi
done

cat <<EOF >> ./java/java.go

// Functions copied from moov-io/ach

EOF
# Copy every exported function
names=($(grep -E "func [A-Z]{1,}[a-zA-Z0-9]*\(" *.go | grep -v Benchmark | grep -v Test | cut -d':' -f2- | cut -d'(' -f1 | cut -d' ' -f2 | sort -u))
for name in "${names[@]}"
do
    if [[ "$name" != "" ]];
    then
        echo "var $name = ach.$name" >> ./java/java.go
    fi
done

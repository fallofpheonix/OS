#!/bin/bash
MODULES=$(grep -oE '\./[^ ]+' go.work)
for mod in $MODULES; do
    echo "=== $mod ==="
    (cd $mod && go test -cover ./... 2>/dev/null | grep -E "^ok|^FAIL|\?" | awk '{print $1 " " $2 " " $5}')
done

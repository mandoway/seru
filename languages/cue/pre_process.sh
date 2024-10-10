#!/bin/bash


target=$(realpath "$1")
cd "$target"

cue mod fix

set -e

matches=$(find . -name "*_tool.cue" -exec echo {} \;)
for file in $matches; do
    echo "Processing $file"
    new_name="${file%_tool.cue}.cue"
    mv "$file" "$new_name"
done

output=$(cue def --inline-imports -s)

echo "$output" > "$target"/in.cue

find . ! -name 'in.cue' ! -name 'test.sh' -exec rm -fr {} +
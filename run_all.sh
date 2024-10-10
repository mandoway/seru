#!/bin/bash

paths=(
  "./test/instances/error/issue_2209/v1"
  "./test/instances/error/issue_2209/final"
  "./test/instances/error/issue_2246/v1"
  "./test/instances/error/issue_2246/final"
  "./test/instances/error/issue_2473/v1"
  "./test/instances/error/issue_2473/final"
  "./test/instances/panic/issue_2490/v1"
  "./test/instances/panic/issue_2490/v2"
  "./test/instances/panic/issue_2490/final"
  "./test/instances/panic/issue_2584/v1"
  "./test/instances/panic/issue_2584/v2"
  "./test/instances/panic/issue_2584/final"
  "./test/instances/semantic/issue_2218/v1"
  "./test/instances/semantic/issue_2218/v2"
  "./test/instances/semantic/issue_2218/final"
)

go generate ./...
go build

for path in "${paths[@]}"; do
  if [ -e "$path" ]; then
    echo "Running 'seru' on '$path'..."
    ./seru -i "$path/in.cue" -t "$path/test.sh" -m
  else
    echo "Warning: '$path' does not exist."
  fi
done

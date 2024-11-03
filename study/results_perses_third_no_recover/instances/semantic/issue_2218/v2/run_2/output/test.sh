#!/bin/bash

readonly tmp_dir="${TMPDIR:-/tmp/}"
readonly cache_dir="${tmp_dir}cue-dd"

function cueVersion() {
  version=$1
  cached_bin="${cache_dir}/cue-${version}"

  if [[ ! -f $cached_bin ]]; then
    # until "go install -o" is supported; see https://github.com/golang/go/issues/44469
    if ! GOBIN=$cache_dir go install cuelang.org/go/cmd/cue@"${version}"; then
      echo "CUE ${version} download failed"
      exit 1
    fi
    mv "${cache_dir}/cue" "${cached_bin}"
  fi

  echo "$cached_bin"
}
		
set -e # forbid failing commands
output=$($(cueVersion v0.5.0-beta.2) eval in.cue || exit 1)
expected_output=': true$'
if grep -qE "$expected_output" <<< "$output"; then
	echo "Unexpected output found: ${output}"
	exit 1
fi
		
set -e # forbid failing commands
output=$($(cueVersion v0.4.3) eval in.cue || exit 1)
expected_output=': true$'
if ! grep -qE "$expected_output" <<< "$output"; then
	echo "Expected output missing, got: ${output}"
	exit 1
fi






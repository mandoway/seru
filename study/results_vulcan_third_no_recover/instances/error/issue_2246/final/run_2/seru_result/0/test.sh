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
		
set +e # allow failing commands
error_output="$("$(cueVersion v0.5.0-beta.5)" export in.cue 2>&1 > /dev/null)"
expected_output='field not allowed'
if ! grep -qE "$expected_output" <<< "$error_output"; then
	echo "Expected error missing, got: ${error_output}"
	exit 1
fi
		
set -e # forbid failing commands
output=$($(cueVersion v0.4.3) export in.cue || exit 1)






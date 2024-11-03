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
output=$(exec sh -c "CUE_EXPERIMENT= $(cueVersion v0.9.0-alpha.5) export -e output" || exit 1)
expected_output=DeploymentsWithRef
if ! grep -qE "$expected_output" <<< "$output"; then
	echo "Expected output missing, got: ${output}"
	exit 1
fi

set +e # allow failing commands
error_output="$(exec sh -c "CUE_EXPERIMENT=evalv3 '$(cueVersion v0.9.0-alpha.5)' export -e output" 2>&1 > /dev/null)"
expected_output='output._generated.0: adding field redis not allowed as field set was already referenced'
if ! grep -qE "$expected_output" <<< "$error_output"; then
	echo "Expected error missing, got: ${error_output}"
	exit 1
fi






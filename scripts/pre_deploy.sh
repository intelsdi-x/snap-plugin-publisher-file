#!/bin/bash

#http://www.apache.org/licenses/LICENSE-2.0.txt
#
#
#Copyright 2016 Intel Corporation
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

set -e
set -u
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__proj_dir="$(dirname "$__dir")"

# shellcheck source=scripts/common.sh
. "${__dir}/common.sh"

plugin_name=${__proj_dir##*/}
build_path="${SNAP_PATH:-"${__proj_dir}/build"}"
git_sha=$(git log --pretty=format:"%H" -1)
git_path="${build_path}/${TRAVIS_BRANCH}/${git_sha}"
latest_path="${build_path}/${TRAVIS_BRANCH}/latest"
release_path="${SNAP_PATH:-"${__proj_dir}/release"}"

mkdir -p "${git_path}"
mkdir -p "${latest_path}"
mkdir -p "${release_path}"

_info "build_path: ${build_path}"
find "${build_path}"

_info "copying plugin binaries to ${git_path}"
cp "${build_path}/${plugin_name}"* "${git_path}"
_info "copying plugin binaries to ${latest_path}"
cp "${build_path}/${plugin_name}"* "${latest_path}"

_info "moving plugin binaries to ${release_path}"
mv "${build_path}/${plugin_name}"* "${release_path}"

find "${build_path}"
find "${release_path}"

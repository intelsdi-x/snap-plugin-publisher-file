#!/bin/bash

set -e
set -u
set -o pipefail

# get the directory the script exists in
__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# source the common bash script 
. "${__dir}/../scripts/common.sh"

# ensure PLUGIN_PATH is set
TMPDIR=${TMPDIR:-"/tmp"}
PLUGIN_PATH=${PLUGIN_PATH:-"${TMPDIR}/snap/plugins"}
mkdir -p $PLUGIN_PATH

_info "Get latest plugins"
(cd $PLUGIN_PATH && curl -sSO http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/master/latest/linux/x86_64/snap-plugin-publisher-file && chmod 755 snap-plugin-publisher-file)
(cd $PLUGIN_PATH && curl -sSO http://snap.ci.snap-telemetry.io/snap/master/latest/snap-plugin-collector-mock1 && chmod 755 snap-plugin-collector-mock1)
(cd $PLUGIN_PATH && curl -sSO http://snap.ci.snap-telemetry.io/snap/master/latest/snap-plugin-processor-passthru && chmod 755 snap-plugin-processor-passthru)


_info "loading plugins"
snapctl plugin load "${PLUGIN_PATH}/snap-plugin-collector-mock1"
snapctl plugin load "${PLUGIN_PATH}/snap-plugin-publisher-file"
snapctl plugin load "${PLUGIN_PATH}/snap-plugin-processor-passthru"

_info "creating and starting a task"
snapctl task create -t "${__dir}/tasks/mock-passthru-file.json"
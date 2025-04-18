#!/usr/bin/env bash

# Modified from https://github.com/kubernetes-sigs/kueue/blob/065451d907fa27a0647436505b3cac38718ef136/hack/update-codegen.sh
# Apache-2.0, Copyright 2023 The Kubernetes Authors

set -o errexit
set -o nounset
set -o pipefail

PKG_ROOT=$(realpath "$(dirname ${BASH_SOURCE[0]})/..")

CODEGEN_PKG=/Users/vishal/worktest/code-generator

cd $PKG_ROOT

source "${CODEGEN_PKG}/kube_codegen.sh"

# TODO: remove the workaround when the issue is solved in code-generator
# (https://github.com/kubernetes/code-generator/issues/165).
# kube_codegen.sh expects this layout:
# .
# └── github.com
#     └── vishalanarase
#         └── memcached-operator
# We can use soft links in order to fake this layout, such that
# ./github.com/vishalanarase/memcached-operator resolves to ././../memcached-operator, or ./.

ln -s . github.com
ln -s .. vishalanarase
trap "rm github.com && rm vishalanarase" EXIT

kube::codegen::gen_helpers \
  --boilerplate /dev/null \
  github.com/vishalanarase/memcached-operator/pkg/apis

kube::codegen::gen_client \
  --output-pkg github.com/vishalanarase/memcached-operator/pkg/generated \
  --boilerplate /dev/null \
  --output-dir github.com/vishalanarase/memcached-operator/pkg/generated \
  --with-watch \
  --with-applyconfig \
  github.com/vishalanarase/memcached-operator/pkg/apis \

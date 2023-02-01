#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


# This is not supposed to be an error-prone script; just a convenience.

# Install CCM
pip install --user cql PyYAML six
git clone https://github.com/pcmanus/ccm.git
pushd ccm
./setup.py install --user
popd

if [ "$1" != "gocql/gocql" ]; then
    USER=$(echo $1 | cut -f1 -d'/')
    cd ../..
    mv ${USER} gocql
fi

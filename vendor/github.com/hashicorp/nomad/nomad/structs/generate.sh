#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

FILES="$(ls *[!_test].go | tr '\n' ' ')"
codecgen -d 100 -o structs.generated.go ${FILES}

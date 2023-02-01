#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


today=$(date +%Y%m%d)

sed -i -r -e 's/const Repo = "([0-9]{8})"/const Repo = "'$today'"/' $GOFILE


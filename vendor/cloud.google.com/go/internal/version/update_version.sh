#!/bin/bash
# Copyright IBM Corp. 2015, 2026


today=$(date +%Y%m%d)

sed -i -r -e 's/const Repo = "([0-9]{8})"/const Repo = "'$today'"/' $GOFILE


#!/bin/bash
# Copyright IBM Corp. 2015, 2026

set -e

FILES="$(ls *[!_test].go | tr '\n' ' ')"
codecgen -d 100 -o structs.generated.go ${FILES}

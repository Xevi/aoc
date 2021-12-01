#!/bin/sh
DAY=$1
CURRENT_YEAR=`date +"%Y"`
YEAR=${2:-$CURRENT_YEAR}

GOPATH=`pwd` go run src/$YEAR/$DAY/main.go
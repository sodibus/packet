#!/bin/bash

SRCROOT=resources/protobuf-definitions
TGTROOT=.

protoc -I=$SRCROOT --go_out=$TGTROOT $SRCROOT/*.proto

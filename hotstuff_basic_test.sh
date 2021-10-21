#!/bin/bash

clear
go test -v -count 1 github.com/devfans/zion/consensus/hotstuff/basic/core -run TestNewRound

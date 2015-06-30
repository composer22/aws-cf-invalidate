#!/usr/bin/bash

# Build all executables

# 386 = 32 bit
# amd = 64 bit
# arm = 32 bit
#
# $GOOS			$GOARCH
# ====================
# darwin		386
# darwin		amd64
# dragonfly		386
# dragonfly		amd64
# freebsd		386
# freebsd		amd64
# freebsd		arm
# linux			386
# linux			amd64
# linux			arm
# netbsd		386
# netbsd		amd64
# netbsd		arm
# openbsd		386
# openbsd		amd64
# plan9			386
# plan9			amd64
# solaris		amd64
# windows		386
# windows		amd64

type setopt >/dev/null 2>&1 && setopt shwordsplit
PLATFORMS="darwin/386 darwin/amd64 dragonfly/386 dragonfly/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm netbsd/386 netbsd/amd64 netbsd/arm openbsd/386 openbsd/amd64 plan9/386 plan9/amd64 solaris/amd64 windows/386 windows/amd64"

function go-compile {
	local GOOS=${1%/*}
	local GOARCH=${1#*/}
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o aws-cf-invalidate-${GOOS}-${GOARCH} -i
}

function run {
	for PLATFORM in $PLATFORMS; do
			local CMD="go-compile ${PLATFORM}"
			echo "$CMD"
			$CMD
	done
}

run

mv aws-cf-invalidate-windows-386 aws-cf-invalidate-windows-386.exe
mv aws-cf-invalidate-windows-amd64 aws-cf-invalidate-windows-amd64.exe

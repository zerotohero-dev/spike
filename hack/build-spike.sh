#!/usr/bin/env bash

#    \\ SPIKE: Keep your secrets secret with SPIFFE.
#  \\\\\ Copyright 2024-present SPIKE contributors.
# \\\\\\\ SPDX-License-Identifier: Apache-2.0

rm keeper
rm nexus
rm pilot

# `boringcrypto` is required for FIPS compliance.

CGO_ENABLED=0 GOEXPERIMENT=boringcrypto go build -o keeper ./app/keeper/cmd/main.go
CGO_ENABLED=0 GOEXPERIMENT=boringcrypto go build -o nexus ./app/nexus/cmd/main.go
CGO_ENABLED=0 GOEXPERIMENT=boringcrypto go build -o pilot ./app/pilot/cmd/main.go
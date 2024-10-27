#!/usr/bin/env bash

#    \\ SPIKE: Keep your secrets secret with SPIFFE.
#  \\\\\ Copyright 2024-present SPIKE contributors.
# \\\\\\\ SPDX-License-Identifier: Apache-2.0

# Verify file was created and is not empty
if [ ! -s .token ]; then
    echo "Error: Token file is empty or was not created" >&2
    exit 1
fi

if [ ! -f .token ]; then
    echo "Error: token does not exist"
    exit 1
fi

JOIN_TOKEN=$(cat .token)
if [ -z "$JOIN_TOKEN" ]; then
    echo "Error: join token is empty"
    exit 1
fi

spire-agent run -config ./config/spire/agent/agent.conf -joinToken "$JOIN_TOKEN"

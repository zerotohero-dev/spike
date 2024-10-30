#!/usr/bin/env bash

#    \\ SPIKE: Keep your secrets secret with SPIFFE.
#  \\\\\ Copyright 2024-present SPIKE contributors.
# \\\\\\\ SPDX-License-Identifier: Apache-2.0

if [ ! -f .token ]; then
    echo "Error: token does not exist"
    exit 1
fi

TOKEN=$(cat .token)

if [ -z "$TOKEN" ]; then
    echo "Error: token is empty"
    exit 1
fi

echo "Token loaded successfully."

# We probably don't need this.
# AGENT_ID="spiffe://spike.ist/spire-agent/join_token/$TOKEN"

KEEPER_PATH="$(pwd)/keeper"
KEEPER_SHA=$(sha256sum "$KEEPER_PATH" | cut -d' ' -f1)

NEXUS_PATH="$(pwd)/nexus"
NEXUS_SHA=$(sha256sum "$NEXUS_PATH" | cut -d' ' -f1)

PILOT_PATH="$(pwd)/spike"
PILOT_SHA=$(sha256sum "$PILOT_PATH" | cut -d' ' -f1)

# Register SPIKE Keeper
spire-server entry create \
    -spiffeID spiffe://spike.ist/spike/keeper \
    -parentID "spiffe://spike.ist/spire-agent" \
    -selector unix:uid:"$(id -u)" \
    -selector unix:path:"$KEEPER_PATH" \
    -selector unix:sha256:"$KEEPER_SHA"

# Register SPIKE Nexus
spire-server entry create \
    -spiffeID spiffe://spike.ist/spike/nexus \
    -parentID "spiffe://spike.ist/spire-agent" \
    -selector unix:uid:"$(id -u)" \
    -selector unix:path:"$NEXUS_PATH" \
    -selector unix:sha256:"$NEXUS_SHA"

# Register SPIKE Pilot
spire-server entry create \
    -spiffeID spiffe://spike.ist/spike/pilot \
    -parentID "spiffe://spike.ist/spire-agent" \
    -selector unix:uid:"$(id -u)" \
    -selector unix:path:"$PILOT_PATH" \
    -selector unix:sha256:"$PILOT_SHA"

#!/usr/bin/env bash

#    \\ SPIKE: Keep your secrets secret with SPIFFE.
#  \\\\\ Copyright 2024-present SPIKE contributors.
# \\\\\\\ SPDX-License-Identifier: Apache-2.0

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

// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

// ChainAliasPrefix is the path segment a chain's routes hang from:
// /v1/chain/<id>/rpc.
//
// It was "bc", short for blockchain, and the abbreviation asserted something
// that is not true of every chain here — a Lux chain may be a DAG, and naming
// the shape in the route means the route is wrong wherever the shape differs.
// "chain" is what these all are.
const ChainAliasPrefix string = "chain"

// VMAliasPrefix denotes a prefix for an alias that belongs to a VM ID.
const VMAliasPrefix string = "vm"

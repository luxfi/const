// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package constants provides Lux network constants
package constants

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/luxfi/ids"
	"github.com/luxfi/math/set"
)

// Network IDs
const (
	LocalID    uint32 = 1337
	MainnetID  uint32 = 1 // EVM-compatible
	TestnetID  uint32 = 2 // EVM-compatible
	UnitTestID uint32 = 369

	// Lux-specific network IDs
	LuxMainnetID uint32 = 96369 // Lux mainnet
	LuxTestnetID uint32 = 96368 // Lux testnet
)

// Network names
const (
	LocalName    = "local"
	MainnetName  = "mainnet"
	TestnetName  = "testnet"
	UnitTestName = "testing"
)

// Human-readable parts (HRP) for addresses
const (
	FallbackHRP = "custom"
	LocalHRP    = "local"
	MainnetHRP  = "lux"
	TestnetHRP  = "test"
	UnitTestHRP = "testing"
)

// Well-known IDs
var (
	PrimaryNetworkID = ids.Empty
	PlatformChainID  = ids.Empty
	XChainID         = ids.ID{'x', 'c', 'h', 'a', 'i', 'n'}
	CChainID         = ids.ID{'c', 'c', 'h', 'a', 'i', 'n'}
)

// Memory constants
const (
	PointerOverhead = 8 // bytes for a pointer on 64-bit systems
)

// Mapping tables
var (
	NetworkIDToNetworkName = map[uint32]string{
		LocalID:         LocalName,
		MainnetID:       MainnetName,
		TestnetID:       TestnetName,
		UnitTestID:      UnitTestName,
		LuxMainnetID:    MainnetName,
		LuxTestnetID:    TestnetName,
		QChainMainnetID: "qchain-mainnet",
		QChainTestnetID: "qchain-testnet",
	}

	NetworkNameToNetworkID = map[string]uint32{
		LocalName:    LocalID,
		MainnetName:  MainnetID,
		TestnetName:  TestnetID,
		UnitTestName: UnitTestID,
	}

	NetworkIDToHRP = map[uint32]string{
		LocalID:         LocalHRP,
		MainnetID:       MainnetHRP,
		TestnetID:       TestnetHRP,
		UnitTestID:      UnitTestHRP,
		LuxMainnetID:    MainnetHRP,
		LuxTestnetID:    TestnetHRP,
		QChainMainnetID: "qchain",
		QChainTestnetID: "qtest",
	}

	NetworkHRPToNetworkID = map[string]uint32{
		LocalHRP:    LocalID,
		MainnetHRP:  MainnetID,
		TestnetHRP:  TestnetID,
		UnitTestHRP: UnitTestID,
	}

	// ProductionNetworkIDs are networks that should use production-grade settings
	ProductionNetworkIDs = set.Of(MainnetID, TestnetID, LuxMainnetID, LuxTestnetID)

	ValidNetworkPrefix = "network-"
)

// Errors
var (
	ErrUnknownNetworkID   = errors.New("unknown network ID")
	ErrUnknownNetworkName = errors.New("unknown network name")
)

// NetworkName returns the name for the given network ID
func NetworkName(networkID uint32) string {
	if name, ok := NetworkIDToNetworkName[networkID]; ok {
		return name
	}
	return fmt.Sprintf("%s%d", ValidNetworkPrefix, networkID)
}

// NetworkID returns the network ID for the given network name
func NetworkID(name string) (uint32, error) {
	name = strings.ToLower(name)
	if id, ok := NetworkNameToNetworkID[name]; ok {
		return id, nil
	}

	if strings.HasPrefix(name, ValidNetworkPrefix) {
		idStr := strings.TrimPrefix(name, ValidNetworkPrefix)
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err == nil {
			return uint32(id), nil
		}
	}

	return 0, ErrUnknownNetworkName
}

// GetHRP returns the human-readable part for the given network ID
func GetHRP(networkID uint32) string {
	if hrp, ok := NetworkIDToHRP[networkID]; ok {
		return hrp
	}
	return FallbackHRP
}

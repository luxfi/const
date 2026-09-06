// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

// BlackholeAddr is the address where assets are burned: 20 bytes whose only
// nonzero byte is the first.
//
// It is a plain array rather than a geth common.Address so that constants — a
// leaf every module can import — does not pull in geth, and through it crypto.
// A caller that needs a common.Address writes common.Address(BlackholeAddr);
// the conversion is free and the dependency is theirs, not this package's.
var BlackholeAddr = [20]byte{
	1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

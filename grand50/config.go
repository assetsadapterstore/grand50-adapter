/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package grand50

import (
	"github.com/assetsadapterstore/grand50-adapter/g50Transaction"
	"github.com/blocktree/go-owcrypt"
)

const (
	//币种
	Symbol    = "G50"
	CurveType = owcrypt.ECC_CURVE_SECP256K1
	Decimals  = int32(8)
)

var (
	MainNetAddressPrefix = g50Transaction.AddressPrefix{P2PKHPrefix: []byte{0x26}, P2WPKHPrefix: []byte{0x08}, P2SHPrefix: nil, Bech32Prefix: "g50"}
	TestNetAddressPrefix = g50Transaction.AddressPrefix{P2PKHPrefix: []byte{0x61}, P2WPKHPrefix: []byte{0x09}, P2SHPrefix: nil, Bech32Prefix: "g50"}
)

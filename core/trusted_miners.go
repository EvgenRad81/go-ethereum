// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
)

func ReadTrustedMiners(reader io.Reader) ([]common.Address, error) {
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	address_list := []string{}
	if err := json.Unmarshal(contents, &address_list); err != nil {
		return nil, err
	}

	trusted_miners := []common.Address{}
	for _, addr := range address_list {
		trusted_miners = append(trusted_miners, common.HexToAddress(addr))
	}

	return trusted_miners, nil
}

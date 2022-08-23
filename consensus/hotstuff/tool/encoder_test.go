/*
 * Copyright (C) 2021 The Zion Authors
 * This file is part of The Zion library.
 *
 * The Zion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Zion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The Zion.  If not, see <http://www.gnu.org/licenses/>.
 */

package tool

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGenesisExtra(t *testing.T) {
	validators := []common.Address{
		common.HexToAddress("0x40FBBE484b8Ee6139Af08446950B088e10b2306A"),
		common.HexToAddress("0x8C161d85fDC086AC6726bCEDe39f2CCB1Afa3bc8"),
		common.HexToAddress("0x22C21D4F64aabA7ec837b6B93639dB8cF514dAD5"),
		common.HexToAddress("0xBE6805C4c904B9cc39065ADC4cCcF4FCaB167AE6"),
	}
	enc, err := EncodeGenesisExtra(validators)
	assert.NoError(t, err)
	t.Log(enc)
}

func TestEncode(t *testing.T) {
	var testNodeKeys = []*Node{
		{
			NodeKey: "562aa98da69477996bd82422b97698541f25e71ba2f803970947b3ad8bdb7afa",
		},
		{
			NodeKey: "8bea3ce27136df435ada62a40a4226404879b3c42e2e86ba9a236b4a61c99c26",
		},
		{
			NodeKey: "53f7d9ec7657cdd3a3eaa8ddd126d36fbc60203448fca1bbfccec0d59d173da6",
		},
		{
			NodeKey: "305baf1e19a2da40b413dfb62b206b0ac74cb3d7e975cb70fe8391cbbe174f2a",
		},
	}

	dumpNodes(t, testNodeKeys)
}

func TestGenerateAndEncode(t *testing.T) {
	nodes := generateNodes(4)
	dumpNodes(t, nodes)
}

func dumpNodes(t *testing.T, nodes []*Node) {
	sortedNodes := SortNodes(nodes)
	staticNodes := make([]string, 0)
	for _, v := range sortedNodes {
		t.Logf("addr: %s, pubKey: %s, nodeKey: %s", v.Address(), v.Pubkey(), v.NodeKey)
		staticNodes = append(staticNodes, v.Static)
	}
	t.Log("==================================================================")

	genesis, err := EncodeGenesisExtra(NodesAddress(sortedNodes))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("genesis extra %s", genesis)

	t.Log("==================================================================")

	staticNodesEnc, err := json.MarshalIndent(staticNodes, "", "\t")
	t.Log(string(staticNodesEnc))
}

func generateNodes(n int) []*Node {
	nodes := make([]*Node, 0)
	for i := 0; i < n; i++ {
		key, _ := crypto.GenerateKey()
		nodekey := hexutil.Encode(crypto.FromECDSA(key))
		node := &Node{
			NodeKey: nodekey,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

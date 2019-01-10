// Copyright 2016 The go-ethereum Authors
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

package network

import (
	"net"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/discover"
	p2ptest "github.com/ethereum/go-ethereum/p2p/testing"
)

/***
 *
 * - after connect, that outgoing subpeersmsg is sent
 *
 */
func TestDiscovery(t *testing.T) {
	params := NewHiveParams()
	s, pp := newHiveTester(t, params, 1, nil)

	nid := s.IDs[0]
	node := discover.NewNode(nid, net.IP{127, 0, 0, 1}, 30303, 30303)

	raddr := NewAddr(node)
	pp.Register(raddr)

	// start the hive and wait for the connection
	pp.Start(s.Server)
	defer pp.Stop()

	// send subPeersMsg to the peer
	err := s.TestExchanges(p2ptest.Exchange{
		Label: "outgoing subPeersMsg",
		Expects: []p2ptest.Expect{
			{
				Code: 1,
				Msg:  &subPeersMsg{Depth: 0},
				Peer: node.ID,
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}

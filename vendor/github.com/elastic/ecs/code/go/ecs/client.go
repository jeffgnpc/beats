// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// A client is defined as the initiator of a network connection for events
// regarding sessions, connections, or bidirectional flow records. For TCP
// events, the client is the initiator of the TCP connection that sends the SYN
// packet(s). For other protocols, the client is generally the initiator or
// requestor in the network transaction. Some systems use the term "originator"
// to refer the client in TCP connections. The client fields describe details
// about the system acting as the client in the network event. Client fields
// are usually populated in conjunction with server fields.  Client fields are
// generally not populated for packet-level events.
// Client / server representations can add semantic context to an exchange,
// which is helpful to visualize the data in certain situations. If your
// context falls in that category, you should still ensure that source and
// destination are filled appropriately.
type Client struct {
	// Some event client addresses are defined ambiguously. The event will
	// sometimes list an IP, a domain or a unix socket.  You should always
	// store the raw address in the `.address` field.
	// Then it should be duplicated to `.ip` or `.domain`, depending on which
	// one it is.
	Address string `ecs:"address"`

	// IP address of the client.
	// Can be one or multiple IPv4 or IPv6 addresses.
	IP string `ecs:"ip"`

	// Port of the client.
	Port int64 `ecs:"port"`

	// MAC address of the client.
	MAC string `ecs:"mac"`

	// Client domain.
	Domain string `ecs:"domain"`

	// Bytes sent from the client to the server.
	Bytes int64 `ecs:"bytes"`

	// Packets sent from the client to the server.
	Packets int64 `ecs:"packets"`
}
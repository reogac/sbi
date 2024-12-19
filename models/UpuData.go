/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpuData struct {
	Drei             *bool    `json:"drei,omitempty"`
	Aol              *bool    `json:"aol,omitempty"`
	SecPacket        string   `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
	RoutingId        string   `json:"routingId,omitempty"`
}

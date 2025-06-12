/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpuData struct {
	RoutingId        string   `json:"routingId,omitempty"`
	Drei             *bool    `json:"drei,omitempty"`
	Aol              *bool    `json:"aol,omitempty"`
	SecPacket        string   `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpuData struct {
	Aol              *bool    `json:"aol,omitempty"`
	SecPacket        string   `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
	RoutingId        string   `json:"routingId,omitempty"`
	Drei             *bool    `json:"drei,omitempty"`
}

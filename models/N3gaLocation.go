/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N3gaLocation struct {
	PortNumber     *int              `json:"portNumber,omitempty"`
	TnapId         *TnapId           `json:"tnapId,omitempty"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty"`
	Gci            string            `json:"gci,omitempty"`
	N3IwfId        string            `json:"n3IwfId,omitempty"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty"`
	Gli            string            `json:"gli,omitempty"`
	N3gppTai       *Tai              `json:"n3gppTai,omitempty"`
	Protocol       TransportProtocol `json:"protocol,omitempty"`
	TwapId         *TwapId           `json:"twapId,omitempty"`
}

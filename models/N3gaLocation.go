/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N3gaLocation struct {
	Gli            string            `json:"gli,omitempty"`
	N3IwfId        string            `json:"n3IwfId,omitempty"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty"`
	PortNumber     *int              `json:"portNumber,omitempty"`
	TwapId         *TwapId           `json:"twapId,omitempty"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty"`
	Gci            string            `json:"gci,omitempty"`
	N3gppTai       *Tai              `json:"n3gppTai,omitempty"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty"`
	Protocol       TransportProtocol `json:"protocol,omitempty"`
	TnapId         *TnapId           `json:"tnapId,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:01:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N3gaLocation struct {
	Protocol       TransportProtocol `json:"protocol,omitempty"`
	TnapId         *TnapId           `json:"tnapId,omitempty"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty"`
	N3gppTai       *Tai              `json:"n3gppTai,omitempty"`
	N3IwfId        string            `json:"n3IwfId,omitempty"`
	PortNumber     *int              `json:"portNumber,omitempty"`
	TwapId         *TwapId           `json:"twapId,omitempty"`
	Gli            string            `json:"gli,omitempty"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty"`
	Gci            string            `json:"gci,omitempty"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty"`
}

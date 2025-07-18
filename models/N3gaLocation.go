/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N3gaLocation struct {
	TnapId         *TnapId           `json:"tnapId,omitempty"`
	TwapId         *TwapId           `json:"twapId,omitempty"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty"`
	Gli            string            `json:"gli,omitempty"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty"`
	PortNumber     *int              `json:"portNumber,omitempty"`
	Protocol       TransportProtocol `json:"protocol,omitempty"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty"`
	Gci            string            `json:"gci,omitempty"`
	N3gppTai       *Tai              `json:"n3gppTai,omitempty"`
	N3IwfId        string            `json:"n3IwfId,omitempty"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty"`
}

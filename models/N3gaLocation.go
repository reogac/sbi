/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N3gaLocation struct {
	TnapId         *TnapId           `json:"tnapId,omitempty"`
	TwapId         *TwapId           `json:"twapId,omitempty"`
	Gli            string            `json:"gli,omitempty"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty"`
	N3IwfId        string            `json:"n3IwfId,omitempty"`
	Protocol       TransportProtocol `json:"protocol,omitempty"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty"`
	PortNumber     *int              `json:"portNumber,omitempty"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty"`
	Gci            string            `json:"gci,omitempty"`
	N3gppTai       *Tai              `json:"n3gppTai,omitempty"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty"`
}

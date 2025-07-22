/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EthFlowDescription struct {
	DestMacAddr    string        `json:"destMacAddr,omitempty"`
	EthType        string        `json:"ethType"`
	FDesc          string        `json:"fDesc,omitempty"`
	FDir           FlowDirection `json:"fDir,omitempty"`
	SourceMacAddr  string        `json:"sourceMacAddr,omitempty"`
	VlanTags       []string      `json:"vlanTags,omitempty"`
	SrcMacAddrEnd  string        `json:"srcMacAddrEnd,omitempty"`
	DestMacAddrEnd string        `json:"destMacAddrEnd,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EthFlowDescription struct {
	FDesc          string        `json:"fDesc,omitempty"`
	FDir           FlowDirection `json:"fDir,omitempty"`
	SourceMacAddr  string        `json:"sourceMacAddr,omitempty"`
	VlanTags       []string      `json:"vlanTags,omitempty"`
	SrcMacAddrEnd  string        `json:"srcMacAddrEnd,omitempty"`
	DestMacAddrEnd string        `json:"destMacAddrEnd,omitempty"`
	DestMacAddr    string        `json:"destMacAddr,omitempty"`
	EthType        string        `json:"ethType"`
}

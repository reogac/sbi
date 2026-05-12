/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EthFlowDescription struct {
	EthType        string        `json:"ethType"`
	FDesc          string        `json:"fDesc,omitempty"`
	FDir           FlowDirection `json:"fDir,omitempty"`
	SourceMacAddr  string        `json:"sourceMacAddr,omitempty"`
	VlanTags       []string      `json:"vlanTags,omitempty"`
	SrcMacAddrEnd  string        `json:"srcMacAddrEnd,omitempty"`
	DestMacAddrEnd string        `json:"destMacAddrEnd,omitempty"`
	DestMacAddr    string        `json:"destMacAddr,omitempty"`
}

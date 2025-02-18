/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:05:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models
type HandoverRequestAcknowledge struct {
	 RanUeId	RanUeId	`json:"ranUeId"`
	 RanNets	[]string	`json:"ranNets,omitempty"`
	 NfSelection	map[string]string	`json:"nfSelection,omitempty"`
	 Sessions	[]N2SmInfoUplinkContent	`json:"sessions,omitempty"`
	 TargetToSourceContent	[]byte	`json:"targetToSourceContent,omitempty"`
	 NasSplit	bool	`json:"nasSplit"`
}

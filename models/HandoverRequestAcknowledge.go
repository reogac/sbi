/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:06 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequestAcknowledge struct {
	TargetToSourceContent []byte                  `json:"targetToSourceContent,omitempty"`
	NasSplit              bool                    `json:"nasSplit"`
	RanUeId               RanUeId                 `json:"ranUeId"`
	RanNets               []string                `json:"ranNets,omitempty"`
	NfSelection           map[string]string       `json:"nfSelection,omitempty"`
	Sessions              []N2SmInfoUplinkContent `json:"sessions,omitempty"`
}

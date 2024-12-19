/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:26:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequestAcknowledge struct {
	NasSplit              bool                    `json:"nasSplit"`
	RanUeId               RanUeId                 `json:"ranUeId"`
	RanNets               []string                `json:"ranNets,omitempty"`
	NfSelection           map[string]string       `json:"nfSelection,omitempty"`
	Sessions              []N2SmInfoUplinkContent `json:"sessions,omitempty"`
	TargetToSourceContent []byte                  `json:"targetToSourceContent,omitempty"`
}

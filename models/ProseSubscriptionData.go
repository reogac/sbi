/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProseSubscriptionData struct {
	ProseServiceAuth *ProseServiceAuth  `json:"proseServiceAuth,omitempty"`
	NrUePc5Ambr      string             `json:"nrUePc5Ambr,omitempty"`
	ProseAllowedPlmn []ProSeAllowedPlmn `json:"proseAllowedPlmn,omitempty"`
}

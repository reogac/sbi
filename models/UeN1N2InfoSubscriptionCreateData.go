/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeN1N2InfoSubscriptionCreateData struct {
	N2InformationClass  N2InformationClass `json:"n2InformationClass,omitempty"`
	N2NotifyCallbackUri string             `json:"n2NotifyCallbackUri,omitempty"`
	N1MessageClass      N1MessageClass     `json:"n1MessageClass,omitempty"`
	N1NotifyCallbackUri string             `json:"n1NotifyCallbackUri,omitempty"`
	NfId                string             `json:"nfId,omitempty"`
	SupportedFeatures   string             `json:"supportedFeatures,omitempty"`
	OldGuami            *Guami             `json:"oldGuami,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonUeN2InfoSubscriptionCreatedData struct {
	SupportedFeatures      string             `json:"supportedFeatures,omitempty"`
	N2InformationClass     N2InformationClass `json:"n2InformationClass,omitempty"`
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
}

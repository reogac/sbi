/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonUeN2InfoSubscriptionCreatedData struct {
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	SupportedFeatures      string             `json:"supportedFeatures,omitempty"`
	N2InformationClass     N2InformationClass `json:"n2InformationClass,omitempty"`
}

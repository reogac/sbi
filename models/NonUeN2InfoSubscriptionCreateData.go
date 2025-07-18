/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonUeN2InfoSubscriptionCreateData struct {
	GlobalRanNodeList   []GlobalRanNodeId  `json:"globalRanNodeList,omitempty"`
	AnTypeList          []string           `json:"anTypeList,omitempty"`
	N2InformationClass  N2InformationClass `json:"n2InformationClass"`
	N2NotifyCallbackUri string             `json:"n2NotifyCallbackUri"`
	NfId                string             `json:"nfId,omitempty"`
	SupportedFeatures   string             `json:"supportedFeatures,omitempty"`
}

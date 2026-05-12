/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonUeN2InfoSubscriptionCreateData struct {
	AnTypeList          []string           `json:"anTypeList,omitempty"`
	N2InformationClass  N2InformationClass `json:"n2InformationClass"`
	N2NotifyCallbackUri string             `json:"n2NotifyCallbackUri"`
	NfId                string             `json:"nfId,omitempty"`
	SupportedFeatures   string             `json:"supportedFeatures,omitempty"`
	GlobalRanNodeList   []GlobalRanNodeId  `json:"globalRanNodeList,omitempty"`
}

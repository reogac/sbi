/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfSelectionSubscriptionData struct {
	HssGroupId            string                `json:"hssGroupId,omitempty"`
	SupportedFeatures     string                `json:"supportedFeatures,omitempty"`
	SubscribedSnssaiInfos map[string]SnssaiInfo `json:"subscribedSnssaiInfos,omitempty"`
	SharedSnssaiInfosId   string                `json:"sharedSnssaiInfosId,omitempty"`
}

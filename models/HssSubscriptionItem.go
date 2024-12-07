/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssSubscriptionItem struct {
	SubscriptionId string       `json:"subscriptionId"`
	ContextInfo    *ContextInfo `json:"contextInfo,omitempty"`
	HssInstanceId  string       `json:"hssInstanceId"`
}

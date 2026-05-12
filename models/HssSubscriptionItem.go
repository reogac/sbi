/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssSubscriptionItem struct {
	SubscriptionId string       `json:"subscriptionId"`
	ContextInfo    *ContextInfo `json:"contextInfo,omitempty"`
	HssInstanceId  string       `json:"hssInstanceId"`
}

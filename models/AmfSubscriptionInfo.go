/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSubscriptionInfo struct {
	ContextInfo                   *ContextInfo `json:"contextInfo,omitempty"`
	AmfInstanceId                 string       `json:"amfInstanceId"`
	SubscriptionId                string       `json:"subscriptionId"`
	SubsChangeNotifyCorrelationId string       `json:"subsChangeNotifyCorrelationId,omitempty"`
}

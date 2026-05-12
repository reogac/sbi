/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSubscriptionInfo struct {
	AmfInstanceId                 string       `json:"amfInstanceId"`
	SubscriptionId                string       `json:"subscriptionId"`
	SubsChangeNotifyCorrelationId string       `json:"subsChangeNotifyCorrelationId,omitempty"`
	ContextInfo                   *ContextInfo `json:"contextInfo,omitempty"`
}

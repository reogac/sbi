/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdpSubscriptionData struct {
	SupportedFeatures        string `json:"supportedFeatures,omitempty"`
	UpdpCallbackBinding      string `json:"updpCallbackBinding,omitempty"`
	UpdpNotifySubscriptionId string `json:"updpNotifySubscriptionId"`
	UpdpNotifyCallbackUri    string `json:"updpNotifyCallbackUri"`
}

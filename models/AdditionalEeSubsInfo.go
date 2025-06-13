/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalEeSubsInfo struct {
	AmfSubscriptionInfoList []AmfSubscriptionInfo `json:"amfSubscriptionInfoList,omitempty"`
	SmfSubscriptionInfo     *SmfSubscriptionInfo  `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo     *HssSubscriptionInfo  `json:"hssSubscriptionInfo,omitempty"`
}

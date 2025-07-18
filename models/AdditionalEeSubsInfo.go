/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalEeSubsInfo struct {
	AmfSubscriptionInfoList []AmfSubscriptionInfo `json:"amfSubscriptionInfoList,omitempty"`
	SmfSubscriptionInfo     *SmfSubscriptionInfo  `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo     *HssSubscriptionInfo  `json:"hssSubscriptionInfo,omitempty"`
}

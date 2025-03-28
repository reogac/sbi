/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GbrQosFlowInformation struct {
	AlternativeQosProfileList []AlternativeQosProfile `json:"alternativeQosProfileList,omitempty"`
	MaxFbrDl                  string                  `json:"maxFbrDl"`
	MaxFbrUl                  string                  `json:"maxFbrUl"`
	GuaFbrDl                  string                  `json:"guaFbrDl"`
	GuaFbrUl                  string                  `json:"guaFbrUl"`
	NotifControl              NotificationControl     `json:"notifControl,omitempty"`
	MaxPacketLossRateDl       *int                    `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl       *int                    `json:"maxPacketLossRateUl,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GbrQosFlowInformation struct {
	MaxFbrDl                  string                  `json:"maxFbrDl"`
	MaxFbrUl                  string                  `json:"maxFbrUl"`
	GuaFbrDl                  string                  `json:"guaFbrDl"`
	GuaFbrUl                  string                  `json:"guaFbrUl"`
	NotifControl              NotificationControl     `json:"notifControl,omitempty"`
	MaxPacketLossRateDl       *int                    `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl       *int                    `json:"maxPacketLossRateUl,omitempty"`
	AlternativeQosProfileList []AlternativeQosProfile `json:"alternativeQosProfileList,omitempty"`
}

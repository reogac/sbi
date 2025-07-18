/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextRetrievedData struct {
	DlDataWaitingInd    *bool                `json:"dlDataWaitingInd,omitempty"`
	AfCoordinationInfo  *AfCoordinationInfo  `json:"afCoordinationInfo,omitempty"`
	UeEpsPdnConnection  string               `json:"ueEpsPdnConnection"`
	SmContext           *SmContext           `json:"smContext,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
}

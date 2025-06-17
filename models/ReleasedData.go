/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleasedData struct {
	N4InfoExt1          *N4Information       `json:"n4InfoExt1,omitempty"`
	N4InfoExt2          *N4Information       `json:"n4InfoExt2,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	N4Info              *N4Information       `json:"n4Info,omitempty"`
}

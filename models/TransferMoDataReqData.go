/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMoDataReqData struct {
	MoExpDataCounter *MoExpDataCounter `json:"moExpDataCounter,omitempty"`
	UeLocation       *UserLocation     `json:"ueLocation,omitempty"`
	MoData           RefToBinaryData   `json:"moData"`
}

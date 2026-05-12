/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMoDataRequest struct {
	JsonData     *TransferMoDataReqData `json:"jsonData,omitempty"`
	BinaryMoData []byte                 `json:"binaryMoData,omitempty"`
}

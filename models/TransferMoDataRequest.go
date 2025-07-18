/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMoDataRequest struct {
	JsonData     *TransferMoDataReqData `json:"jsonData,omitempty"`
	BinaryMoData []byte                 `json:"binaryMoData,omitempty"`
}

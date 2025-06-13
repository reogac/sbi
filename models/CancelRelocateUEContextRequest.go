/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CancelRelocateUEContextRequest struct {
	BinaryDataGtpcMessage []byte                       `json:"binaryDataGtpcMessage,omitempty"`
	JsonData              *UeContextCancelRelocateData `json:"jsonData,omitempty"`
}

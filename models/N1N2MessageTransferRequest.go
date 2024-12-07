/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferRequest struct {
	BinaryMtData            []byte                      `json:"binaryMtData,omitempty"`
	JsonData                *N1N2MessageTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN1Message     []byte                      `json:"binaryDataN1Message,omitempty"`
	BinaryDataN2Information []byte                      `json:"binaryDataN2Information,omitempty"`
}

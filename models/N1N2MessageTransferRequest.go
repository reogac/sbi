/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferRequest struct {
	JsonData                *N1N2MessageTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN1Message     []byte                      `json:"binaryDataN1Message,omitempty"`
	BinaryDataN2Information []byte                      `json:"binaryDataN2Information,omitempty"`
	BinaryMtData            []byte                      `json:"binaryMtData,omitempty"`
}

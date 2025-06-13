/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonUeN2MessageTransferRequest struct {
	JsonData                *N2InformationTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN2Information []byte                        `json:"binaryDataN2Information,omitempty"`
}

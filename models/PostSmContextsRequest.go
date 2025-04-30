/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PostSmContextsRequest struct {
	JsonData                      *SmContextCreateData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage         []byte               `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation     []byte               `json:"binaryDataN2SmInformation,omitempty"`
	BinaryDataN2SmInformationExt1 []byte               `json:"binaryDataN2SmInformationExt1,omitempty"`
}

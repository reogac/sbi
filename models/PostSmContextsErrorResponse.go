/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PostSmContextsErrorResponse struct {
	BinaryDataN1SmMessage []byte                `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmMessage []byte                `json:"binaryDataN2SmMessage,omitempty"`
	JsonData              *SmContextCreateError `json:"jsonData,omitempty"`
}

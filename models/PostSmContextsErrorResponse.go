/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PostSmContextsErrorResponse struct {
	JsonData              *SmContextCreateError `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage []byte                `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmMessage []byte                `json:"binaryDataN2SmMessage,omitempty"`
}

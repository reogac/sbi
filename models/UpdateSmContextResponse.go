/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdateSmContextResponse struct {
	JsonData                  *SmContextUpdatedData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage     []byte                `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty"`
}

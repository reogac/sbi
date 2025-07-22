/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdateSmContextResponse struct {
	BinaryDataN1SmMessage     []byte                `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty"`
	JsonData                  *SmContextUpdatedData `json:"jsonData,omitempty"`
}

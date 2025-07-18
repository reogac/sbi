/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdateSmContextErrorResponse struct {
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty"`
	JsonData                  *SmContextUpdateError `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage     []byte                `json:"binaryDataN1SmMessage,omitempty"`
}

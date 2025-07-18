/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PostPduSessionsRequest struct {
	BinaryDataUnknownN1SmInfo []byte                `json:"binaryDataUnknownN1SmInfo,omitempty"`
	JsonData                  *PduSessionCreateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe  []byte                `json:"binaryDataN1SmInfoFromUe,omitempty"`
}

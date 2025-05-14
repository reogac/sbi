/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PostPduSessionsRequest struct {
	JsonData                  *PduSessionCreateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe  []byte                `json:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo []byte                `json:"binaryDataUnknownN1SmInfo,omitempty"`
}

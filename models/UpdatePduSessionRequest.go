/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdatePduSessionRequest struct {
	BinaryDataN4InformationExt1 []byte          `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 []byte          `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *HsmfUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe    []byte          `json:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo   []byte          `json:"binaryDataUnknownN1SmInfo,omitempty"`
	BinaryDataN4Information     []byte          `json:"binaryDataN4Information,omitempty"`
}

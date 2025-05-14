/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdatePduSessionResponse struct {
	BinaryDataN4InformationExt2 []byte           `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *HsmfUpdatedData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoToUe      []byte           `json:"binaryDataN1SmInfoToUe,omitempty"`
	BinaryDataN4Information     []byte           `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 []byte           `json:"binaryDataN4InformationExt1,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleasePduSessionResponse struct {
	BinaryDataN4InformationExt1 []byte        `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 []byte        `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *ReleasedData `json:"jsonData,omitempty"`
	BinaryDataN4Information     []byte        `json:"binaryDataN4Information,omitempty"`
}

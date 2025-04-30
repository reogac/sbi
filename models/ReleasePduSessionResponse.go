/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleasePduSessionResponse struct {
	BinaryDataN4InformationExt2 []byte        `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *ReleasedData `json:"jsonData,omitempty"`
	BinaryDataN4Information     []byte        `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 []byte        `json:"binaryDataN4InformationExt1,omitempty"`
}

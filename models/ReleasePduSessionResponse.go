/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleasePduSessionResponse struct {
	JsonData                    *ReleasedData `json:"jsonData,omitempty"`
	BinaryDataN4Information     []byte        `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 []byte        `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 []byte        `json:"binaryDataN4InformationExt2,omitempty"`
}

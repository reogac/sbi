/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleasePduSessionRequest struct {
	JsonData                    *ReleaseData `json:"jsonData,omitempty"`
	BinaryDataN4Information     []byte       `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 []byte       `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 []byte       `json:"binaryDataN4InformationExt2,omitempty"`
}

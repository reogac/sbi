/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UEContextTransferResponse struct {
	JsonData                    *UeContextTransferRspData `json:"jsonData,omitempty"`
	BinaryDataN2Information     []byte                    `json:"binaryDataN2Information,omitempty"`
	BinaryDataN2InformationExt1 []byte                    `json:"binaryDataN2InformationExt1,omitempty"`
	BinaryDataN2InformationExt2 []byte                    `json:"binaryDataN2InformationExt2,omitempty"`
}

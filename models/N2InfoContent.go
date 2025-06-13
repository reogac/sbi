/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InfoContent struct {
	NgapIeType      NgapIeType      `json:"ngapIeType,omitempty"`
	NgapData        RefToBinaryData `json:"ngapData"`
	NgapMessageType *int            `json:"ngapMessageType,omitempty"`
}

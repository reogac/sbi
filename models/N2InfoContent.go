/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InfoContent struct {
	NgapMessageType *int            `json:"ngapMessageType,omitempty"`
	NgapIeType      NgapIeType      `json:"ngapIeType,omitempty"`
	NgapData        RefToBinaryData `json:"ngapData"`
}

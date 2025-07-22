/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedundantTransmissionExpReq struct {
	RedTOrderCriter RedTransExpOrderingCriterion `json:"redTOrderCriter,omitempty"`
	Order           MatchingDirection            `json:"order,omitempty"`
}

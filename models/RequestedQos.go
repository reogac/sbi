/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:58 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RequestedQos struct {
	FiveQi int    `json:"5qi"`
	GbrUl  string `json:"gbrUl,omitempty"`
	GbrDl  string `json:"gbrDl,omitempty"`
}

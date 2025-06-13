/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RequestedQos struct {
	GbrUl  string `json:"gbrUl,omitempty"`
	GbrDl  string `json:"gbrDl,omitempty"`
	FiveQi int    `json:"5qi"`
}

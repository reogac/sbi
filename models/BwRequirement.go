/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BwRequirement struct {
	MirBwDl string `json:"mirBwDl,omitempty"`
	MirBwUl string `json:"mirBwUl,omitempty"`
	AppId   string `json:"appId"`
	MarBwDl string `json:"marBwDl,omitempty"`
	MarBwUl string `json:"marBwUl,omitempty"`
}

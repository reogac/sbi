/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BwRequirement struct {
	AppId   string `json:"appId"`
	MarBwDl string `json:"marBwDl,omitempty"`
	MarBwUl string `json:"marBwUl,omitempty"`
	MirBwDl string `json:"mirBwDl,omitempty"`
	MirBwUl string `json:"mirBwUl,omitempty"`
}

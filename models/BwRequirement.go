/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BwRequirement struct {
	MirBwUl string `json:"mirBwUl,omitempty"`
	AppId   string `json:"appId"`
	MarBwDl string `json:"marBwDl,omitempty"`
	MarBwUl string `json:"marBwUl,omitempty"`
	MirBwDl string `json:"mirBwDl,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSetConfiguration struct {
	SetId          string              `json:"setId"`
	SupportedPlmns []SupportedPlmnItem `json:"supportedPlmns,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSetConfiguration struct {
	SetId          string              `json:"setId"`
	SupportedPlmns []SupportedPlmnItem `json:"supportedPlmns,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UEContextRelease struct {
	Supi                string    `json:"supi,omitempty"`
	UnauthenticatedSupi *bool     `json:"unauthenticatedSupi,omitempty"`
	NgapCause           NgApCause `json:"ngapCause"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CagInfo struct {
	CagOnlyIndicator *bool    `json:"cagOnlyIndicator,omitempty"`
	AllowedCagList   []string `json:"allowedCagList"`
}

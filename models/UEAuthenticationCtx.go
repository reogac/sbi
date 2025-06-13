/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UEAuthenticationCtx struct {
	AuthType           AuthType        `json:"authType"`
	FiveGAuthData      FiveGAuthData   `json:"5gAuthData"`
	Links              map[string]Link `json:"_links"`
	ServingNetworkName string          `json:"servingNetworkName,omitempty"`
}

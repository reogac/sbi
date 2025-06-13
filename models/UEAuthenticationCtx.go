/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UEAuthenticationCtx struct {
	ServingNetworkName string          `json:"servingNetworkName,omitempty"`
	AuthType           AuthType        `json:"authType"`
	FiveGAuthData      FiveGAuthData   `json:"5gAuthData"`
	Links              map[string]Link `json:"_links"`
}

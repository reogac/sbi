/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthEvent struct {
	AuthType                   AuthType `json:"authType"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	Success                    bool     `json:"success"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
}

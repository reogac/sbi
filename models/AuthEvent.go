/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthEvent struct {
	NfInstanceId               string   `json:"nfInstanceId"`
	AuthType                   AuthType `json:"authType"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	Success                    bool     `json:"success"`
	TimeStamp                  string   `json:"timeStamp"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
}

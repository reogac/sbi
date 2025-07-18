/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthEvent struct {
	ResetIds                   []string `json:"resetIds,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	Success                    bool     `json:"success"`
	TimeStamp                  string   `json:"timeStamp"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	AuthType                   AuthType `json:"authType"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
}

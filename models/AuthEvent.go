/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthEvent struct {
	AuthType                   AuthType `json:"authType"`
	TimeStamp                  string   `json:"timeStamp"`
	Success                    bool     `json:"success"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
}

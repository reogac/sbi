/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Apr 29 09:34:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	Scope                string      `json:"scope"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	NfInstanceId         string      `json:"nfInstanceId"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	NfType               NFType      `json:"nfType,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
}

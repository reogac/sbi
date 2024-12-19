/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:26:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	Scope                string      `json:"scope"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	NfType               NFType      `json:"nfType,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
}

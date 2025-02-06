/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Feb  6 13:48:14 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	NfType               NFType      `json:"nfType,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	Scope                string      `json:"scope"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	NfType               NFType      `json:"nfType,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	Scope                string      `json:"scope"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
}

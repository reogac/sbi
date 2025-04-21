/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Apr 21 15:04:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	GrantType            GrantType   `json:"grant_type"`
	NfType               NFType      `json:"nfType,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	Scope                string      `json:"scope"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
}

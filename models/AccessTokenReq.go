/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	NfType               NFType      `json:"nfType,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	Scope                string      `json:"scope"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
}

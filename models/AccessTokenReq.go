/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Jul 21 09:39:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	GrantType            GrantType   `json:"grant_type"`
	Scope                string      `json:"scope"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	NfType               NFType      `json:"nfType,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	GrantType            GrantType   `json:"grant_type"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	Scope                string      `json:"scope"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	NfType               NFType      `json:"nfType,omitempty"`
}

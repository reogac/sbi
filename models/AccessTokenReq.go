/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	GrantType            GrantType   `json:"grant_type"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	NfType               NFType      `json:"nfType,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	Scope                string      `json:"scope"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
}

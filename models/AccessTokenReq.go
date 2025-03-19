/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Mar 19 09:38:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	GrantType            GrantType   `json:"grant_type"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	NfType               NFType      `json:"nfType,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	Scope                string      `json:"scope"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
}

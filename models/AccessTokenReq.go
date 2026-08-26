/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 11:15:54 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessTokenReq struct {
	NfType               NFType      `json:"nfType,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	Scope                string      `json:"scope"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
}

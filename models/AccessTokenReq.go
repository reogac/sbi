/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:05:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models
type AccessTokenReq struct {
	 TargetNfType	NFType	`json:"targetNfType,omitempty"`
	 RequesterSnssaiList	[]Snssai	`json:"requesterSnssaiList,omitempty"`
	 Scope	string	`json:"scope"`
	 RequesterPlmn	*PlmnId	`json:"requesterPlmn,omitempty"`
	 TargetPlmn	*PlmnId	`json:"targetPlmn,omitempty"`
	 TargetNsiList	[]string	`json:"targetNsiList,omitempty"`
	 HnrfAccessTokenUri	string	`json:"hnrfAccessTokenUri,omitempty"`
	 GrantType	GrantType	`json:"grant_type"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 RequesterPlmnList	[]PlmnId	`json:"requesterPlmnList,omitempty"`
	 RequesterSnpnList	[]PlmnIdNid	`json:"requesterSnpnList,omitempty"`
	 TargetSnpn	*PlmnIdNid	`json:"targetSnpn,omitempty"`
	 TargetSnssaiList	[]Snssai	`json:"targetSnssaiList,omitempty"`
	 TargetNfServiceSetId	string	`json:"targetNfServiceSetId,omitempty"`
	 NfType	NFType	`json:"nfType,omitempty"`
	 TargetNfInstanceId	string	`json:"targetNfInstanceId,omitempty"`
	 SourceNfInstanceId	string	`json:"sourceNfInstanceId,omitempty"`
	 RequesterFqdn	string	`json:"requesterFqdn,omitempty"`
	 TargetNfSetId	string	`json:"targetNfSetId,omitempty"`
}

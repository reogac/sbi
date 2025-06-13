/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	UeContext           UeContext          `json:"ueContext"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	InitialAmfName      string             `json:"initialAmfName"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnType              AccessType         `json:"anType"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
}

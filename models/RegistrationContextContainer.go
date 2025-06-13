/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	UeContext           UeContext          `json:"ueContext"`
	AnType              AccessType         `json:"anType"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	UserLocation        UserLocation       `json:"userLocation"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
}

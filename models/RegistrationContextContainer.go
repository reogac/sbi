/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	AnType              AccessType         `json:"anType"`
	UserLocation        UserLocation       `json:"userLocation"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	InitialAmfName      string             `json:"initialAmfName"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
}

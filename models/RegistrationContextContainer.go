/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	InitialAmfName      string             `json:"initialAmfName"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	UserLocation        UserLocation       `json:"userLocation"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	AnType              AccessType         `json:"anType"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
}

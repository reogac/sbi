/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
	UserLocation        UserLocation       `json:"userLocation"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	AnType              AccessType         `json:"anType"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnN2ApId            int                `json:"anN2ApId"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
}

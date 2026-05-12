/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	InitialAmfName      string             `json:"initialAmfName"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnType              AccessType         `json:"anType"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
}

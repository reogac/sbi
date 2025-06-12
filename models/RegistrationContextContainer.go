/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	InitialAmfName      string             `json:"initialAmfName"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	AnType              AccessType         `json:"anType"`
	AnN2ApId            int                `json:"anN2ApId"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
}

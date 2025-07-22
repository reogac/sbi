/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationContextContainer struct {
	AnN2IPv4Addr        string             `json:"anN2IPv4Addr,omitempty"`
	ConfiguredNssai     []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInTa   []Snssai           `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId      *PlmnId            `json:"selectedPlmnId,omitempty"`
	CeModeBInd          *CeModeBInd        `json:"ceModeBInd,omitempty"`
	AuthenticatedInd    *bool              `json:"authenticatedInd,omitempty"`
	RanNodeId           GlobalRanNodeId    `json:"ranNodeId"`
	RrcEstCause         string             `json:"rrcEstCause,omitempty"`
	NpnAccessInfo       *NpnAccessInfo     `json:"npnAccessInfo,omitempty"`
	RejectedNssaiInPlmn []Snssai           `json:"rejectedNssaiInPlmn,omitempty"`
	LteMInd             *LteMInd           `json:"lteMInd,omitempty"`
	UserLocation        UserLocation       `json:"userLocation"`
	AnN2IPv6Addr        string             `json:"anN2IPv6Addr,omitempty"`
	InitialAmfName      string             `json:"initialAmfName"`
	UeContextRequest    *bool              `json:"ueContextRequest,omitempty"`
	IabNodeInd          *bool              `json:"iabNodeInd,omitempty"`
	LocalTimeZone       string             `json:"localTimeZone,omitempty"`
	AnN2ApId            int                `json:"anN2ApId"`
	InitialAmfN2ApId    *int               `json:"initialAmfN2ApId,omitempty"`
	AllowedNssai        *AllowedNssai      `json:"allowedNssai,omitempty"`
	UeContext           UeContext          `json:"ueContext"`
	AnType              AccessType         `json:"anType"`
}

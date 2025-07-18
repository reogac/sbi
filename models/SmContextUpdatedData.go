/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdatedData struct {
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
}

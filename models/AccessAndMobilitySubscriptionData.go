/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
}

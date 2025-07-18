/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
}

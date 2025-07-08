/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
}

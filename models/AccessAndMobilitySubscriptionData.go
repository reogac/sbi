/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
}

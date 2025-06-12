/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	Dnn                             string                          `json:"dnn"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
}

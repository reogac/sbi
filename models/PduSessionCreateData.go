/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
}

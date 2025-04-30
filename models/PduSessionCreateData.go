/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
}

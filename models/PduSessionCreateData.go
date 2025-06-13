/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	Dnn                             string                          `json:"dnn"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
}

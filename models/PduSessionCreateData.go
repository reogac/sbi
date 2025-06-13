/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	Dnn                             string                          `json:"dnn"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
}

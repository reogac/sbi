/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	Dnn                             string                          `json:"dnn"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	AnType                          AccessType                      `json:"anType"`
}

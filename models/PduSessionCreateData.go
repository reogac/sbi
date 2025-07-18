/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	Dnn                             string                          `json:"dnn"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
}

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	Supi                    string                    `json:"supi"`
	Online                  *bool                     `json:"online,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	Dnn                     string                    `json:"dnn"`
	NotificationUri         string                    `json:"notificationUri"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
}

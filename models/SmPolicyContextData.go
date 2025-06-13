/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	Supi                    string                    `json:"supi"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	Dnn                     string                    `json:"dnn"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
}

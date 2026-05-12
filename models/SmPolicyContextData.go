/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	SmfId                   string                    `json:"smfId,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	Supi                    string                    `json:"supi"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	NotificationUri         string                    `json:"notificationUri"`
	Pei                     string                    `json:"pei,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	Dnn                     string                    `json:"dnn"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
}

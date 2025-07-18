/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	Supi                    string                    `json:"supi"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	Dnn                     string                    `json:"dnn"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
}

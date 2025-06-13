/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	Dnn                     string                    `json:"dnn"`
	NotificationUri         string                    `json:"notificationUri"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	Supi                    string                    `json:"supi"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
}

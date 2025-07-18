/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
}

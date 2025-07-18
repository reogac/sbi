/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MdtConfiguration struct {
	EventThresholdRsrpNr     *int                      `json:"eventThresholdRsrpNr,omitempty"`
	PositioningMethod        PositioningMethodMdt      `json:"positioningMethod,omitempty"`
	EventThresholdRsrp       *int                      `json:"eventThresholdRsrp,omitempty"`
	MdtAllowedPlmnIdList     []PlmnId                  `json:"mdtAllowedPlmnIdList,omitempty"`
	JobType                  JobType                   `json:"jobType"`
	MeasurementLteList       []string                  `json:"measurementLteList,omitempty"`
	AddPositioningMethodList []string                  `json:"addPositioningMethodList,omitempty"`
	ReportType               ReportTypeMdt             `json:"reportType,omitempty"`
	AreaScope                *AreaScope                `json:"areaScope,omitempty"`
	ReportInterval           ReportIntervalMdt         `json:"reportInterval,omitempty"`
	ReportAmount             ReportAmountMdt           `json:"reportAmount,omitempty"`
	CollectionPeriodRmmNr    CollectionPeriodRmmNrMdt  `json:"collectionPeriodRmmNr,omitempty"`
	InterFreqTargetList      []InterFreqTargetInfo     `json:"interFreqTargetList,omitempty"`
	EventList                []string                  `json:"eventList,omitempty"`
	LoggingDurationNr        LoggingDurationNrMdt      `json:"loggingDurationNr,omitempty"`
	CollectionPeriodRmmLte   CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	MeasurementPeriodLte     MeasurementPeriodLteMdt   `json:"measurementPeriodLte,omitempty"`
	MeasurementNrList        []string                  `json:"measurementNrList,omitempty"`
	ReportingTriggerList     []string                  `json:"reportingTriggerList,omitempty"`
	EventThresholdRsrq       *int                      `json:"eventThresholdRsrq,omitempty"`
	LoggingDuration          LoggingDurationMdt        `json:"loggingDuration,omitempty"`
	MbsfnAreaList            []MbsfnArea               `json:"mbsfnAreaList,omitempty"`
	SensorMeasurementList    []string                  `json:"sensorMeasurementList,omitempty"`
	ReportIntervalNr         ReportIntervalNrMdt       `json:"reportIntervalNr,omitempty"`
	EventThresholdRsrqNr     *int                      `json:"eventThresholdRsrqNr,omitempty"`
	LoggingInterval          LoggingIntervalMdt        `json:"loggingInterval,omitempty"`
	LoggingIntervalNr        LoggingIntervalNrMdt      `json:"loggingIntervalNr,omitempty"`
}

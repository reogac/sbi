/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type JobType string

// Define constant values for JobType
const (
	JOBTYPE_IMMEDIATE_MDT_ONLY      JobType = "IMMEDIATE_MDT_ONLY"
	JOBTYPE_LOGGED_MDT_ONLY         JobType = "LOGGED_MDT_ONLY"
	JOBTYPE_TRACE_ONLY              JobType = "TRACE_ONLY"
	JOBTYPE_IMMEDIATE_MDT_AND_TRACE JobType = "IMMEDIATE_MDT_AND_TRACE"
	JOBTYPE_RLF_REPORTS_ONLY        JobType = "RLF_REPORTS_ONLY"
	JOBTYPE_RCEF_REPORTS_ONLY       JobType = "RCEF_REPORTS_ONLY"
	JOBTYPE_LOGGED_MBSFN_MDT        JobType = "LOGGED_MBSFN_MDT"
)

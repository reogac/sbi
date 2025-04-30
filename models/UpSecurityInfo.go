/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpSecurityInfo struct {
	UpSecurity                      UpSecurity                    `json:"upSecurity"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SecurityResult                  *SecurityResult               `json:"securityResult,omitempty"`
}

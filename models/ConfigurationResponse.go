/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfigurationResponse struct {
	Smf     *SessionManagementConfiguration `json:"smf,omitempty"`
	Upf     *UserPlaneConfiguration         `json:"upf,omitempty"`
	Udm     *UdmConfiguration               `json:"udm,omitempty"`
	Udr     *UdrConfiguration               `json:"udr,omitempty"`
	Sepp    *SeppConfiguration              `json:"sepp,omitempty"`
	Slicing *SlicingConfiguration           `json:"slicing,omitempty"`
}

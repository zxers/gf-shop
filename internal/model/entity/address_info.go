// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddressInfo is the golang structure for table address_info.
type AddressInfo struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:""`
	Pid       int         `json:"pid"       description:""`
	Status    int         `json:"status"    description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}

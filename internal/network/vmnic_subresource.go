/*
 *  Copyright 2023 VMware, Inc.
 *    SPDX-License-Identifier: MPL-2.0
 */

package network

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	validationutils "github.com/vmware/terraform-provider-vcf/internal/validation"
	"github.com/vmware/vcf-sdk-go/models"
)

// VMNicSchema this helper function extracts the VMNic Schema, so that
// it's made available for both workload domain and cluster creation.
func VMNicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "ESXI host vmnic ID to be associated with a VDS, once added to cluster",
				ValidateFunc: validation.NoZeroValues,
			},
			"uplink": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Uplink to be associated with vmnic",
				ValidateFunc: validation.NoZeroValues,
			},
			"vds_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Name of the VDS to associate with the ESXi host",
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

func TryConvertToVmNic(object map[string]interface{}) (*models.VMNic, error) {
	if object == nil {
		return nil, fmt.Errorf("cannot convert to VMNic, object is nil")
	}
	id := object["id"].(string)
	if len(id) == 0 {
		return nil, fmt.Errorf("cannot convert to VMNic, id is required")
	}
	result := &models.VMNic{}
	result.ID = id
	if uplink, ok := object["uplink"]; ok && !validationutils.IsEmpty(uplink) {
		result.Uplink = uplink.(string)
	}
	if vdsName, ok := object["vds_name"]; ok && !validationutils.IsEmpty(vdsName) {
		result.VdsName = vdsName.(string)
	}
	return result, nil
}

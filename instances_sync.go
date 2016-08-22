package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
)

type InstancesSync struct {
	D      *schema.ResourceData
	Client BareMetalClient
	Res    *baremetal.InstanceList
}

func (s *InstancesSync) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(
		s.D,
		"availability_domain",
		"page",
		"limit",
	)

	s.Res, e = s.Client.ListInstances(compartmentID, opts...)
	return

}

func (s *InstancesSync) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Instances {
			res := map[string]interface{}{
				"availability_domain": v.AvailabilityDomain,
				"compartment_id":      v.CompartmentID,
				"display_name":        v.DisplayName,
				"id":                  v.ID,
				"image":               v.Image,
				"metadata":            v.Metadata,
				"region":              v.Region,
				"shape":               v.Shape,
				"state":               v.State,
				"time_created":        v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("instances", resources)
	}
	return
}

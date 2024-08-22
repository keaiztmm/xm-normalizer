package elastic_container

type ElasticContainer struct {
	CloudPlatform string `json:"cloud_platform"`
	CloudService  string `json:"cloud_service"`

	RegionID string `json:"region_id"`
	ZoneID   string `json:"zone_id"`

	InstanceID   string `json:"id"`
	InstanceName string `json:"name"`
}

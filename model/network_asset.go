package model

import "time"

type NetworkAsset struct {
	Name             string     `json:"name" db:"name"`
	SystemName       string     `json:"system_name" db:"systemname"`
	Address          string     `json:"address" db:"address"`
	ShortDescription string     `json:"short_description" db:"shortdescription"`
	SubnetMask       string     `json:"subnet_mask" db:"subnetmask"`
	ProtocolType     string     `json:"protocol_type" db:"protocoltype"`
	Description      string     `json:"description" db:"description"`
	AddressType      string     `json:"address_type" db:"addresstype"`
	DNSHostName      string     `json:"dns_host_name" db:"dnshostname"`
	CreateDate       time.Time  `json:"create_date" db:"createdate"`
	DatasetId        int        `json:"dataset_id" db:"datasetid"`
	ModifiedDate     *time.Time `json:"modified_date" db:"modifieddate"`
	LastModifiedBy   string     `json:"last_modified_by" db:"lastmodifiedby"`
	InstanceId       string     `json:"instance_id" db:"instanceid"`
	RequestId        string     `json:"request_id" db:"requestid"`
}

type NetworkAssetList struct {
	Name             string    `json:"name" db:"name"`
	SystemName       string    `json:"system_name" db:"systemname"`
	Address          string    `json:"address" db:"address"`
	ShortDescription string    `json:"short_description" db:"shortdescription"`
	ProtocolType     string    `json:"protocol_type" db:"protocoltype"`
	AddressType      string    `json:"address_type" db:"addresstype"`
	DNSHostName      string    `json:"dns_host_name" db:"dnshostname"`
	CreateDate       time.Time `json:"create_date" db:"createdate"`
}
type NetworkAssetFilter struct {
	Name         string `json:"name,omitempty" query:"name"`
	Address      string `json:"address,omitempty" query:"address"`
	ProtocolType string `json:"protocol_type,omitempty" query:"protocol_type"`
	AddressType  string `json:"address_type,omitempty" query:"address_type"`
	DnsHostname  string `json:"dns_host_name,omitempty" query:"dns_host_name"`
	DatasetId    int    `json:"dataset_id,omitempty" query:"dataset_id"`
	Page         int    `json:"page" query:"page"`
	Limit        int    `json:"limit" query:"limit"`
}

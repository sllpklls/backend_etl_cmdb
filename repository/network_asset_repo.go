package repository

import (
	"context"

	"github.com/sllpklls/template-backend-go/model"
)

type NetworkAssetRepo interface {
	GetAllNetworkAssets(ctx context.Context, page, limit int) ([]model.NetworkAssetList, error)
	GetNetworkAssetByName(ctx context.Context, name string) (*model.NetworkAsset, error)
	GetNetworkAssetsByFilter(ctx context.Context, filter model.NetworkAssetFilter) ([]model.NetworkAssetList, error)
	GetNetworkAssetsByDNSHostName(ctx context.Context, dnsHostName string, page, limit int) ([]model.NetworkAssetList, error)
	GetTotalNetworkAssetsByDNSHostName(ctx context.Context, dnsHostName string) (int, error)
	GetTotalNetworkAssets(ctx context.Context) (int, error)
	GetTotalNetworkAssetsByFilter(ctx context.Context, filter model.NetworkAssetFilter) (int, error)
	CreateNetworkAsset(ctx context.Context, asset model.NetworkAsset) error
	UpdateNetworkAsset(ctx context.Context, name string, asset model.NetworkAsset) error
	DeleteNetworkAsset(ctx context.Context, name string) error

	GetIPEndpointByDNSHostName(ctx context.Context, dnsHostName string) (bool, error)
}

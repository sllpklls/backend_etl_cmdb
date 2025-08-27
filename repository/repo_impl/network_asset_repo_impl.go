package repo_impl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/sllpklls/template-backend-go/db"
	"github.com/sllpklls/template-backend-go/model"
)

type NetworkAssetRepoImpl struct {
	sql *db.Sql
}

func NewNetworkAssetRepo(sql *db.Sql) *NetworkAssetRepoImpl {
	return &NetworkAssetRepoImpl{sql: sql}
}

func (r *NetworkAssetRepoImpl) GetAllNetworkAssets(ctx context.Context, page, limit int) ([]model.NetworkAssetList, error) {
	offset := (page - 1) * limit

	query := `
		SELECT name, systemname, address, shortdescription, protocoltype, 
		       addresstype, dnshostname, createdate
		FROM NetworkAssets 
		ORDER BY createdate DESC 
		LIMIT $1 OFFSET $2`

	rows, err := r.sql.Db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query network assets: %w", err)
	}
	defer rows.Close()

	var assets []model.NetworkAssetList
	for rows.Next() {
		var asset model.NetworkAssetList
		err := rows.Scan(
			&asset.Name,
			&asset.SystemName,
			&asset.Address,
			&asset.ShortDescription,
			&asset.ProtocolType,
			&asset.AddressType,
			&asset.DNSHostName,
			&asset.CreateDate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan network asset: %w", err)
		}
		assets = append(assets, asset)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return assets, nil
}

func (r *NetworkAssetRepoImpl) GetNetworkAssetByName(ctx context.Context, name string) (*model.NetworkAsset, error) {
	query := `
		SELECT name, systemname, address, shortdescription, subnetmask, protocoltype, 
		       description, addresstype, dnshostname, createdate, datasetid, 
		       modifieddate, lastmodifiedby, instanceid, requestid
		FROM NetworkAssets 
		WHERE name = $1`

	var asset model.NetworkAsset
	row := r.sql.Db.QueryRowContext(ctx, query, name)

	err := row.Scan(
		&asset.Name,
		&asset.SystemName,
		&asset.Address,
		&asset.ShortDescription,
		&asset.SubnetMask,
		&asset.ProtocolType,
		&asset.Description,
		&asset.AddressType,
		&asset.DNSHostName,
		&asset.CreateDate,
		&asset.DatasetId,
		&asset.ModifiedDate,
		&asset.LastModifiedBy,
		&asset.InstanceId,
		&asset.RequestId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get network asset: %w", err)
	}

	return &asset, nil
}

func (r *NetworkAssetRepoImpl) GetNetworkAssetsByDNSHostName(ctx context.Context, dnsHostName string, page, limit int) ([]model.NetworkAssetList, error) {
	offset := (page - 1) * limit

	query := `
		SELECT name, systemname, address, shortdescription, protocoltype, 
		       addresstype, dnshostname, createdate
		FROM NetworkAssets 
		WHERE dnshostname ILIKE $1
		ORDER BY createdate DESC 
		LIMIT $2 OFFSET $3`

	rows, err := r.sql.Db.QueryContext(ctx, query, "%"+dnsHostName+"%", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query network assets by DNS hostname: %w", err)
	}
	defer rows.Close()

	var assets []model.NetworkAssetList
	for rows.Next() {
		var asset model.NetworkAssetList
		err := rows.Scan(
			&asset.Name,
			&asset.SystemName,
			&asset.Address,
			&asset.ShortDescription,
			&asset.ProtocolType,
			&asset.AddressType,
			&asset.DNSHostName,
			&asset.CreateDate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan network asset: %w", err)
		}
		assets = append(assets, asset)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return assets, nil
}
func (r *NetworkAssetRepoImpl) GetTotalNetworkAssetsByDNSHostName(ctx context.Context, dnsHostName string) (int, error) {
	query := "SELECT COUNT(*) FROM NetworkAssets WHERE dnshostname ILIKE $1"

	var total int
	err := r.sql.Db.QueryRowContext(ctx, query, "%"+dnsHostName+"%").Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total network assets by DNS hostname: %w", err)
	}

	return total, nil
}
func (r *NetworkAssetRepoImpl) GetIPEndpointByDNSHostName(ctx context.Context, dnsHostName string) (bool, error) {
	query := `
		SELECT 1
		FROM NetworkAssets 
		WHERE dnshostname = $1
		LIMIT 1`

	var exists int
	err := r.sql.Db.QueryRowContext(ctx, query, dnsHostName).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to query network assets by DNS hostname: %w", err)
	}

	return true, nil
}

func (r *NetworkAssetRepoImpl) GetTotalNetworkAssets(ctx context.Context) (int, error) {
	var total int
	query := "SELECT COUNT(*) FROM NetworkAssets"

	err := r.sql.Db.QueryRowContext(ctx, query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total network assets: %w", err)
	}

	return total, nil
}

func (r *NetworkAssetRepoImpl) GetTotalNetworkAssetsByFilter(ctx context.Context, filter model.NetworkAssetFilter) (int, error) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	baseQuery := "SELECT COUNT(*) FROM NetworkAssets"

	if filter.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Name+"%")
		argIndex++
	}

	if filter.Address != "" {
		conditions = append(conditions, fmt.Sprintf("address ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Address+"%")
		argIndex++
	}

	if filter.ProtocolType != "" {
		conditions = append(conditions, fmt.Sprintf("protocoltype = $%d", argIndex))
		args = append(args, filter.ProtocolType)
		argIndex++
	}

	if filter.AddressType != "" {
		conditions = append(conditions, fmt.Sprintf("addresstype = $%d", argIndex))
		args = append(args, filter.AddressType)
		argIndex++
	}

	if filter.DatasetId > 0 {
		conditions = append(conditions, fmt.Sprintf("datasetid = $%d", argIndex))
		args = append(args, filter.DatasetId)
		argIndex++
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	var total int
	err := r.sql.Db.QueryRowContext(ctx, baseQuery, args...).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total network assets by filter: %w", err)
	}
	// log.Info(total)
	return total, nil
}
func (r *NetworkAssetRepoImpl) GetNetworkAssetsByFilter(ctx context.Context, filter model.NetworkAssetFilter) ([]model.NetworkAssetList, error) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	baseQuery := `
		SELECT name, systemname, address, shortdescription, protocoltype, 
		       addresstype, dnshostname, createdate
		FROM NetworkAssets`

	if filter.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Name+"%")
		argIndex++
	}

	if filter.Address != "" {
		conditions = append(conditions, fmt.Sprintf("address ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Address+"%")
		argIndex++
	}

	if filter.ProtocolType != "" {
		conditions = append(conditions, fmt.Sprintf("protocoltype = $%d", argIndex))
		args = append(args, filter.ProtocolType)
		argIndex++
	}

	if filter.AddressType != "" {
		conditions = append(conditions, fmt.Sprintf("addresstype = $%d", argIndex))
		args = append(args, filter.AddressType)
		argIndex++
	}
	if filter.DnsHostname != "" {
		conditions = append(conditions, fmt.Sprintf("dnshostname ILIKE $%d", argIndex))
		args = append(args, "%"+filter.DnsHostname+"%")
		argIndex++
	}
	if filter.DatasetId > 0 {
		conditions = append(conditions, fmt.Sprintf("datasetid = $%d", argIndex))
		args = append(args, filter.DatasetId)
		argIndex++
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	baseQuery += " ORDER BY createdate DESC"

	rows, err := r.sql.Db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query network assets with filter: %w", err)
	}
	defer rows.Close()

	var assets []model.NetworkAssetList
	for rows.Next() {
		var asset model.NetworkAssetList
		err := rows.Scan(
			&asset.Name,
			&asset.SystemName,
			&asset.Address,
			&asset.ShortDescription,
			&asset.ProtocolType,
			&asset.AddressType,
			&asset.DNSHostName,
			&asset.CreateDate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan network asset: %w", err)
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (r *NetworkAssetRepoImpl) CreateNetworkAsset(ctx context.Context, asset model.NetworkAsset) error {
	query := `
		INSERT INTO NetworkAssets (
			name, systemname, address, shortdescription, subnetmask, protocoltype,
			description, addresstype, dnshostname, datasetid, lastmodifiedby,
			instanceid, requestid
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := r.sql.Db.ExecContext(ctx, query,
		asset.Name,
		asset.SystemName,
		asset.Address,
		asset.ShortDescription,
		asset.SubnetMask,
		asset.ProtocolType,
		asset.Description,
		asset.AddressType,
		asset.DNSHostName,
		asset.DatasetId,
		asset.LastModifiedBy,
		asset.InstanceId,
		asset.RequestId,
	)

	if err != nil {
		return fmt.Errorf("failed to create network asset: %w", err)
	}

	return nil
}

func (r *NetworkAssetRepoImpl) UpdateNetworkAsset(ctx context.Context, name string, asset model.NetworkAsset) error {
	query := `
		UPDATE NetworkAssets SET
			systemname = $1, address = $2, shortdescription = $3, subnetmask = $4,
			protocoltype = $5, description = $6, addresstype = $7, dnshostname = $8,
			datasetid = $9, modifieddate = NOW(), lastmodifiedby = $10,
			instanceid = $11, requestid = $12
		WHERE name = $13`

	result, err := r.sql.Db.ExecContext(ctx, query,
		asset.SystemName,
		asset.Address,
		asset.ShortDescription,
		asset.SubnetMask,
		asset.ProtocolType,
		asset.Description,
		asset.AddressType,
		asset.DNSHostName,
		asset.DatasetId,
		asset.LastModifiedBy,
		asset.InstanceId,
		asset.RequestId,
		name,
	)

	if err != nil {
		return fmt.Errorf("failed to update network asset: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("network asset not found")
	}

	return nil
}

func (r *NetworkAssetRepoImpl) DeleteNetworkAsset(ctx context.Context, name string) error {
	query := "DELETE FROM NetworkAssets WHERE name = $1"

	result, err := r.sql.Db.ExecContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("failed to delete network asset: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("network asset not found")
	}

	return nil
}

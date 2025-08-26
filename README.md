# 1. Lấy danh sách Network Assets (có phân trang)
GET /api/v1/network-assets

Query params:
    page (int, default=1)
    limit (int, default=10, max=100)

curl "http://localhost:3000/api/v1/network-assets?page=1&limit=5"

# 2. Tìm kiếm Network Assets (lọc nhiều trường)
GET /api/v1/network-assets/search

Query params:
    name (string, tìm theo tên, LIKE)
    address (string)
    protocol_type (string, exact match)
    address_type (string)
    dns_host_name (string, LIKE)
    dataset_id (int)
    page (int, default=1)
    limit (int, default=10)

curl "http://localhost:3000/api/v1/network-assets/search?name=server&protocol_type=TCP&page=1&limit=5"

# 3. Tìm kiếm theo DNS Hostname
GET /api/v1/network-assets/search-dns?dns_host_name=example.com&page=1&limit=10

curl "http://localhost:3000/api/v1/network-assets/search-dns?dns_host_name=myserver"

# 4. Lấy chi tiết Network Asset theo name
GET /api/v1/network-assets/:name

curl "http://localhost:3000/api/v1/network-assets/myserver01"

# 5. Tạo mới Network Asset
POST /api/v1/network-assets

Body Json:
{
  "name": "myserver01",
  "system_name": "Server System",
  "address": "192.168.1.10",
  "short_description": "Test server",
  "subnet_mask": "255.255.255.0",
  "protocol_type": "TCP",
  "description": "Internal server",
  "address_type": "IPv4",
  "dns_host_name": "myserver.local",
  "dataset_id": 1,
  "last_modified_by": "admin",
  "instance_id": "inst-001",
  "request_id": "req-001"
}

curl -X POST "http://localhost:3000/api/v1/network-assets" \
-H "Content-Type: application/json" \
-d '{"name":"myserver01","address":"192.168.1.10"}'
 # 6. Cập nhật Network Asset
 PUT /api/v1/network-assets/:name

 Body Json:
{
  "name": "myserver01",
  "system_name": "Server System",
  "address": "192.168.1.10",
  "short_description": "Test server",
  "subnet_mask": "255.255.255.0",
  "protocol_type": "TCP",
  "description": "Internal server",
  "address_type": "IPv4",
  "dns_host_name": "myserver.local",
  "dataset_id": 1,
  "last_modified_by": "admin",
  "instance_id": "inst-001",
  "request_id": "req-001"
}

curl -X PUT "http://localhost:3000/api/v1/network-assets/myserver01" \
-H "Content-Type: application/json" \
-d '{"system_name":"Updated System","address":"192.168.1.20"}'

# 7. Xóa Network Asset
DELETE /api/v1/network-assets/:name

curl -X DELETE "http://localhost:3000/api/v1/network-assets/myserver01"

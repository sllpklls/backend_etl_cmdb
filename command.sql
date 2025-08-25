CREATE TABLE NetworkAssets (
    Name VARCHAR(50),              -- Tên IP
    SystemName VARCHAR(100),       -- Tên hệ thống hoặc hostname
    Address VARCHAR(50),           -- Địa chỉ IP
    ShortDescription VARCHAR(255), -- Mô tả ngắn
    SubnetMask VARCHAR(50),        -- Địa chỉ mask
    ProtocolType VARCHAR(20),      -- Loại giao thức (TCP, UDP, ICMP...)
    Description TEXT,               -- Mô tả chi tiết
    AddressType VARCHAR(50),        -- IPv4, IPv6
    DNSHostName VARCHAR(100),       -- Tên DNS
    CreateDate TIMESTAMPTZ DEFAULT NOW(),  -- Ngày tạo
    DatasetId INT,                  -- Data Set ID
    ModifiedDate TIMESTAMPTZ,       -- Ngày chỉnh sửa cuối
    LastModifiedBy VARCHAR(100),    -- Người chỉnh sửa cuối
    InstanceId VARCHAR(100),        -- ID instance
    RequestId VARCHAR(100)          -- ID phiếu
);



INSERT INTO NetworkAssets
(Name, SystemName, Address, ShortDescription, SubnetMask, ProtocolType, Description, AddressType, DNSHostName, CreateDate, DatasetId, ModifiedDate, LastModifiedBy, InstanceId, RequestId)
VALUES
-- Web & DB
('WebServer01', 'srv-web-01', '192.168.1.10', 'Web server chính', '255.255.255.0', 'TCP', 'Chạy ứng dụng web nội bộ', 'IPv4', 'web01.local', '2025-08-01 10:00:00', 1001, '2025-08-10 12:30:00', 'admin', 'inst-001', 'req-123'),
('DBServer01', 'srv-db-01', '192.168.1.20', 'CSDL nội bộ', '255.255.255.0', 'TCP', 'Máy chủ PostgreSQL', 'IPv4', 'db01.local', '2025-08-02 11:00:00', 1002, '2025-08-11 09:45:00', 'dba', 'inst-002', 'req-124'),

-- Hạ tầng mạng
('Firewall01', 'fw-01', '10.0.0.1', 'Firewall chính', '255.255.0.0', 'UDP', 'Thiết bị tường lửa', 'IPv4', 'fw01.local', '2025-08-05 08:30:00', 1003, '2025-08-12 15:20:00', 'security', 'inst-003', 'req-125'),
('RouterCore01', 'rtr-core-01', '10.0.0.254', 'Router core mạng LAN', '255.255.0.0', 'ICMP', 'Router trung tâm cho toàn bộ LAN', 'IPv4', 'core01.local', '2025-08-06 07:50:00', 1005, '2025-08-14 19:40:00', 'netadmin', 'inst-005', 'req-128'),

-- App
('AppServer01', 'srv-app-01', '2001:db8::1', 'Ứng dụng ERP', 'ffff:ffff:ffff:ffff::', 'TCP', 'Chạy ứng dụng ERP cho công ty', 'IPv6', 'erp.local', '2025-08-07 14:00:00', 1004, '2025-08-13 18:10:00', 'erp-admin', 'inst-004', 'req-126'),
('AppServer02', 'srv-app-02', '192.168.1.30', 'Ứng dụng CRM', '255.255.255.0', 'TCP', 'Chạy CRM phục vụ kinh doanh', 'IPv4', 'crm.local', '2025-08-08 09:20:00', 1006, '2025-08-14 13:25:00', 'crm-admin', 'inst-006', 'req-129'),

-- Proxy & Email
('Proxy01', 'pxy-01', '192.168.2.10', 'Proxy server', '255.255.255.0', 'TCP', 'Lọc web, caching', 'IPv4', 'proxy.local', '2025-08-09 15:10:00', 1007, '2025-08-15 11:50:00', 'netadmin', 'inst-007', 'req-130'),
('MailServer01', 'srv-mail-01', '192.168.2.20', 'Mail Exchange', '255.255.255.0', 'TCP', 'Mail server cho công ty', 'IPv4', 'mail.local', '2025-08-10 08:40:00', 1008, '2025-08-15 16:30:00', 'mail-admin', 'inst-008', 'req-131'),

-- Monitoring & Backup
('Monitor01', 'mon-01', '192.168.3.10', 'Zabbix server', '255.255.255.0', 'TCP', 'Giám sát hệ thống', 'IPv4', 'monitor.local', '2025-08-11 17:00:00', 1009, '2025-08-16 12:15:00', 'ops', 'inst-009', 'req-132'),
('Backup01', 'bkp-01', '192.168.3.20', 'Backup NAS', '255.255.255.0', 'TCP', 'Lưu trữ backup hệ thống', 'IPv4', 'backup.local', '2025-08-12 20:10:00', 1010, '2025-08-16 18:20:00', 'storage', 'inst-010', 'req-133'),

-- Test/Dev
('DevServer01', 'srv-dev-01', '192.168.4.10', 'Máy chủ phát triển', '255.255.255.0', 'TCP', 'Dùng cho đội dev test ứng dụng', 'IPv4', 'dev01.local', '2025-08-13 09:00:00', 1011, '2025-08-17 14:40:00', 'developer', 'inst-011', 'req-134'),
('TestServer01', 'srv-test-01', '192.168.4.20', 'Máy chủ test QA', '255.255.255.0', 'TCP', 'Test ứng dụng trước triển khai', 'IPv4', 'test01.local', '2025-08-14 10:20:00', 1012, '2025-08-18 08:25:00', 'qa', 'inst-012', 'req-135'),

-- Cloud & Container
('K8sMaster01', 'k8s-master-01', '172.16.1.10', 'Kubernetes Master', '255.255.0.0', 'TCP', 'Điều phối cluster k8s', 'IPv4', 'k8s-master.local', '2025-08-15 12:00:00', 1013, '2025-08-18 21:30:00', 'devops', 'inst-013', 'req-136'),
('DockerHost01', 'docker-host-01', '172.16.1.20', 'Docker Host', '255.255.0.0', 'TCP', 'Chạy container ứng dụng', 'IPv4', 'docker.local', '2025-08-16 13:10:00', 1014, '2025-08-19 11:15:00', 'devops', 'inst-014', 'req-137');

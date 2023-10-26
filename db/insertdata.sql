INSERT INTO
    device_location (
        serial_number,
        device_make_model,
        model,
        device_type,
        data_center,
        region,
        dc_location,
        device_location,
        device_row_number,
        device_rack_number,
        device_ru_number
    )
VALUES
    ('123ABC456', 'Oracle T5-4', 'T5-4', 'Server', 'Primary DC', 'Delhi', 'Sansad Marg', 'IDC2, 3rd Floor', 6, 4, '22U-26U'),
    ('789DEF123', 'Hitachi HCP-G10', 'HCP-G10', 'Object Storage', 'Primary DC', 'Delhi', 'Sansad Marg', 'IDC1, 2nd Floor', 5, 2, '10U-11U'),
    ('456xyz789', 'Cisco NEXUS-93108', 'NEXUS-93108', 'Switch', 'Disaster DR', 'Mumbai', 'Belapur', 'DC1, Ground Floor', 4, 3, '5U'),
    ('234KLM567', 'Brocade 6520', '6520', 'SAN Switch', 'Disaster DR', 'Mumbai', 'Belapur', 'DC1, Ground Floor', 4, 5, '8U-9U');


INSERT INTO
    device_amc_owner (
        serial_number,
        device_make_model,
        model,
        po_number,
        po_order_date,
        eosl_date,
        amc_start_date,
        amc_end_date,
        device_owner
    )
VALUES
    ('123ABC456', 'Oracle T8-4', 'T8-4', 'PNB/SADBA/2022/151', '2022-02-03', '2027-03-31', '2023-04-01', '2024-03-31', 'Mr. Abc'),
    ('789DEF123', 'Hitachi HCP-G10', 'HCP-G10', 'PNB/CISD/2023/02', '2023-01-04', '2024-04-10', '2023-04-01', '2024-03-31', 'Mr. Xyz'),
    ('456xyz789', 'Cisco NEXUS-93108', 'NEXUS-93108', 'PNB/SADBA/2022/153', '2022-05-18', '2027-03-31', '2023-04-01', '2024-03-31', 'Mr. Abc'),
    ('234KLM567', 'Brocade 6520', '6520', 'PNB/CISD/2023/04', '2023-01-20', '2024-04-10', '2023-04-01', '2024-03-31', 'Mr. Xyz');
INSERT INTO
    device_power (
        serial_number,
        device_make_model,
        model,
        device_type,
        total_power_watt,
        total_btu,
        total_power_cable,
        power_socket_type
    )
VALUES
    ('123ABC456', 'Oracle T8-4', 'T8-4', 'Server', 1100, 3751.00, 4, 'C13'),
    ('789DEF123', 'Hitachi HCP-G10', 'HCP-G10', 'Object Storage', 950, 3239.50, 12, 'C13, C19'),
    ('456xyz789', 'Cisco NEXUS-93108', 'NEXUS-93108', 'Switch', 500, 1705.00, 2, 'C13'),
    ('234KLM567', 'Brocade 6520', '6520', 'SAN Switch', 400, 1364.00, 2, 'C13');
INSERT INTO
    device_ethernet_fiber (
        serial_number,
        device_make_model,
        model,
        device_type,
        device_physical_port,
        device_port_type,
        device_port_mac_address_wwn,
        connected_device_port
    )
VALUES
    ('123ABC456', 'Oracle T8-4', 'T8-4', 'Server', 'PCI1/Port2', 'Network/ETH', '00:09:0F:AA:00:01', 'Nexus 92108/ ETH1/1'),
    ('123ABC456', 'Oracle T8-4', 'T8-4', 'Server', 'PCI2/Port1', 'SAN/Fiber', '50:02:77:a4:10:0c:4e:21', 'SAN-1 6520 / Port5'),
    ('789DEF123', 'Hitachi HCP-G10', 'HCP-G10', 'Object Storage', 'Node1/PCI1/Port1', 'Network/Fiber', '00:06:0F:AB:00:02', 'Nexus 92108/ ETH2/5'),
    ('789DEF123', 'Hitachi HCP-G10', 'HCP-G10', 'Object Storage', 'Node1/PCI1/Port2', 'Network/Fiber', '00:09:0F:AA:00:05', 'Nexus 92108/ ETH2/7'),
    ('456xyz789', 'Cisco NEXUS-93108', 'NEXUS-93108', 'Switch', 'ETH1', 'Network/ETH', '00:06:0B:AA:00:03', 'Nexus 92108/ ETH1/10'),
    ('456xyz789', 'Cisco NEXUS-93108', 'NEXUS-93108', 'Switch', 'ETH2', 'Network/Fiber', '00:09:0F:AA:00:07', 'Nexus 92108/ ETH1/15'),
    ('234KLM567', 'Brocade 6520', '6520', 'SAN Switch', 'Port1', 'SAN/Fiber', '50:02:77:a4:10:0c:4a:22', 'SAN-2 6520 / Port10'),
    ('234KLM567', 'Brocade 6520', '6520', 'SAN Switch', 'Port2', 'SAN/Fiber', '50:02:77:a4:10:0c:4b:20', 'SAN-1 6520 / Port18');

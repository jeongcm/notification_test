# Openstack Notification Message

---

#### 목차

- [notification event type](#notification-event-type)
	- [keystone](#keystone)
	- [nova](#nova)
	- [cinder](#cinder)
	- [Neutron](#neutron)

---

## keystone

- [keystone event type](#keystone-event-type)
	- [identity.project.created](#identity-project-created)
	- [identity.project.updated](#identity-project-updated)
	- [identity.project.deleted](#identity-project-deleted)

#### identity-project-created

    "payload": {
        "typeURI": "http://schemas.dmtf.org/cloud/audit/1.0/event",
        "initiator": {
            "username": "admin",
            "typeURI": "service/security/account/user",
            "user_id": "783c77cb3daa4f7bbdeb798514db9399",
            "host": {"agent": "python-keystoneclient",
            "address": "172.16.194.168"},
            "request_id": "req-460ba344-01a4-49d3-ab6d-af30f72e0f6b",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "783c77cb3daa4f7bbdeb798514db9399"
        },
        "target": {
            "typeURI": "data/security/project",
            "id": "7fef504ab5e045db9f8bf8441c311f5d"
        },
        "observer": {
            "typeURI": "service/security",
            "id": "a7cdb291f2054d3f8d8062b99448acfe"
        },
        "eventType": "activity",
        "eventTime": "2021-08-10T05:38:55.895747+0000",
        "action": "created.project",
        "outcome": "success",
        "id": "d2cac1f4-3b15-5af1-bf43-9219928281d8",
        "resource_info": "7fef504ab5e045db9f8bf8441c311f5d" <- tenant_id
    }

#### identity-project-updated

    "payload": {
        "typeURI": "http://schemas.dmtf.org/cloud/audit/1.0/event",
        "initiator": {
            "username": "admin",
            "typeURI": "service/security/account/user",
            "user_id": "783c77cb3daa4f7bbdeb798514db9399",
            "host": {
                "agent": "python-keystoneclient",
                "address": "172.16.194.168"
            }, 
            "request_id": "req-b35bea78-3874-49a3-8024-307af6406233",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "783c77cb3daa4f7bbdeb798514db9399"
        },
        "target": {
            "typeURI": "data/security/project",
            "id": "7fef504ab5e045db9f8bf8441c311f5d"
        },
        "observer": {
            "typeURI": "service/security",
            "id": "a7cdb291f2054d3f8d8062b99448acfe"
        },
        "eventType": "activity",
        "eventTime": "2021-08-10T06:33:05.212203+0000",
        "action": "updated.project",
        "outcome": "success",
        "id": "bc124e55-a298-5d63-94e4-feaf0acaf3cb",
        "resource_info": "7fef504ab5e045db9f8bf8441c311f5d" <- tenant_id
    }

#### identity-project-deleted

    "payload": {
        "typeURI": "http://schemas.dmtf.org/cloud/audit/1.0/event",
        "initiator": {
            "username": "admin",
            "typeURI": "service/security/account/user",
            "user_id": "783c77cb3daa4f7bbdeb798514db9399",
            "host": {
                "agent": "python-keystoneclient",
                "address": "172.16.194.168"
            }, 
            "request_id": "req-5a2d61f7-00e7-41c7-8879-09ae2267b156",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "783c77cb3daa4f7bbdeb798514db9399"
        }, 
        "target": {
            "typeURI": "data/security/project",
            "id": "7fef504ab5e045db9f8bf8441c311f5d"
        },
        "observer": {
            "typeURI": "service/security",
            "id": "a7cdb291f2054d3f8d8062b99448acfe"
        },
        "eventType": "activity",
        "eventTime": "2021-08-10T06:40:46.357913+0000",
        "action": "deleted.project",
        "outcome": "success",
        "id": "88dd8941-e8c3-5a45-b5a7-9275deef5b9b",
        "resource_info": "7fef504ab5e045db9f8bf8441c311f5d" <- tenant_id
    }

## nova

- [nova event type](#nova-event-type)
    - [compute.instance.create.end](#nova-instance-created)
	- [compute.instance.update](#nova-instance-updated)
	- [compute.instance.delete.end](#nova-instance-deleted)
	- [volume.attach.end](#nova-instance-volume-attached)
	- [compute.instance.suspend.end](#nova-instance-suspend)
	- [snapshot.create.end](#nova-instance-snapshot-created)
	
### nova-instance-created 
(compute.instance.create.start를 보고 작성, instance_id만 확인하면 되지만 .end 를 한번 봐야할 필요는 있음)
	
	"payload": {
		"state_description": "",
		"availability_zone": "nova",
		"terminated_at": "",
		"ephemeral_gb": 0,
		"instance_type_id": 1,
		"message": "\\uc644\\ub8cc",
		"deleted_at": "",
		"fixed_ips": [
		{
			"version": 4,
			"vif_mac": "fa:16:3e:46:f1:a3",
			"floating_ips": [],
			"label": "testNetwork",
			"meta": {},
			"address": "192.168.122.71",
			"type": "fixed"
		}
		],
		"instance_id": "ea6048c1-3f0f-4e55-9f0f-70d09877d66f",
		"display_name": "in1",
		"reservation_id": "r-qy6y10dm",
		"hostname": "in1",
		"state": "active",
		"progress": "",
		"launched_at": "2021-08-24T08:29:33.046697",
		"metadata": {},
		"node": "jcm",
		"ramdisk_id": "",
		"access_ip_v6": null,
		"disk_gb": 1,
		"access_ip_v4": null,
		"kernel_id": "",
		"host": "jcm",
		"user_id": "007f371aefc6466da99f21ee6db42c5c",
		"image_ref_url": "http://172.16.194.168:9292/images/",
		"cell_name": "",
		"root_gb": 1,
		"tenant_id": "e6bff8e243dc42c38615572bd9e63d00",
		"created_at": "2021-08-24 08:29:18+00:00",
		"memory_mb": 512,
		"instance_type": "m1.tiny",
		"vcpus": 1,
		"image_meta": {
			"min_disk": "1",
			"container_format": "bare",
			"min_ram": "0",
			"disk_format": "iso",
			"base_image_ref": ""
		},
		"architecture": null,
		"os_type": null,
		"instance_flavor_id": "1"
	}

### nova-instance-updated
	
	"payload": {
		"state_description": "",
		"availability_zone": null,
		"terminated_at": "",
		"ephemeral_gb": 0,
		"instance_type_id": 1,
		"bandwidth": {},
		"deleted_at": "",
		"reservation_id": "r-ytxq0c5i",
		"instance_id": "559ee2ad-b3a1-49d5-b6e8-c68d4daeb081",
		"display_name": "testins",
		"hostname": "testins",
		"state": "error",
		"old_state": "building",
		"progress": "",
		"launched_at": "",
		"metadata": {},
		"node": null,
		"ramdisk_id": "",
		"access_ip_v6": null,
		"disk_gb": 1,
		"access_ip_v4": null,
		"kernel_id": "",
		"host": null,
		"user_id": "783c77cb3daa4f7bbdeb798514db9399",
		"image_ref_url": "http://172.16.194.168:9292/images/",
		"cell_name": "",
		"audit_period_beginning": "2021-08-01T00:00:00.000000",
		"root_gb": 1,
		"tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
		"created_at": "2021-08-11 06:50:54+00:00",
		"old_task_state": "block_device_mapping",
		"memory_mb": 512,
		"instance_type": "m1.tiny",
		"vcpus": 1,
		"image_meta": {
			"min_disk": "1",
			"container_format": "bare",
			"min_ram": "0",
			"disk_format": "iso",
			"base_image_ref": ""
		},
		"architecture": null,
		"new_task_state": null,
		"audit_period_ending": "2021-08-11T06:51:03.002867",
		"os_type": null,
		"instance_flavor_id": "1"
	}

### nova-instance-deleted

	"payload": {
		"state_description": "",
		"availability_zone": null,
		"terminated_at": "2021-08-11T06:57:35.000000",
		"ephemeral_gb": 0,
		"instance_type_id": 1,
		"deleted_at": "2021-08-11T06:57:35.312636",
		"reservation_id": "r-ytxq0c5i",
		"instance_id": "559ee2ad-b3a1-49d5-b6e8-c68d4daeb081",
		"display_name": "testins",
		"hostname": "testins",
		"state": "deleted",
		"progress": "",
		"launched_at": "",
		"metadata": {},
		"node": null,
		"ramdisk_id": "",
		"access_ip_v6": null,
		"disk_gb": 1,
		"access_ip_v4": null,
		"kernel_id": "",
		"host": null,
		"user_id": "783c77cb3daa4f7bbdeb798514db9399",
		"image_ref_url": "http://172.16.194.168:9292/images/",
		"cell_name": "",
		"root_gb": 1,
		"tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
		"created_at": "2021-08-11 06:50:54+00:00",
		"memory_mb": 512,
		"instance_type": "m1.tiny",
		"vcpus": 1,
		"image_meta": {
			"min_disk": "1",
			"container_format": "bare",
			"min_ram": "0",
			"disk_format": "iso",
			"base_image_ref": ""
		},
		"architecture": null,
		"os_type": null,
		"instance_flavor_id": "1"
	}

### nova-instance-volume-attached
	"payload": {
		"status": "in-use",
		"display_name": "testVolume",
		"volume_attachment": [
			{
				"instance_uuid": "c18fc73e-ac32-4bd8-8737-3c81192bed86",
				"detach_time": null,
				"attach_time": "2021-08-13T05:47:29.000000",
				"deleted": false,
				"attach_mode": "rw",
				"created_at": "2021-08-13T05:46:10.000000",
				"attached_host": "jcm.localdomain",
				"updated_at": "2021-08-13T05:47:31.000000",
				"attach_status": "attached",
				"volume": {
					"migration_status": null,
					"provider_id": null,
					"availability_zone": "nova",
					"terminated_at": null,
					"updated_at": "2021-08-13T05:47:31.000000",
					"provider_geometry": null,
					"replication_extended_status": null,
					"replication_status": null,
					"snapshot_id": null,
					"ec2_id": null,
					"deleted_at": null,
					"id": "1e947156-ca02-4475-8328-75066498bc4e",
					"size": 1,
					"user_id": "783c77cb3daa4f7bbdeb798514db9399",
					"display_description": "",
					"cluster_name": null,
					"project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
					"launched_at": "2021-07-12T07:01:45.000000",
					"scheduled_at": "2021-07-12T07:01:45.000000",
					"status": "in-use",
					"volume_type_id": "98927d17-0d92-4b5b-bbf1-e2a6a6fdda54",
					"multiattach": false,
					"deleted": false,
					"service_uuid": "f702a95b-86c8-4691-9afc-44e801b9ea44",
					"provider_location": "172.16.194.168:3260,iqn.2010-10.org.openstack:volume-1e947156-ca02-4475-8328-75066498bc4e iqn.2010-10.org.openstack:volume-1e947156-ca02-4475-8328-75066498bc4e 0",
					"host": "jcm@lvm#lvm",
					"consistencygroup_id": null,
					"source_volid": null,
					"provider_auth": "CHAP kjRwYpBnpu9GT47Mfux8 Jo57JJGdFmiLqNtK",
					"previous_status": null,
					"display_name": "testVolume",
					"bootable": false,
					"created_at": "2021-07-12T07:01:45.000000",
					"attach_status": "attached",
					"_name_id": null,
					"encryption_key_id": null,
					"replication_driver_data": null,
					"group_id": null,
					"shared_targets": false
				},
				"connection_info": {
					"access_mode": "rw",
					"attachment_id": "b427c3ef-6588-44b2-9559-0db7c873e412",
					"target_discovered": false,
					"encrypted": false,
					"driver_volume_type": "iscsi",
					"qos_specs": null,
					"target_iqn": "iqn.2010-10.org.openstack:volume-1e947156-ca02-4475-8328-75066498bc4e",
					"target_portal": "172.16.194.168:3260",
					"volume_id": "1e947156-ca02-4475-8328-75066498bc4e",
					"target_lun": 0,
					"auth_password": "Jo57JJGdFmiLqNtK",
					"auth_username": "kjRwYpBnpu9GT47Mfux8",
					"auth_method": "CHAP"
				},
				"volume_id": "1e947156-ca02-4475-8328-75066498bc4e",
				"mountpoint": "/dev/vdb",
				"deleted_at": null,
				"id": "b427c3ef-6588-44b2-9559-0db7c873e412",
				"connector": {
					"initiator": "iqn.1994-05.com.redhat:c377366444a2",
					"ip": "172.16.194.168",
					"system uuid": "07F64D56-331D-AE14-3003-964DF8FB8EC1",
					"platform": "x86_64",
					"host": "jcm.localdomain",
					"do_local_attach": false,
					"mountpoint": "/dev/vdb",
					"os_type": "linux2",
					"multipath": false
				}
			}
		],
		"availability_zone": "nova",
		"tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
		"created_at": "2021-07-12T07:01:45+00:00",
		"volume_id": "1e947156-ca02-4475-8328-75066498bc4e",
		"volume_type": "98927d17-0d92-4b5b-bbf1-e2a6a6fdda54",
		"host": "jcm@lvm#lvm",
		"replication_driver_data": null,
		"replication_status": null,
		"snapshot_id": null,
		"replication_extended_status": null,
		"user_id": "783c77cb3daa4f7bbdeb798514db9399",
		"metadata": [],
		"launched_at": "2021-07-12T07:01:45+00:00",
		"size": 1
	}

### nova-instance-suspend

	"payload": {
		"state_description": "",
		"availability_zone": "nova",
		"terminated_at": "",
		"ephemeral_gb": 0,
		"instance_type_id": 1,
		"deleted_at": "",
		"reservation_id": "r-y0ymva00",
		"instance_id": "c18fc73e-ac32-4bd8-8737-3c81192bed86",
		"display_name": "in1",
		"hostname": "in1",
		"state": "suspended",
		"progress": "",
		"launched_at": "2021-04-22T08:08:14.000000",
		"metadata": {},
		"node": "jcm",
		"ramdisk_id": "",
		"access_ip_v6": null,
		"disk_gb": 1,
		"access_ip_v4": null,
		"kernel_id": "",
		"host": "jcm.localdomain",
		"user_id": "783c77cb3daa4f7bbdeb798514db9399",
		"image_ref_url": "http://172.16.194.168:9292/images/",
		"cell_name": "",
		"root_gb": 1,
		"tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
		"created_at": "2021-04-22 08:07:58+00:00",
		"memory_mb": 512,
		"instance_type": "m1.tiny",
		"vcpus": 1,
		"image_meta": {
			"min_disk": "1",
			"container_format": "bare",
			"min_ram": "0",
			"disk_format": "iso",
			"base_image_ref": ""
		},
		"architecture": null,
		"os_type": null,
		"instance_flavor_id": "1"
	}

### nova-instance-snapshot-created (cinder-snapshot-created 와 일치)
	
	"payload": {
		"status": "available",
		"display_name": "snapshot for test_ins_snap",
		"availability_zone": "nova",
		"deleted": "",
		"tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
		"created_at": "2021-08-13T06:16:12+00:00",
		"snapshot_id": "2f89dbda-f70d-40c4-9337-4805aea9fb77",
		"volume_size": 1,
		"volume_id": "1e947156-ca02-4475-8328-75066498bc4e",
		"user_id": "783c77cb3daa4f7bbdeb798514db9399",
		"metadata": ""
	}

## cinder

- [cinder event type](#cinder-event-type)
    - [volume.create.end](#cinder-volume-created)
    - [volume.update.end](#cinder-volume-updated)
    - [volume.delete.end](#cinder-volume-deleted)
    - [snapshot.create.end](#cinder-snapshot-created)
    - [snapshot.update.end](#cinder-snapshot-updated)
    - [snapshot.delete.end](#cinder-snapshot-deleted)
    - [volume_type.create](#cinder-volume_type-created)
    - [volume_type.update](#cinder-volume_type-updated)
    - [volume_type.delete](#cinder-volume_type-deleted)
    - [volume_type_project.access.add](#cinder-volume_type-access-add)
    - [volume_type_extra_specs.create](#cinder-volume_type_extra_specs-created)
    - [volume_type_extra_specs.delete](#cinder-volume_type_extra_specs-deleted)

### cinder-volume-created

    "payload": {
        "status": "available",
        "display_name": "ttttttttttttt",
        "volume_attachment": [],
        "availability_zone": "nova",
        "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
        "created_at": "2021-08-10T05:38:14+00:00",
        "volume_id": "c2af83e6-d386-4672-8938-1df26677049f",
        "volume_type": "98927d17-0d92-4b5b-bbf1-e2a6a6fdda54",
        "host": "jcm@lvm#lvm",
        "replication_driver_data": null,
        "replication_status": null,
        "snapshot_id": null,
        "replication_extended_status": null,
        "user_id": "783c77cb3daa4f7bbdeb798514db9399",
        "metadata": [],
        "launched_at": "2021-08-10T05:38:15.147695+00:00",
        "size": 1
    }

### cinder-volume-updated
    "payload": {
        "status": "available",
        "display_name": "tttttttttttttdfdf",
        "volume_attachment": [],
        "availability_zone": "nova",
        "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
        "created_at": "2021-08-10T05:38:14+00:00",
        "volume_id": "c2af83e6-d386-4672-8938-1df26677049f",
        "volume_type": "98927d17-0d92-4b5b-bbf1-e2a6a6fdda54",
        "host": "jcm@lvm#lvm",
        "replication_driver_data": null,
        "replication_status": null,
        "snapshot_id": null,
        "replication_extended_status": null,
        "user_id": "783c77cb3daa4f7bbdeb798514db9399",
        "metadata": [],
        "launched_at": "2021-08-10T05:38:15+00:00",
        "size": 1
    }

### cinder-volume-deleted

    "payload": {
        "status": "deleted",
        "display_name": "tttttttttttttdfdf",
        "volume_attachment": [],
        "availability_zone": "nova",
        "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
        "created_at": "2021-08-10T05:38:14+00:00",
        "volume_id": "c2af83e6-d386-4672-8938-1df26677049f",
        "volume_type": "98927d17-0d92-4b5b-bbf1-e2a6a6fdda54",
        "host": "jcm@lvm#lvm",
        "replication_driver_data": null,
        "replication_status": null,
        "snapshot_id": null,
        "replication_extended_status": null,
        "user_id": "783c77cb3daa4f7bbdeb798514db9399",
        "metadata": [],
        "launched_at": "2021-08-10T05:38:15+00:00",
        "size": 1
    }

### cinder-snapshot-created

    "payload": {
        "status": "available",
        "display_name": "testsnap",
        "availability_zone": "nova",
        "deleted": "",
        "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
        "created_at": "2021-08-11T03:13:07+00:00",
        "snapshot_id": "b31fc802-1c94-43fc-8745-def92b5b6154",
        "volume_size": 1,
        "volume_id": "1e947156-ca02-4475-8328-75066498bc4e",
        "user_id": "783c77cb3daa4f7bbdeb798514db9399",
        "metadata": ""
    }

### cinder-snapshot-updated

    "payload": {
        "status": "available",
        "display_name": "testSnapshot",
        "availability_zone": "nova",
        "deleted": "",
        "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
        "created_at": "2021-08-11T03:13:07+00:00",
        "snapshot_id": "b31fc802-1c94-43fc-8745-def92b5b6154",
        "volume_size": 1,
        "volume_id": "1e947156-ca02-4475-8328-75066498bc4e",
        "user_id": "783c77cb3daa4f7bbdeb798514db9399",
        "metadata": ""
    }

### cinder-snapshot-deleted

    "payload": {
        "status": "deleted",
        "display_name": "testSnapshot",
        "availability_zone": "nova",
        "deleted": "True",
        "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
        "created_at": "2021-08-11T03:13:07+00:00",
        "snapshot_id": "b31fc802-1c94-43fc-8745-def92b5b6154",
        "volume_size": 1,
        "volume_id": "1e947156-ca02-4475-8328-75066498bc4e",
        "user_id": "783c77cb3daa4f7bbdeb798514db9399",
        "metadata": ""
    }

### cinder-volume_type-created

    "payload": {
        "volume_types": {
            "name": "testVolumeType",
            "qos_specs_id": null,
            "deleted": false,
            "created_at": "2021-08-11T03:29:38.000000",
            "updated_at": null,
            "extra_specs": {},
            "is_public": true,
            "deleted_at": null,
            "id": "5e22242f-4fd6-4004-a93c-d6ac31d24b16",
            "description": ""
        }
    }

### cinder-volume_type-updated

    "payload": {
        "volume_types": {
            "name": "iscsi2",
            "qos_specs_id": null,
            "deleted": false,
            "created_at": "2021-04-22T07:34:42.000000",
            "updated_at": "2021-08-11T03:26:01.000000",
            "extra_specs": {
                "volume_backend_name": "lvm"
            },
            "is_public": true,
            "deleted_at": null,
            "id": "98927d17-0d92-4b5b-bbf1-e2a6a6fdda54",
            "description": ""
        }
    }

### cinder-volume_type-deleted (직접 찍어보진 않음)

    "payload": {
        "volume_types": {
            "name": "iscsi2",
            "qos_specs_id": null,
            "deleted": false,
            "created_at": "2021-04-22T07:34:42.000000",
            "updated_at": "2021-08-11T03:26:01.000000",
            "extra_specs": {
                "volume_backend_name": "lvm"
            },
            "is_public": true,
            "deleted_at": null,
            "id": "98927d17-0d92-4b5b-bbf1-e2a6a6fdda54",
            "description": ""
        }
    }

### cinder-volume_type-access-add

    "payload": {
        "volume_type_id": "5e22242f-4fd6-4004-a93c-d6ac31d24b16",
        "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
    }

### cinder-volume_type_extra_specs-created

    "payload": {
        "created_at": "2021-08-11T03:29:38.000000",
        "type_id": "5e22242f-4fd6-4004-a93c-d6ac31d24b16",
        "specs": {"111": "111"},
        "updated_at": "2021-08-11T03:31:38.000000"
    }

### cinder-volume_type_extra_specs-deleted

    "payload": {
        "type_id": "5e22242f-4fd6-4004-a93c-d6ac31d24b16",
        "created_at": "2021-08-11T03:29:38.000000",
        "deleted_at": null,
        "id": "111",
        "updated_at": "2021-08-11T03:31:38.000000"
    }


## neutron

- [neutron event type](#neutron-event-type)
	- [network.create.end](#neutron-network-created)
	- [network.update.end](#neutron-network-updated)
	- [network.delete.end](#neutron-network-deleted)
	- [subnet.create.end](#neutron-subnet-created)
	- [subnet.update.end](#neutron-subnet-updated)
	- [security_group.create.end](#neutron-security_group-created)
	- [security_group.update.end](#neutron-security_group-updated)
	- [security_group.delete.end](#neutron-security_group-deleted)
    - [security_group_rule.create.end](#neutron-security_group_rule-created)
    - [security_group_rule.update.end](#neutron-security_group_rule-updated)
    - [security_group_rule.delete.end](#neutron-security_group_rule-deleted)
    - [router.create.end](#neutron-router-created)
    - [router.update.end](#neutron-router-updated)
    - [router.delete.end](#neutron-router-deleted)
    - [router.interface.create](#neutron-router-interface-created)
    - [floatingip.create.end](#neutron-floating_ip-created)
    - [floatingip.update.end](#neutron-floating_ip-updated)
    - [floatingip.delete.end](#neutron-floating_ip-deleted)
    
### neutron-network-created

    "payload": {
        "network": {
            "provider:physical_network": null,
            "ipv6_address_scope": null,
            "revision_number": 1,
            "port_security_enabled": true,
            "mtu": 1442,
            "id": "32c1bc8c-af55-4a40-9f22-e849baa65fd3",
            "router:external": false,
            "availability_zone_hints": [],
            "availability_zones": [],
            "provider:segmentation_id": 11,
            "ipv4_address_scope": null,
            "shared": false,
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "status": "ACTIVE",
            "subnets": [],
            "description": "",
            "tags": [],
            "updated_at": "2021-08-10T06:57:21Z",
            "is_default": false,
            "qos_policy_id": null,
            "name": "testnetworkdfdkfjdf",
            "admin_state_up": true,
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T06:57:21Z",
            "provider:network_type": "geneve"
        }
    }

### neutron-network-updated

    "payload": {
        "network": {
            "provider:physical_network": null,
            "ipv6_address_scope": null,
            "revision_number": 4,
            "port_security_enabled": true,
            "mtu": 1442,
            "id": "32c1bc8c-af55-4a40-9f22-e849baa65fd3",
            "router:external": false,
            "availability_zone_hints": [],
            "availability_zones": [],
            "provider:segmentation_id": 11,
            "ipv4_address_scope": null,
            "shared": false,
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "status": "ACTIVE",
            "subnets": ["529e929f-8997-4722-a045-3a96a7356a05"],
            "description": "",
            "tags": [],
            "updated_at": "2021-08-10T07:10:58Z",
            "qos_policy_id": null,
            "name": "testnetworkdfdkfjdfdfdfdf",
            "admin_state_up": true,
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T06:57:21Z",
            "provider:network_type": "geneve"
        }
    }

### neutron-network-deleted

    "payload": {
        "network_id":"32c1bc8c-af55-4a40-9f22-e849baa65fd3",
        "network":{
            "provider:physical_network":null,
            "ipv6_address_scope":null,
            "revision_number":4,
            "port_security_enabled":true,
            "provider:network_type":"geneve",
            "id":"32c1bc8c-af55-4a40-9f22-e849baa65fd3",
            "router:external":false,
            "availability_zone_hints":[],
            "availability_zones":[],
            "provider:segmentation_id":11,
            "ipv4_address_scope":null,
            "shared":false,
            "project_id":"f123f8eec9f2497d8a0453b6ab1d853b",
            "status":"ACTIVE",
            "subnets":["529e929f-8997-4722-a045-3a96a7356a05"],
            "description":"",
            "tags":[],
            "updated_at":"2021-08-10T07:10:58Z",
            "qos_policy_id":null,
            "name":"testnetworkdfdkfjdfdfdfdf",
            "admin_state_up":true,
            "tenant_id":"f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at":"2021-08-10T06:57:21Z",
            "mtu":1442,
            "vlan_transparent":null
        }
    }

### neutron-subnet-created

    "payload": {
        "subnet": {
            "service_types": [],
            "description": "",
            "enable_dhcp": true,
            "tags": [],
            "network_id": "32c1bc8c-af55-4a40-9f22-e849baa65fd3",
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T06:57:22Z",
            "dns_nameservers": [],
            "updated_at": "2021-08-10T06:57:22Z",
            "ipv6_ra_mode": null,
            "allocation_pools": [
                {
                    "start": "192.168.121.2",
                    "end": "192.168.121.254"
                }
            ],
            "gateway_ip": "192.168.121.1",
            "revision_number": 0,
            "ipv6_address_mode": null,
            "ip_version": 4,
            "host_routes": [],
            "cidr": "192.168.121.0/24",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "529e929f-8997-4722-a045-3a96a7356a05",
            "subnetpool_id": null,
            "name": "dfdf"
        }
    }

### neutron-subnet-updated

    "payload": {
        "subnet": {
            "service_types": [],
            "description": "",
            "enable_dhcp": true,
            "tags": [],
            "network_id": "32c1bc8c-af55-4a40-9f22-e849baa65fd3",
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T06:57:22Z",
            "dns_nameservers": [],
            "updated_at": "2021-08-10T07:04:44Z",
            "ipv6_ra_mode": null,
            "allocation_pools": [
                {
                    "start": "192.168.121.2",
                    "end": "192.168.121.254"
                }
            ],
            "gateway_ip": "192.168.121.1",
            "revision_number": 1,
            "ipv6_address_mode": null,
            "ip_version": 4,
            "host_routes": [],
            "cidr": "192.168.121.0/24",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "529e929f-8997-4722-a045-3a96a7356a05",
            "subnetpool_id": null,
            "name": "dfdfdfdf"
        }
    }

### neutron-security_group-created

    "payload": {
        "security_group": {
            "description": "",
            "tags": [],
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T07:16:38Z",
            "updated_at": "2021-08-10T07:16:38Z",
            "security_group_rules": [
                {
                    "direction": "egress",
                    "protocol": null,
                    "description": null,
                    "tags": [],
                    "port_range_max": null,
                    "updated_at": "2021-08-10T07:16:38Z",
                    "revision_number": 0,
                    "id": "749de20a-24a4-40fc-a137-d34c65b61cc0",
                    "remote_group_id": null,
                    "remote_ip_prefix": null,
                    "created_at": "2021-08-10T07:16:38Z",
                    "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
                    "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
                    "port_range_min": null,
                    "ethertype": "IPv6",
                    "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
                },
                {
                    "direction": "egress",
                    "protocol": null,
                    "description": null,
                    "tags": [],
                    "port_range_max": null,
                    "updated_at": "2021-08-10T07:16:38Z",
                    "revision_number": 0,
                    "id": "913e1bf3-8031-4235-bf78-0d848b02de24",
                    "remote_group_id": null,
                    "remote_ip_prefix": null,
                    "created_at": "2021-08-10T07:16:38Z",
                    "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
                    "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
                    "port_range_min": null,
                    "ethertype": "IPv4",
                    "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
                }
            ],
            "revision_number": 1,
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
            "name": "testSecurity"
        }
    }

### neutron-security_group-updated
    
    "payload": {
        "security_group": {
            "description": "",
            "tags": [],
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T07:16:38Z",
            "updated_at": "2021-08-10T08:56:24Z",
            "security_group_rules": [
            {
                "direction": "egress",
                "protocol": null,
                "description": null,
                "tags": [],
                "port_range_max": null,
                "updated_at": "2021-08-10T07:16:38Z",
                "revision_number": 0,
                "id": "749de20a-24a4-40fc-a137-d34c65b61cc0",
                "remote_group_id": null,
                "remote_ip_prefix": null,
                "created_at": "2021-08-10T07:16:38Z",
                "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
                "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
                "port_range_min": null,
                "ethertype": "IPv6",
                "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
            },
            {
                "direction": "egress",
                "protocol": null,
                "description": null,
                "tags": [],
                "port_range_max": null,
                "updated_at": "2021-08-10T07:16:38Z",
                "revision_number": 0,
                "id": "913e1bf3-8031-4235-bf78-0d848b02de24",
                "remote_group_id": null,
                "remote_ip_prefix": null,
                "created_at": "2021-08-10T07:16:38Z",
                "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
                "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
                "port_range_min": null,
                "ethertype": "IPv4",
                "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
            }
            ],
            "revision_number": 4,
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
            "name": "testSecurity2"
        }
    }

### neutron-security_group-deleted

    "payload": {
        "security_group": {
            "description": "",
            "tags": [],
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T07:16:38Z",
            "updated_at": "2021-08-10T08:56:24Z",
            "security_group_rules": [
            {
                "direction": "egress",
                "protocol": null,
                "description": null,
                "tags": [],
                "port_range_max": null,
                "updated_at": "2021-08-10T07:16:38Z",
                "revision_number": 0,
                "id": "749de20a-24a4-40fc-a137-d34c65b61cc0",
                "remote_group_id": null,
                "remote_ip_prefix": null,
                "created_at": "2021-08-10T07:16:38Z",
                "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
                "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
                "port_range_min": null,
                "ethertype": "IPv6",
                "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
            },
            {
                "direction": "egress",
                "protocol": null,
                "description": null,
                "tags": [],
                "port_range_max": null,
                "updated_at": "2021-08-10T07:16:38Z",
                "revision_number": 0,
                "id": "913e1bf3-8031-4235-bf78-0d848b02de24",
                "remote_group_id": null,
                "remote_ip_prefix": null,
                "created_at": "2021-08-10T07:16:38Z",
                "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
                "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
                "port_range_min": null,
                "ethertype": "IPv4",
                "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
            }
            ],
            "revision_number": 4,
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
            "name": "testSecurity2"
        },
        "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7"
    }

### neutron-security_group_rule-created
    "payload": {
        "security_group_rule": {
            "remote_group_id": null,
            "direction": "ingress",
            "protocol": "tcp",
            "description": "asdf",
            "ethertype": "IPv4",
            "remote_ip_prefix": "0.0.0.0/0",
            "port_range_max": 44444,
            "updated_at": "2021-08-10T07:26:09Z",
            "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
            "port_range_min": 44444,
            "revision_number": 0,
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T07:26:09Z",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "f71e9bb9-d04d-4066-bc60-5cd9d0722ab4"
        }
    }

### neutron-security_group_rule-deleted

    "payload": {
        "security_group_rule": {
            "remote_group_id": null,
            "direction": "ingress",
            "protocol": "tcp",
            "description": "asdf",
            "ethertype": "IPv4",
            "remote_ip_prefix": "0.0.0.0/0",
            "port_range_max": 44444,
            "updated_at": "2021-08-10T07:26:09Z",
            "security_group_id": "aadbedc9-102f-4d17-b613-7ac501371ac7",
            "port_range_min": 44444,
            "revision_number": 0,
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T07:26:09Z",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "f71e9bb9-d04d-4066-bc60-5cd9d0722ab4"
        }
    }

### neutron-router-created
    "payload": {
        "router": {
            "status": "ACTIVE",
            "external_gateway_info": {
                "network_id": "6d2328f9-7967-4107-8612-a81c489c1822",
                "enable_snat": true,
                "external_fixed_ips": [
                    {
                        "subnet_id": "96e3b4e6-9cf7-4aa1-be04-62e0807796f8",
                        "ip_address": "192.168.15.109"
                    }
                ]
            },
            "availability_zone_hints": [],
            "availability_zones": [],
            "description": "",
            "tags": [],
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-11T02:35:27Z",
            "admin_state_up": true,
            "updated_at": "2021-08-11T02:35:29Z",
            "revision_number": 3,
            "routes": [],
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "38c5d4e1-d354-4572-a6f3-1b8661b44fc8",
            "name": "testr"
        }
    }
### neutron-router-updated

    "payload":  {
        "router": {
            "status": "ACTIVE",
            "external_gateway_info": {
                "network_id": "6d2328f9-7967-4107-8612-a81c489c1822",
                "enable_snat": true,
                "external_fixed_ips": [
                    {
                        "subnet_id": "96e3b4e6-9cf7-4aa1-be04-62e0807796f8",
                        "ip_address": "192.168.15.109"
                    }
                ]
            },
            "availability_zone_hints": [],
            "availability_zones": [],
            "description": "",
            "tags": [],
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-11T02:35:27Z",
            "admin_state_up": true,
            "updated_at": "2021-08-11T02:38:11Z",
            "revision_number": 4,
            "routes": [],
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "38c5d4e1-d354-4572-a6f3-1b8661b44fc8",
            "name": "testrdff"
        }
    }

### neutron-router-interface-created
	"payload": {
		"router_interface": {
		"network_id": "fd65e9f1-bbd2-4e25-a41b-09e031f85c2a",
		"tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
		"subnet_id": "fd0b1fcd-a117-456d-9d25-2fb5c66f3d8e",
		"subnet_ids": [
			"fd0b1fcd-a117-456d-9d25-2fb5c66f3d8e"
		],
		"port_id": "ce24beea-3e0a-49ef-be52-e1bb1bf7fff2",
		"id": "d33fda10-2cd2-4a32-a7db-ac00d560c208"
		}
	}

### neutron-router-deleted

    "payload": {
        "router_id": "38c5d4e1-d354-4572-a6f3-1b8661b44fc8",
        "router": {
            "status": "ACTIVE",
            "external_gateway_info": {
                "network_id": "6d2328f9-7967-4107-8612-a81c489c1822",
                "enable_snat": true,
                "external_fixed_ips": [
                    {
                        "subnet_id": "96e3b4e6-9cf7-4aa1-be04-62e0807796f8",
                        "ip_address": "192.168.15.109"
                    }
                ]
            },
            "availability_zone_hints": [],
            "availability_zones": [],
            "description": "",
            "tags": [],
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-11T02:35:27Z",
            "admin_state_up": true,
            "updated_at": "2021-08-11T02:38:11Z",
            "revision_number": 4,
            "routes": [],
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "id": "38c5d4e1-d354-4572-a6f3-1b8661b44fc8",
            "name": "testrdff"
        }
    }

### neutron-floating_ip-created

    "payload": {
        "floatingip": {
            "router_id": null,
            "status": "DOWN",
            "description": "",
            "tags": [],
            "updated_at": "2021-08-10T05:37:46Z",
            "dns_domain": "",
            "floating_network_id": "6d2328f9-7967-4107-8612-a81c489c1822",
            "fixed_ip_address": null,
            "floating_ip_address": "192.168.15.77",
            "revision_number": 0,
            "port_id": null,
            "id": "b9bc8f81-6722-48fb-b602-73745133111c",
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "created_at": "2021-08-10T05:37:46Z",
            "dns_name": "",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
        }
    }

### neutron-floating_ip-updated

    "payload": 

### neutron-floating_ip-deleted

    "payload": {
        "floatingip_id": "b9bc8f81-6722-48fb-b602-73745133111c",
        "floatingip": {
            "router_id": null,
            "status": "DOWN",
            "description": "",
            "tags": [],
            "updated_at": "2021-08-10T05:37:46Z",
            "dns_domain": "",
            "floating_network_id": "6d2328f9-7967-4107-8612-a81c489c1822",
            "fixed_ip_address": null,
            "floating_ip_address": "192.168.15.77",
            "revision_number": 0,
            "port_id": null,
            "id": "b9bc8f81-6722-48fb-b602-73745133111c",
            "dns_name": "",
            "created_at": "2021-08-10T05:37:46Z",
            "tenant_id": "f123f8eec9f2497d8a0453b6ab1d853b",
            "project_id": "f123f8eec9f2497d8a0453b6ab1d853b"
        }
    }


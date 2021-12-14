package sync

import (
	"log"
	"sync"
)

const (
	// ResourceTenant 테넌트
	ResourceTenant = "tenant"
	// ResourceInstance 인스턴스
	ResourceInstance = "instance"
	// ResourceVolumeList 볼륨 목록
	ResourceVolumeList = "volumeList"
	// ResourceVolume 볼륨
	ResourceVolume = "volume"
	// ResourceVolumeSnapshot 볼륨 스냅샷
	ResourceVolumeSnapshot = "volumeSnapshot"
	// ResourceVolumeType 볼륨 타입
	ResourceVolumeType = "volumeType"
	// ResourceNetwork 네트워크
	ResourceNetwork = "network"
	// ResourceSubnet 서브넷
	ResourceSubnet = "subnet"
	// ResourceSecurityGroup 보안 그룹
	ResourceSecurityGroup = "securityGroup"
	// ResourceSecurityGroupRule 보안 그룹 규칙
	ResourceSecurityGroupRule = "securityGroupRule"
	// ResourceRouter 라우터
	ResourceRouter = "router"
	// ResourceFloatingIP floating ip
	ResourceFloatingIP = "floatingIP"
)

var reservedSynchronizerMapLock = sync.Mutex{}
var reservedSynchronizerMap map[uint64]*ReservedSynchronizer

// ReservedSynchronizer 예약 동기화 구조체
type ReservedSynchronizer struct {
	projects           map[string]bool
	instances          map[string]bool
	volumes            map[string]bool
	volumeSnapshots    map[string]bool
	volumeTypes        map[string]bool
	networks           map[string]bool
	subnets            map[string]bool
	securityGroups     map[string]bool
	securityGroupRules map[string]bool
	routers            map[string]bool
	routerInterfaces   map[string]bool
	floatingIPs        map[string]bool
	volumeList         bool
}

// ReservedSync 는 예약된 클러스터 리소스의 동기화를 하는 함수이다.
func ReservedSync() {
	results := make(map[uint64]*ReservedSynchronizer)

	reservedSynchronizerMapLock.Lock()

	for k, v := range reservedSynchronizerMap {
		results[k] = v
	}

	reservedSynchronizerMap = nil

	reservedSynchronizerMapLock.Unlock()

	for clusterID, resources := range results {
		if err := resources.reservedSync(clusterID); err != nil {
			log.Printf("Could not sync cluster(%d). cause: %+v", clusterID, err)
		}
	}
}

func (rs *ReservedSynchronizer) reservedSync(clusterID uint64) error {

	log.Printf("cluster id : %d\n", clusterID)

	for uuid := range rs.projects {
		log.Printf("projects (%s)\n", uuid)
	}

	for uuid := range rs.securityGroups {
		log.Printf("securityGroups (%s)\n", uuid)
	}

	for uuid := range rs.securityGroupRules {
		log.Printf("securityGroupRules (%s)\n", uuid)
	}

	for uuid := range rs.networks {
		log.Printf("networks (%s)\n", uuid)
	}

	for uuid := range rs.subnets {
		log.Printf("subnets (%s)\n", uuid)
	}

	for uuid := range rs.routers {
		log.Printf("routers (%s)\n", uuid)
	}

	for uuid := range rs.volumeTypes {
		log.Printf("volumeTypes (%s)\n", uuid)
	}

	for uuid := range rs.volumes {
		log.Printf("volumes (%s)\n", uuid)
	}

	for uuid := range rs.volumeSnapshots {
		log.Printf("volumeSnapshots (%s)\n", uuid)
	}

	for uuid := range rs.instances {
		log.Printf("instance (%s)\n", uuid)
	}

	return nil
}

// Reserve 동기화 예약을 위한 함수
func Reserve(clusterID uint64, resourceType, uuid string) {
	reservedSynchronizerMapLock.Lock()
	defer reservedSynchronizerMapLock.Unlock()

	if reservedSynchronizerMap == nil {
		reservedSynchronizerMap = make(map[uint64]*ReservedSynchronizer)
	}

	if _, ok := reservedSynchronizerMap[clusterID]; !ok {
		reservedSynchronizerMap[clusterID] = &ReservedSynchronizer{
			projects:           make(map[string]bool),
			instances:          make(map[string]bool),
			volumes:            make(map[string]bool),
			volumeSnapshots:    make(map[string]bool),
			volumeTypes:        make(map[string]bool),
			networks:           make(map[string]bool),
			subnets:            make(map[string]bool),
			securityGroups:     make(map[string]bool),
			securityGroupRules: make(map[string]bool),
			routers:            make(map[string]bool),
			routerInterfaces:   make(map[string]bool),
			floatingIPs:        make(map[string]bool),
		}
	}

	reservedSynchronizerMap[clusterID].reserve(resourceType, uuid)
}

func (rs *ReservedSynchronizer) reserve(resourceType, uuid string) {
	switch resourceType {
	case ResourceTenant:
		rs.projects[uuid] = true

	case ResourceInstance:
		rs.instances[uuid] = true

	case ResourceVolume:
		rs.volumes[uuid] = true

	case ResourceVolumeSnapshot:
		rs.volumeSnapshots[uuid] = true

	case ResourceVolumeType:
		rs.volumeTypes[uuid] = true

	case ResourceNetwork:
		rs.networks[uuid] = true

	case ResourceSubnet:
		rs.subnets[uuid] = true

	case ResourceSecurityGroup:
		rs.securityGroups[uuid] = true

	case ResourceSecurityGroupRule:
		rs.securityGroupRules[uuid] = true

	case ResourceRouter:
		rs.routers[uuid] = true

	case ResourceFloatingIP:
		rs.floatingIPs[uuid] = true

	case ResourceVolumeList:
		rs.volumeList = true
	}
}

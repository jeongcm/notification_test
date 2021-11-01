package monitor

import (
	"errors"
)

type clusterMonitorCreationFunc func(string) Monitor

var clusterMonitorCreationFuncMap map[string]clusterMonitorCreationFunc

// RegisterClusterMonitorCreationFunc 는 클러스터 타입별 Monitor 구조체 생성 함수의 맵이다.
func RegisterClusterMonitorCreationFunc(typeCode string, fn clusterMonitorCreationFunc) {
	if clusterMonitorCreationFuncMap == nil {
		clusterMonitorCreationFuncMap = make(map[string]clusterMonitorCreationFunc)
	}

	clusterMonitorCreationFuncMap[typeCode] = fn
}

// Monitor monitor 인터페이스
type Monitor interface {
	Start() error
	Stop()
	Connect() error
}

// New 는 클러스터 타입별 모니터 인터페이스를 초기화하는 함수
func New(typeCode, serverURL string) (Monitor, error) {
	if fn, ok := clusterMonitorCreationFuncMap[typeCode]; ok {
		return fn(serverURL), nil
	}

	return nil, errors.New("not found cluster")
}

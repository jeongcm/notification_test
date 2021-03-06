package notification

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2/logger"
	"log"
	"net"
	"net/url"
	"test/monitor"
)

func init() {
	monitor.RegisterClusterMonitorCreationFunc("type.openstack", New)
}

type notificationHandler struct {
	cluster uint64
}

// Start openstack notification Start
func (n *notification) Start() error {
	if err := n.Connect(); err != nil {
		return err
	}

	if err := n.DeleteQueue(); err != nil {
		logger.Warnf("Could not delete openstack notification queue. cause: %v", err)
	}

	logger.Info("Success to connect cluster notification.")

	ns := notificationHandler{
		cluster: 1,
	}

	_, err := n.Subscribe(ns.handleEvent)
	if err != nil {
		logger.Errorf("Could not register to cluster notification. Cause: %v", err)
		_ = n.Disconnect()
		return err
	}

	logger.Info("Start cluster notification.\n")

	return nil
}

// Stop openstack notification stop
func (n *notification) Stop() {
	_ = n.Disconnect()
	logger.Infof("Close cluster notification.")
}

func (ns *notificationHandler) handleEvent(p Event) error {
	var e map[string]interface{}

	err := json.Unmarshal(p.Message().Body, &e)
	if err != nil {
		log.Printf("Could not subscribe event. cause: %v\n", err)
		return err
	}

	var m OsloMessage
	if err := json.Unmarshal([]byte(e["oslo.message"].(string)), &m); err != nil {
		log.Printf("Could not subscribe event. cause: %v\n", err)
		return err
	}

	log.Println(m.EventType)

	//TODO message type 별 동기화진행
	switch m.EventType {
	case "identity.project.created":
		fallthrough
	case "identity.project.updated":
		fallthrough
	case "identity.project.deleted":
		log.Printf("project notification %s\n", m.Payload["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "compute.instance.create.end":
		fallthrough
	case "compute.instance.update":
		fallthrough
	case "compute.instance.delete.end":
		fallthrough
	case "compute.instance.suspend.end":
		log.Printf("instance notification %s\n", m.Payload["instance_id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "volume.attach.end", "volume.detach.end":
		log.Printf("instance notification %s\n", m.Payload["instance_id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

		for _, attachment := range m.Payload["volume_attachment"].([]map[string]interface{}) {
			log.Printf("attachment instance uuid = %s\n", attachment["instance_uuid"].(string))
		}

	case "volume.create.end":
		fallthrough
	case "volume.update.end":
		fallthrough
	case "volume.delete.end":
		log.Printf("volume notification %s\n", m.Payload["volume_id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "volume.resize.end":
		log.Printf("%s\n", string(p.Message().Body))

	case "snapshot.create.end":
		fallthrough
	case "snapshot.update.end":
		fallthrough
	case "snapshot.delete.end":
		log.Printf("snapshot notification %s\n", m.Payload["snapshot_id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "snapshot.reset_status.end":
		log.Printf("%s\n", string(p.Message().Body))

	case "volume_type.create":
		fallthrough
	case "volume_type.update":
		fallthrough
	case "volume_type.delete":
		log.Printf("storage notification %s\n", m.Payload["volume_types"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "volume_type_project.access.add":
	case "volume_type_extra_specs.create":
	case "volume_type_extra_specs.delete":
	case "network.create.end":
		fallthrough
	case "network.update.end":
		fallthrough
	case "network.delete.end":
		log.Printf("network notification %s\n", m.Payload["network"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "subnet.create.end":
		fallthrough
	case "subnet.update.end":
		fallthrough
	case "subnet.delete.end":
		log.Printf("subnet notification %s\n", m.Payload["subnet"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "security_group.create.end":
		fallthrough
	case "security_group.update.end":
		fallthrough
	case "security_group.delete.end":
		log.Printf("sg notification %s\n", m.Payload["security_group"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "security_group_rule.create.end":
		fallthrough
	case "security_group_rule.update.end":
		fallthrough
	case "security_group_rule.delete.end":
		log.Printf("sg rule notification %s\n", m.Payload["security_group_rule"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "router.create.end":
		fallthrough
	case "router.update.end":
		fallthrough
	case "router.delete.end":
		log.Printf("router notification %s\n", m.Payload["router"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "router.interface.create", "router.interface.delete":
		log.Printf("router notification %s\n", m.Payload["router_interface"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "floatingip.create.end":
		fallthrough
	case "floatingip.delete.end":
		log.Printf("floating ip notification %s\n", m.Payload["floatingip"].(map[string]interface{})["id"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "floatingip.update.end":
		// floating ip attach instance
		if m.Payload["floatingip"].(map[string]interface{})["port_id"] != nil {
			log.Printf("floating ip notification port id := %s\n", m.Payload["floatingip"].(map[string]interface{})["port_id"].(string))
		} else {
			log.Printf("floating ip notification floating id := %s\n", m.Payload["floatingip"].(map[string]interface{})["id"].(string))
		}

		log.Printf("%s\n", string(p.Message().Body))

	case "port.update.end", "port.delete.end":
		// instance attach port interface
		if m.Payload["port"].(map[string]interface{})["device_owner"].(string) == "compute:nova" {
			log.Printf("port notification %s\n", m.Payload["port"].(map[string]interface{})["device_id"].(string))
			log.Printf("%s\n", string(p.Message().Body))
		}

		log.Printf("port notification %s\n", m.Payload["port"].(map[string]interface{})["device_owner"].(string))
		log.Printf("%s\n", string(p.Message().Body))

	case "scheduler.retype":
		log.Printf("%s\n", string(p.Message().Body))

	}
	if err != nil {
		log.Printf("Failed to sync cluster from event notification. cause: %v\n", err)
		return nil
	}

	return nil
}

// New 함수는 새로운 monitor interface 를 생성한다.
func New(serverURL string) monitor.Monitor {
	//TODO auth 의 경우 임시로 ID:PASSWORD(ex.guest:guest)를 쓰지만
	//	   사용자 입력에 의한 Cluster 의 MetaData 로 저장될 필요가 있음.
	//     마찬가지로 임시로 client 의 api server url 과 고정된 port(ex.192.168.1.1:5672) 를 쓰지만
	//     사용자 입력에 의한 Cluster 의 MetaData 로 저장될 필요가 있음.
	auth := "guest:guest"
	defaultPort := "5672"

	u, err := url.Parse(serverURL)
	if err != nil {
		logger.Error(err)
	}
	ip, _, err := net.SplitHostPort(u.Host)
	if err != nil {
		logger.Error(err)
	}

	return &notification{
		auth:    auth,
		address: fmt.Sprintf("%s:%s", ip, defaultPort),
	}
}

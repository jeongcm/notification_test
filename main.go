package main

import (
	"encoding/json"
	"fmt"
	"log"
	"test/monitor"
	"test/notification"
)

type OsloMessage struct {
	EventType string                 `json:"event_type"`
	Payload   map[string]interface{} `json:"payload"`
}

// OpenstackCredential 클러스터 인증정보
type OpenstackCredential struct {
	Methods  []string              `json:"methods,omitempty"`
	Token    OpenstackTokenAuth    `json:"token,omitempty"`
	Password OpenstackPasswordAuth `json:"password,omitempty"`
}

// OpenstackTokenAuth 클러스터 토큰 인증
type OpenstackTokenAuth struct {
	ID string `json:"id,omitempty"`
}

// OpenstackPasswordAuth 클러스터 비밀번호 인증
type OpenstackPasswordAuth struct {
	User ClusterPasswordAuthUser `json:"user,omitempty"`
}

// ClusterPasswordAuthUser 클러스터 인증 사용자
type ClusterPasswordAuthUser struct {
	Name     string                        `json:"name,omitempty"`
	Domain   ClusterPasswordAuthUserDomain `json:"domain,omitempty"`
	Password string                        `json:"password,omitempty"`
}

// ClusterPasswordAuthUserDomain 클러스터 인증 사용자의 도메인
type ClusterPasswordAuthUserDomain struct {
	Name string `json:"name,omitempty"`
}

func clusterNotificationSubscriber() {
	go func() {
		forever := make(chan bool)

		b := notification.NewBroker(fmt.Sprintf("%s:%s", "192.168.10.32", "5672"))
		if err := b.Connect(); err != nil {
			log.Printf("Could not register to cluster notification. Cause: %+v", err)
			return
		}
		defer func() {
			_ = b.Disconnect()
		}()

		s, err := b.Subscribe(1, "notifications.info", subscribeEvent, false)
		if err != nil {
			log.Printf("Could not register to cluster notification. Cause: %+v", err)
			return
		}
		defer func() {
			_ = s.Unsubscribe()
		}()

		<-forever
	}()
}

func subscribeEvent(clusterID int, p monitor.Event) error {
	var e map[string]interface{}

	err := json.Unmarshal(p.Message().Body, &e)
	if err != nil {
		log.Fatal(err)
	}

	var m OsloMessage
	if err := json.Unmarshal([]byte(e["oslo.message"].(string)), &m); err != nil {
		log.Fatal(err)
		return err
	}

	//TODO message type 별 동기화진행
	switch m.EventType {
	case "identity.project.created":
		log.Printf("project created\n")
	case "identity.project.updated":
		log.Printf("project updated\n")
	case "identity.project.deleted":
		log.Printf("project deleted\n")
	case "compute.instance.create.end":
	case "compute.instance.update":
	case "compute.instance.delete.end":
	case "volume.attach.end":
	case "compute.instance.suspend.end":
	case "volume.create.end":
		log.Printf("volume created\n")

	case "volume.update.end":
		log.Printf("volume updated\n")
	case "volume.delete.end":
		log.Printf("volume deleted\n")
	case "snapshot.create.end":
	case "snapshot.update.end":
	case "snapshot.delete.end":
	case "volume_type.create":
	case "volume_type.update":
	case "volume_type.delete":
	case "volume_type_project.access.add":
	case "volume_type_extra_specs.create":
	case "volume_type_extra_specs.delete":
	case "network.create.end":
	case "network.update.end":
	case "network.delete.end":
	case "subnet.create.end":
	case "subnet.update.end":
	case "security_group.create.end":
	case "security_group.update.end":
	case "security_group.delete.end":
	case "security_group_rule.create.end":
	case "security_group_rule.update.end":
	case "security_group_rule.delete.end":
	case "router.create.end":
	case "router.update.end":
	case "router.delete.end":
	case "router.interface.create":
	case "floatingip.create.end":
	case "floatingip.update.end":
	case "floatingip.delete.end":

	}

	return nil
}

func main() {
	//amqpServerURL := "amqp://guest:guest@192.168.10.32:5672/"
	//
	//// Create a new RabbitMQ connection.
	//connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	//if err != nil {
	//	panic(err)
	//}
	//defer connectRabbitMQ.Close()
	//
	//// Opening a channel to our RabbitMQ instance over
	//// the connection we have already established.
	//channelRabbitMQ, err := connectRabbitMQ.Channel()
	//if err != nil {
	//	panic(err)
	//}
	//defer channelRabbitMQ.Close()
	//// Subscribing to QueueService1 for getting messages.
	//messages, err := channelRabbitMQ.Consume(
	//	"notifications.info", // queue name
	//	"",              // consumer
	//	true,            // auto-ack
	//	false,           // exclusive
	//	false,           // no local
	//	false,           // no wait
	//	nil,             // arguments
	//)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//// Build a welcome message.
	//log.Println("Successfully connected to RabbitMQ")
	//log.Println("Waiting for messages")
	//
	//// Make a channel to receive messages into infinite loop.
	//forever := make(chan bool)
	//
	//go func() {
	//	for message := range messages {
	//		// For example, show received message in a console.
	//		log.Printf("%s\n", message.Body)
	//		var e map[string]interface{}
	//		if err := json.Unmarshal(message.Body, &e); err != nil {
	//			log.Fatalln(err)
	//			return
	//		}
	//
	//		var m OsloMessage
	//		if err := json.Unmarshal([]byte(e["oslo.message"].(string)), &m); err != nil {
	//			log.Fatalln(err)
	//			return
	//		}
	//
	//		fmt.Println(m.EventType)
	//		fmt.Println(m.Payload)
	//
	//		switch m.EventType {
	//		case "identity.project.created":
	//			fmt.Println(m.Payload["id"].(string))
	//		case "identity.project.deleted":
	//			fmt.Println(m.Payload["id"].(string))
	//		}
	//	}
	//}()
	//
	//<-forever

	//d := "{\"methods\": [\"password\"], \"password\": {\"user\": {\"domain\": {\"name\": \"default\"},\"name\": \"admin\",\"password\": \"admin\"}}}"
	//
	//var cred OpenstackCredential
	//if err := json.Unmarshal([]byte(d), &cred); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(cred)
	//
	//id := 1
	//
	//b, err := json.Marshal(&id)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//var cid int32
	//
	//if err = json.Unmarshal(b, &cid); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(cid)
	//
	//m := map[string]interface{}{"driver": "unknown"}
	//
	//_, ok := m["volume_backend_name"]
	//if !ok {
	//	fmt.Println("not found backend name")
	//	return
	//}
	//
	//
	//
	//fmt.Println(m["driver"].(string))
	//
	//	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	//	//defer cancel()
	//	i := make(chan interface{})
	//
	//	go func() {
	//		for {
	//
	//			fmt.Println("hihi")
	//
	//			select {
	//			case <-i:
	//				fmt.Println("wait done")
	//				return
	//			case <-time.After(time.Second * 1):
	//				continue
	//			}
	//		}
	//	}()

	//clusters := []string{"ee"}
	//
	//for _, c := range clusters {
	//	log.Printf("hello")
	//	log.Printf(c)
	//}
	forever := make(chan bool)
	clusterNotificationSubscriber()

	<-forever

}

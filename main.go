package main

import (
	"log"
	"test/monitor"
	_ "test/notification"
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

/*
func (ns *notificationSubscriber) subscribeEvent(p notification.Event) error {
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
*/
func main() {
	//amqpServerURL := "amqp://guest:guest@192.168.1.89:5672/"
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

	//_, err = channelRabbitMQ.QueueDeclare(
	//	"test",
	//	false, // durable
	//	true,  // autoDelete
	//	false, // exclusive
	//	false, // noWait
	//	nil,  // args
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	////if err = channelRabbitMQ.ExchangeDeclare(
	////	"openstack",
	////	"topic",
	////	false,
	////	false,
	////	false,
	////	false,
	////	nil,
	////); err != nil {
	////	panic(err)
	////}
	//
	//deliveries, err := channelRabbitMQ.Consume(
	//	"test",
	//	"",      // consumer
	//	false, // auto-ack
	//	false,   // exclusive
	//	false,   // no local
	//	false,   // no wait
	//	nil,     // arguments
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := channelRabbitMQ.QueueBind(
	//	"test", // queue
	//	"#",      // key
	//	"keystone",
	//	false,    // noWait
	//	nil,     // args
	//); err != nil {
	//	panic(err)
	//}
	//
	//
	//go func() {
	//	for message := range deliveries {
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
	//		case "volume.create.end":
	//			fallthrough
	//		case "volume.update.end":
	//			fallthrough
	//		case "volume.delete.end":
	//			fmt.Println(m.Payload["volume_id"].(string))
	//		}
	//	}
	//}()
	//deliveries2, err := channelRabbitMQ.Consume(
	//	"test2",
	//	"",      // consumer
	//	false, // auto-ack
	//	false,   // exclusive
	//	false,   // no local
	//	false,   // no wait
	//	nil,     // arguments
	//)
	//if err != nil {
	//	panic(err)
	//}
	//

	// Build a welcome message.
	//log.Println("Successfully connected to RabbitMQ")
	//log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	//forever := make(chan bool)
	//
	//go func() {
	//	for message := range deliveries2 {
	//		// For example, show received message in a console.
	//		//log.Printf("%s\n", message.Body)
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
	//		//fmt.Println(m.EventType)
	//		//fmt.Println(m.Payload)
	//
	//		switch m.EventType {
	//		case "identity.project.created":
	//			fmt.Println(m.Payload["id"].(string))
	//		case "identity.project.deleted":
	//			fmt.Println(m.Payload["id"].(string))
	//		case "volume.create.end":
	//			fmt.Println(m.Payload["volume_id"].(string))
	//		case "volume.update.end":
	//			fmt.Println(m.Payload["volume_id"].(string))
	//		case "volume.delete.end":
	//			fmt.Println(m.Payload["volume_id"].(string))
	//		case "network.create.end":
	//			fmt.Println(m.Payload["network"].(map[string]interface{})["id"].(string))
	//		case "network.update.end":
	//			fmt.Println(m.Payload["network"].(map[string]interface{})["id"].(string))
	//		case "network.delete.end":
	//			fmt.Println(m.Payload["network"].(map[string]interface{})["id"].(string))
	//		}
	//	}
	//}()

	forever := make(chan interface{})
	n, err := monitor.New("type.openstack", "http://192.168.10.32:5672")
	if err != nil {
		log.Fatalln("monitor init failed")
	}

	err = n.Start()
	if err != nil {
		log.Fatalf("monitor start failed. cause: %v\n", err)
	}

	//err = n.Connect()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//n.Stop()
	<-forever

}

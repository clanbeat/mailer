package main

import (
	"encoding/json"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/clanbeat/broker"
	"github.com/clanbeat/mailer/src/sender"
	"log"
)

func registerBrokerHandlers() {
	emailsQueueInfo := &broker.Consumer{
		Info: &broker.ConnectionInfo{
			ExchangeName: "emails",
			QueueName:    "mailer_outgoing_emails",
			RoutingKey:   "emails.#",
		},
		Callback: emailHandler(),
	}
	brokerConn.Register(emailsQueueInfo)
}

func emailHandler() func(msg broker.Delivery) {
	return func(msg broker.Delivery) {
		m := sender.Mailable{}
		if err := json.Unmarshal(msg.Body, &m); err != nil {
			errorTracker.Error(err)
		}

		log.Println("email", m)
		if err := sender.Send(&m); err != nil {
			errorTracker.Error(err)
		}
	}
}

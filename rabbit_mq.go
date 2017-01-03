package main


import (
  "fmt"
  "github.com/streadway/amqp"
  "log"
  "net/smtp"
)
/**
 * Fail intercept
 */
func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}
/**
 *
 */
func sendMail (){

  // Set up authentication information.
    auth := smtp.PlainAuth(
        "",
        "XXXXXXXXXXXXXXXXX",
        "XXXXXXXXXXXXXXXXX",
        "XXXXXXXXXXXXXXXXX",
    )

  getJson([]byte(`{"text":"I'm a text.","number":1234,"floats":[1.1,2.2,3.3],"innermap":{"foo":1,"bar":2}}`))
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
  failOnError(err, "Failed to connect to RabbitMQ")
  defer conn.Close()

  ch, err := conn.Channel()
  failOnError(err, "Failed to open a channel")
  defer ch.Close()

  msgs, err := ch.Consume(
      "action_q", // queue
      "",     // consumer
      true,   // auto-ack
      false,  // exclusive
      false,  // no-local
      false,  // no-wait
      nil,    // args
  )
  failOnError(err, "Failed to register a consumer")

  forever := make(chan bool)

  go func() {
    for d := range msgs {
       log.Printf(" [*] Waiting for messages. To exit press CTRL+C %s", d.Body)
       fmt.Println(getJson(d.Body))
        go sendMail(auth, getJson(d.Body))
    }
  }()
  //fake := `{"to":["fathallah.houssem@gmail.com"],"data":{"foo":1,"bar":2}}`
  // sendMail(auth, getJson([]byte(fake)))
  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}

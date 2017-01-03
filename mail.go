package main

import (
  "net/smtp"
  "log"
  "fmt"
	"encoding/base64"
  "text/template"
  "regexp"
  "bytes"
)
var reg, err = regexp.Compile("{{(\\s)*");

type Event struct {
    Body string `json:"body"`
    Data map[string]interface{}  `json:"data"`
}
/**
 *
 */
func _getTemplate(templateString string, data map[string]interface{}) []byte {
  replacement := reg.ReplaceAllString(templateString , "{{.")
  tmpl, err := template.New("test").Parse(replacement)
	if err != nil {
	  panic(err)
	}
	/*_data := make(map[string]map[string]interface{})
	_data["data"] = data*/
	buffer := new(bytes.Buffer)
	err = tmpl.Execute(buffer  , data)
	if err != nil { panic(err)
	}
	return buffer.Bytes()
}
/**
 *
 */
func _sendMail(auth smtp.Auth, recipients []string, message string ){
  // Connect to the server, authenticate, set the sender and recipient,
  // and send the email all in one step.
  err := smtp.SendMail(
      "smtp.gmail.com" + ":25",
      auth,
      "XXXXXXXXXX",
      recipients,
      []byte(message),
  )

  if err != nil {
      log.Fatal(err)
  }

}
/**
 *   []string -- {"recipient@example.net"}
  *  {string} message --
 */
func sendMail(auth smtp.Auth,  event TrackerResponse){
    header := make(map[string]string)
  	header["From"] ="XXXXXXXXXX"
  	header["Subject"] =  "event"
  	header["MIME-Version"] = "1.0"
  	header["Content-Type"] = "text/plain; charset=\"utf-8\""
  	header["Content-Transfer-Encoding"] = "base64"

    fmt.Println(">>>>>>>>>>>>>>>>>>")
    message := ""
  	for k, v := range header {
  		message += fmt.Sprintf("%s: %s\r\n", k, v)
  	}

    fmt.Println(">>>>>>>>>>>>>>>>>>")
    var currentTemplate []byte
    for _, action := range event.Actions {
      if (action.Type == "EMAIL"){
        // message + "\r\n" + base64.StdEncoding.EncodeToString([]byte(string(TrackerResponse))) +  "\r\n" + action.Body
        currentTemplate = _getTemplate(action.Body, event.Data)
        go _sendMail(auth, action.To,  message + "\r\n" + base64.StdEncoding.EncodeToString(currentTemplate))
      }
    }
}

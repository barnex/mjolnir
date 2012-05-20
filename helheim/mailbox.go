package helheim

import (
	"os/exec"
	"time"
	"fmt"
)

type Mailbox struct {
	email   string // send to this address
	message string
}

func(m*Mailbox)Post(message string){
	m.message += fmt.Sprint(time.Now(), message, "\n")
}

func (m *Mailbox) Sendmail() {
	defer func() {
		m.Clear()
		err := recover()
		if err != nil {
			Debug(err)
		}
	}()

	if m.email == "" || m.message == "" {
		return
	}

	sendmail := exec.Command("sendmail", m.email)
	stdin, _ := sendmail.StdinPipe()
	Check(sendmail.Start())
	_, err := stdin.Write(([]byte)(m.message))
	Check(err)
	Check(stdin.Close())
	Check(sendmail.Wait())
	m.message = ""
}

func(m*Mailbox)Clear(){
	m.message=""
}

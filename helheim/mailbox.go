package helheim

import (
	"fmt"
	"os/exec"
	"time"
)

type Mailbox struct {
	email     string // send to this address
	message   string
	firstpost time.Time // timestamp of first Posted message
	nSend int
}

func (m *Mailbox) Walltime() time.Duration {
	if m.firstpost.IsZero() {
		return 0
	}
	return time.Now().Sub(m.firstpost)
}

func (m *Mailbox) Println(message ...interface{}) {
	m.message += fmt.Sprint(time.Now(), ": ", fmt.Sprintln(message...))
	if m.firstpost.IsZero() {
		m.firstpost = time.Now()
	}
}

func (m *Mailbox) Sendmail() {
	defer func() {
		m.Clear()
		m.nSend++
		err := recover()
		if err != nil {
			Debug(err)
		}
	}()

	if m.email == "" || m.message == "" {
		return
	}

	Debug("sendmail", m.email, m.message)
	sendmail := exec.Command("mail", "-s", "[ragnarok] status", m.email)
	stdin, _ := sendmail.StdinPipe()
	Check(sendmail.Start())
	_, err := stdin.Write(([]byte)(m.message))
	Check(err)
	Check(stdin.Close())
	Check(sendmail.Wait())
	m.message = ""
}

func (m *Mailbox) Clear() {
	m.message = ""
	var zero time.Time
	m.firstpost = zero
}

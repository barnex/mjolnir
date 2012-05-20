package helheim

import (
	"os/exec"
)

type Mailbox struct {
	receipient string
	message    string
}

func (m *Mailbox) Sendmail() {
	defer func() {
		err := recover()
		if err != nil {
			Debug(err)
		}
	}()
	sendmail := exec.Command("sendmail", m.receipient)
	stdin, _ := sendmail.StdinPipe()
	Check(sendmail.Start())
	_, err := stdin.Write(([]byte)(m.message))
	Check(err)
	Check(stdin.Close())
	Check(sendmail.Wait())
	m.message = ""
}

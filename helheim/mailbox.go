package helheim

import(
	"os/exec"
)

type Mailbox struct{
	receipient string
	messages []string
}

func (m*Mailbox)Sendmail(){
	sendmail := exec.Command("sendmail", m.receipient)	
	stdin, _ := sendmail.StdinPipe()
	Check(sendmail.Start())
	for _,msg:=range m.messages{
		_, err := stdin.Write(([]byte)(msg))
		Check(err)
	}
	Check(stdin.Close())
	Check(sendmail.Wait())
	m.messages = []string{}
}

package helheim

import("testing")

func TestMail(t*testing.T){
	box:=Mailbox{"Arne.Vansteenkiste@UGent.be", []string{"hello mail!"}}
	box.Sendmail()
}

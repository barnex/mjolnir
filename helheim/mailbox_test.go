package helheim

import (
	"testing"
)

func TestMail(t *testing.T) {
	var box Mailbox
	box.email = "Arne.Vansteenkiste@UGent.be"
	box.Println("test1")
	box.Println("test2")
	box.Sendmail()
	box.Println("test3")
	box.Sendmail()
	box.Clear()
	box.Sendmail()
}

package helheim

import (
	"testing"
)

func TestMail(t *testing.T) {
	box := Mailbox{"Arne.Vansteenkiste@UGent.be", ""}
	box.Post("test1")
	box.Post("test2")
	box.Sendmail()
	box.Post("test3")
	box.Sendmail()
	box.Clear()
	box.Sendmail()
}

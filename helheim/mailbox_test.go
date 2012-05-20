package helheim

import (
	"testing"
)

func TestMail(t *testing.T) {
	box := Mailbox{"Arne.Vansteenkiste@UGent.be", "hello mail!"}
	box.Sendmail()
}

package helheim

import (
	"fmt"
	"io"
)

type Group struct {
	name  string
	share int // Relative cluster share of the group
	users []*User
}

type User struct {
	name  string
	share int // Relative group share of the user
}

func AddGroup(name string, share int) *Group {
	group := &Group{name, share, []*User{}}
	groups = append(groups, group)
	return group
}

func (g *Group) AddUser(name string, share int) {
	g.users = append(g.users, &User{name, share})
}

// API func, prints user info.
func Users(out io.Writer) error {
	for _, gr := range groups {
		fmt.Fprint(out, gr)
		for _, usr := range gr.users {
			fmt.Fprintln(out, "\n\t", usr)
		}
	}
	return nil
}

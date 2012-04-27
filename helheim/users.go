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

// Add new group to global list and return it as well.
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
		fmt.Fprint(out, gr.name, " (share ", gr.share, "/", TotalGroupShare(), ")\n")
		for _, usr := range gr.users {
			fmt.Fprint(out, "\t", usr.name, " (share ", usr.share, "/", gr.TotalUserShare(), ")\n")
		}
	}
	return nil
}

func TotalGroupShare() int {
	total := 0
	for _, gr := range groups {
		total += gr.share
	}
	return total
}

func (g *Group) TotalUserShare() int {
	total := 0
	for _, usr := range g.users {
		total += usr.share
	}
	return total
}

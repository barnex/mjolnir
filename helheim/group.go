package helheim

import (
	"fmt"
	"io"
)

// Group of users who have made a collective investment in the cluster.
type Group struct {
	name  string
	share int // Relative cluster share of the group
	users []*User
}

// Add new group to global list and return it as well.
func AddGroup(name string, share int) *Group {
	group := &Group{name, share, []*User{}}
	groups = append(groups, group)
	return group
}

// Add a new user to the group.
func (g *Group) AddUser(name string, share int) {
	if _, ok := users[name]; ok {
		panic("user " + name + " already added")
	}
	user := &User{name, share, 0, NewJobQueue(), g}
	g.users = append(g.users, user)
	users[name] = user
}

// Sum of shares of all groups.
func TotalGroupShare() int {
	total := 0
	for _, gr := range groups {
		total += gr.share
	}
	return total
}

// Sum of shares of all users in the group.
func (g *Group) TotalShare() int {
	total := 0
	for _, usr := range g.users {
		total += usr.share
	}
	return total
}

// Total number of running jobs by this group
func (g *Group) TotalUse() int {
	total := 0
	for _, usr := range g.users {
		total += usr.use
	}
	return total
}

// Roughly the fractional use of the group.
// Group with largest share wins if no jobs are running yet
func (g *Group) FracUse() float64 {
	return (float64(g.TotalUse()) + 1e-3) / float64(g.TotalShare())
}

// Returns if this group has any jobs queued
func (g *Group) HasJobs() bool {
	for _, usr := range g.users {
		if usr.HasJobs() {
			return true
		}
	}
	return false
}

// API func, prints user info.
func Groups(out io.Writer) error {
	for _, gr := range groups {
		fmt.Fprint(out, gr.name, " (share ", gr.share, "/", TotalGroupShare(), ")\n")
		for _, usr := range gr.users {
			fmt.Fprint(out, "\t", usr.name, " (share ", usr.share, "/", gr.TotalShare(), ")\n")
		}
	}
	return nil
}

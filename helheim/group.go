package helheim

import ()

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
	g.users = append(g.users, &User{name, share})
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
func (g *Group) TotalUserShare() int {
	total := 0
	for _, usr := range g.users {
		total += usr.share
	}
	return total
}

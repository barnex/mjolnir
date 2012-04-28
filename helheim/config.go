package helheim

func Configure() {
	dynamat := AddGroup("dynamat", 8)
	dynamat.AddUser("arne", 4)

	AddGroup("eelab", 0)

	AddNode([]string{"ssh", "localhost"})
}

package helheim

// executable
var MUMAX2 = "/home/arne/mumax2.git/bin/mumax2"

func Configure() {
	dynamat := AddGroup("dynamat", 8)
	dynamat.AddUser("arne", 2)
	dynamat.AddUser("test", 1)
	dynamat.AddUser("mykola", 2)
	dynamat.AddUser("jonas", 1)

	eelab := AddGroup("eelab", 4)
	eelab.AddUser("ben", 4)

	AddNode([]string{"ssh", "localhost"})
	AddNode([]string{"ssh", "dynamag"}) // faulty node
}

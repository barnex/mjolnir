package helheim

// executable
const (
	MUMAX2 = "/home/arne/mumax2.git/bin/mumax2"
	//MUMAX2 = "/home/mumax/mumax2/bin/mumax2"
	MUNINN = "/home/arne/go/bin/muninn"
	//MUNINN = "/home/mumax/go/bin/muninn"
)

func Configure() {
	dynamat := AddGroup("dynamat", 8)
	dynamat.AddUser("arne", 3)
	dynamat.AddUser("mykola", 3)
	dynamat.AddUser("jonas", 1)
	dynamat.AddUser("mathias", 1)

	eelab := AddGroup("eelab", 1)
	eelab.AddUser("ben", 1)

	AddNode("localhost", "ssh", "localhost")
	//	AddNode("fermi0" , "ssh", "192.168.0.2")
	//	AddNode("fermi1" , "ssh", "192.168.0.3")
	//	AddNode("kepler0", "ssh", "192.168.0.4")
	//	AddNode("kepler1", "ssh", "192.168.0.5")
}

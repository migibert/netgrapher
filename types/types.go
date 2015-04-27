package types

type Namespace struct {
	Name string
}

type Device struct {
	Name string
}

type Link struct {
	Device string
	Flags []string
	Mtu int
	ActiveQueuingMechanism string
	QueueLength int
	State string
	Mode string
	Group string
	Type string
	DeviceMacAddress string
	BroadcastMacAddress string
}
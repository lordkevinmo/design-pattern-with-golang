package factory

// Boat is a concrete implementation of Transport type
type Boat struct {
	Transport
}

// NewBoat Return instance of logistics method
func NewBoat() Logistics {
	return &Truck{
		Transport: Transport{
			name:   "Boat",
			engine: 1900,
		},
	}
}

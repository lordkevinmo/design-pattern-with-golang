package factory

// Truck is a concrete implementation of Transport data type
type Truck struct {
	Transport
}

// NewTruck Return instance of logistics method
func NewTruck() Logistics {
	return &Truck{
		Transport: Transport{
			name:   "Truck",
			engine: 700,
		},
	}
}

package parkingSlot

// PricingStrategy defines an interface for different pricing strategies
type PricingStrategy interface {
	CalculatePrice(hours int) float64
}

// HourlyPricing implements hourly pricing
type HourlyPricing struct {
	HourlyRate float64
}

// CalculatePrice calculates the price based on hours
func (hp *HourlyPricing) CalculatePrice(hours int) float64 {
	if hours <= 0 {
		return hp.HourlyRate // Minimum 1 hour charge
	}
	return float64(hours) * hp.HourlyRate
}

// GetPricingStrategy returns appropriate pricing strategy for vehicle type
func GetPricingStrategy(vehicleType VehicleType) PricingStrategy {
	switch vehicleType {
	case Car:
		return &HourlyPricing{HourlyRate: 2.0}
	case Motorcycle:
		return &HourlyPricing{HourlyRate: 1.0}
	case Bus:
		return &HourlyPricing{HourlyRate: 5.0}
	case Van:
		return &HourlyPricing{HourlyRate: 3.0}
	default:
		return &HourlyPricing{HourlyRate: 2.0}
	}
}

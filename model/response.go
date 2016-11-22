package model

type Response struct {
	Success   bool
	Error     bool
	Message   string
	Booking   Booking
	MealPrice int
	StripeKey string
}

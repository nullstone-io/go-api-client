package types

type Organization struct {
	Name               string `json:"name"`
	Seats              int    `json:"seats"`
	Plan               string `json:"plan"`
	SubscriptionStatus string `json:"subscriptionStatus"`
}

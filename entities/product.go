package entities

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	Email        string `json:"email"`
	Title        string `json:"title"`
	DisplayName  string `json:"display_name"`
	Height       int    `json:"height"`
	Weight       int    `json:"weight"`
	BodyFat      int    `json:"body_fat"`
	TargetWeight int    `json:"target_weight"`
	TestDate     string `json:"test_date"`
	NextTestDate string `json:"next_test_date"`
}

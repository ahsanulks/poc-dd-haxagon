package usecase

type RegistrationRequest struct {
	Name        string
	PhoneNumber string
}

type AddAddressRequest struct {
	UserID    int
	Street    string
	City      string
	ZipCode   string
	Latitude  float64
	Longitude float64
}

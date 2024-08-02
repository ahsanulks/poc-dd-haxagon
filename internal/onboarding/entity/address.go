package entity

type Address struct {
	ID        int
	Street    string
	City      string
	ZipCode   string
	Latitude  float64
	Longitude float64
}

type AddressBuilder struct {
	street    string
	city      string
	zipCode   string
	latitude  float64
	longitude float64
}

func (a *AddressBuilder) Street(street string) *AddressBuilder {
	a.street = street
	return a
}

func (a *AddressBuilder) City(city string) *AddressBuilder {
	a.city = city
	return a
}

func (a *AddressBuilder) ZipCode(zipCode string) *AddressBuilder {
	a.zipCode = zipCode
	return a
}

func (a *AddressBuilder) Latitude(latitude float64) *AddressBuilder {
	a.latitude = latitude
	return a
}

func (a *AddressBuilder) Longitude(longitude float64) *AddressBuilder {
	a.longitude = longitude
	return a
}

func (a *AddressBuilder) Build() (*Address, error) {
	if a.street == "" {
		return nil, ErrAddressStreetRequired
	}

	if a.city == "" {
		return nil, ErrAddressCityRequired
	}

	if a.zipCode == "" {
		return nil, ErrAddressZipCodeRequired
	}

	return &Address{
		Street:    a.street,
		City:      a.city,
		ZipCode:   a.zipCode,
		Latitude:  a.latitude,
		Longitude: a.longitude,
	}, nil
}

const (
	maximumAdress = 5
)

func (u *User) AddAddress(address *Address) error {
	if len(u.Addresses) == maximumAdress {
		return ErrUserAddressExceededLimit
	}
	u.Addresses = append(u.Addresses, address)
	return nil
}

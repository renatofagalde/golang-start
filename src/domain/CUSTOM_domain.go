package domain

type customDomain struct {
	id                      uint
	custom                   string
	fullName             string
	email       string
}

func (cd *customDomain) GetID() uint {
	return cd.id
}

func (cd *customDomain) GetCustom() string {
	return cd.custom
}

func (cd *customDomain) GetFullName() string {
	return cd.fullName
}

func (sd *customDomain) GetEmail() string {
	return sd.email
}
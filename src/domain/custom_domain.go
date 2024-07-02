package domain

type customDomain struct {
	id       string
	custom   string
	fullName string
	email    string
}

func (cd *customDomain) GetID() string {
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

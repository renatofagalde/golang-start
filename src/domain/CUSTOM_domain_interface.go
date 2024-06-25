package domain

func NewCustomDomain(
	id string,
	custom string,
	fullName string,
	email string,
) CustomDomainInterface {
	return &customDomain{
		id:       id,
		custom:   custom,
		fullName: fullName,
		email:    email,
	}
}

type CustomDomainInterface interface {
	GetID() string
	GetCustom() string
	GetFullName() string
	GetEmail() string
}

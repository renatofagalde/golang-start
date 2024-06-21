package model

func NewCustomDomain(ID string, custom string, fullName string, email string,) CustomDomainInterface {
	return &customDomain{
		id: id,
		custom:	custon,
		fullName:                   fullName,
		email:             email,
	}
}

type CustomDomainInterface interface {
	GetID() string
	GetFullName() string
	GetEmail() string
}

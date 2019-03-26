package gblcodes

const (
	// AuthInvalidJWT denotes that the JWT passed is invalid
	AuthInvalidJWT appCode = "authentication.invalid_jwt"
	// AuthInvalidScope denotes that the caller (user) is modifying or
	// creating a record that's not scoped to their JWT issuer
	AuthInvalidScope appCode = "authenticaiton.invalid_scope"
)

var (
	authMessages = map[AppCode]string{
		AuthInvalidJWT:   "Invalid JWT signature. This can mean the JWT is malformed or invalid (secret was used to validate it).",
		AuthInvalidScope: "The JWT's issuer does not match the scope allowed for this user",
	}
)

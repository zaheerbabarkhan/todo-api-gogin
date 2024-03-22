package types

type accountType struct {
	APP    string
	SOCIAL string
}

var AccountTypes = accountType{
	SOCIAL: "SOCIAL",
	APP:    "APP",
}

package v1

// Function responsible for validating a bearerToken
// and --
func ValidateBearerToken(bearerToken string) bool {
	if (bearerToken == "5230-SF20b-&21c1-%8vs1x41sd"){
		return true
	}
	return false 
}


// Responsible for generating a bearer Token and 
// storing the token with the apropriate user 
// upon login. 
func GenerateBearerToken() string {
	return "5230-SF20b-&21c1-%8vs1x41sd"
}
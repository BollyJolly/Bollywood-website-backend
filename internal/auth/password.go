package auth

import "golang.org/x/crypto/bcrypt"


// HashPassword converts a plain-text password
// into a secure hash before saving it to MongoDB.
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}


// ComparePassword checks whether the password entered
// by the user matches the password stored in the database.
func ComparePassword(
	password string,
	hashedPassword string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}
package errs

import "errors"

var (
	InvalidBearerFormat = errors.New("Invalid Authorization Bearer Format")
	InvalidToken        = errors.New("Invalid Token")
	InvalidIssuer       = errors.New("Invalid Token Issuer")
	InvalidIDParam      = errors.New("Invalid ID Parameter")

	ForbiddenAccess = errors.New("user is forbidden to access this resource")

	InvalidRequestBody = errors.New("invalid request body")

	EmailAlreadyUsed              = errors.New("email already used")
	UsernameAlreadyUsed           = errors.New("username already used")
	PasswordDoesntMatch           = errors.New("password doesn't match")
	PasswordContainUsername       = errors.New("password must not contain username")
	PasswordSameAsBefore          = errors.New("Password cannot be same as before")
	UsernamePasswordIncorrect     = errors.New("username or password incorrect")
	ErrorGeneratingHashedPassword = errors.New("Error Generating Hashed Password")
	ErrorCreatingAccount = errors.New("Error Creating Account")


	InvalidAccountId          = errors.New("Invalid Account Id")
	ErrorOnCreatingNewProfile = errors.New("Error on Creating new profile")
	ErrorOnDeletingProfile    = errors.New("Error on Deleting profile")
	InvalidProfileIDParam     = errors.New("invalid Profile ID param")
	InvalidProfileID          = errors.New("invalid Profile ID")

	
	ErrorOnCreatingNewFarm = errors.New("Error on Creating new farm")
	ErrorOnDeletingFarm    = errors.New("Error on Deleting farm")
	InvalidFarmIDParam     = errors.New("invalid Farm ID param")
	InvalidFarmID          = errors.New("invalid Farm ID")

	ErrorOnCreatingNewSystemUnit = errors.New("Error on Creating new system unit")
	ErrorOnDeletingSystemUnit    = errors.New("Error on Deleting system unit")
	InvalidSystemUnitIDParam     = errors.New("invalid system unit ID param")
	InvalidSystemUnitID          = errors.New("invalid system unit ID")
)

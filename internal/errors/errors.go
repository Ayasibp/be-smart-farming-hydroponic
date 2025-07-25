package errs

import "errors"

var (
	InvalidBearerFormat = errors.New("Invalid Authorization Bearer Format")
	InvalidToken        = errors.New("Invalid Token")
	InvalidIssuer       = errors.New("Invalid Token Issuer")
	InvalidIDParam      = errors.New("Invalid ID Parameter")

	ForbiddenAccess = errors.New("user is forbidden to access this resource")

	InvalidRequestBody = errors.New("invalid request body")

	ErrorOnParsingStringToUUID = errors.New("error on parsing string to uuid")
	ParseUUIDError             = errors.New("Error occurred while parsing UUID")

	ErrorCreatingSystemLog = errors.New("Error on creating system log")

	EmailAlreadyUsed              = errors.New("email already used")
	UsernameAlreadyUsed           = errors.New("username already used")
	PasswordDoesntMatch           = errors.New("password doesn't match")
	PasswordContainUsername       = errors.New("password must not contain username")
	PasswordSameAsBefore          = errors.New("Password cannot be same as before")
	UsernamePasswordIncorrect     = errors.New("username or password incorrect")
	ErrorGeneratingHashedPassword = errors.New("Error Generating Hashed Password")
	ErrorCreatingAccount          = errors.New("Error Creating Account")

	InvalidAccountId          = errors.New("Invalid Account Id")
	ErrorOnCreatingNewProfile = errors.New("Error on Creating new profile")
	ErrorOnCheckingProfile    = errors.New("Error on Checking profile")
	ErrorOnDeletingProfile    = errors.New("Error on Deleting profile")
	InvalidProfileIDParam     = errors.New("invalid Profile ID param")
	InvalidProfileID          = errors.New("invalid Profile ID")
	ProfileAlreadyCreated     = errors.New("Profile already created")

	ErrorOnCreatingNewFarm = errors.New("Error on Creating new farm")
	ErrorOnDeletingFarm    = errors.New("Error on Deleting farm")
	InvalidFarmIDParam     = errors.New("invalid Farm ID param")
	InvalidFarmID          = errors.New("invalid Farm ID")

	ErrorOnCreatingNewSystemUnit = errors.New("Error on Creating new system unit")
	ErrorOnDeletingSystemUnit    = errors.New("Error on Deleting system unit")
	InvalidSystemUnitIDParam     = errors.New("invalid system unit ID param")
	InvalidSystemUnitID          = errors.New("invalid system unit ID")

	InvalidUnitKey      = errors.New("invalid Unit Key")
	InvalidUnitKeyParam = errors.New("invalid Unit Key param")

	InvalidId = errors.New("Invalid Account Id")

	ErrorOnCreatingNewTankTrans = errors.New("Error on Creating new tank transaction")

	ErrorOnCreatingNewGrowthHist  = errors.New("Error on Creating new growth hist")
	EmptyPeriodQueryParams        = errors.New("Empty period query params")
	EmptyFarmIdParams             = errors.New("Empty farm_id query params")
	EmptySystemIdParams           = errors.New("Empty system_id query params")
	EmptyStartDateQueryParams     = errors.New("Empty start_date query param")
	EmptyEndDateQueryParams       = errors.New("Empty end_date query param")
	InvalidValuePeriodQueryParams = errors.New("Invalid Period Value")
	StartDateExceedEndDate        = errors.New("start_date exceed end_date")
	ErrorOnGettingAggregatedData  = errors.New("error on getting aggregated data")
)

package constants

const (
	// Client Error Messages
	MsgInvalidID              = "The provided ID is invalid."
	MsgInvalidEmail           = "The provided email is invalid."
	MsgInvalidOTP             = "The provided OTP is invalid."
	MsgEmailNotFound          = "The provided email doesn't exist."
	MsgEventNotFound          = "No event was found with the corresponding ID."
	MsgNotFound               = "The requested resource was not found."
	MsgUserFound              = "The requested user was not found."
	MsgUnauthorized           = "You are not authorized to perform this action."
	MsgInternalServerError    = "An unexpected server error occurred."
	MsgBadRequest             = "The request is invalid."
	MsgDifferentPayement      = "The cart event has different payment method. Please checkout seperately."
	MsgForbidden              = "You do not have permission to access this resource."
	MsgConflict               = "There is a conflict with the current state of the resource."
	MsgUnprocessableEntity    = "The request cannot be processed."
	MsgTooManyRequests        = "You have made too many requests. Please try again later."
	MsgServiceUnavailable     = "The service is currently unavailable. Please try again later."
	MsgInvalidInput           = "The input provided is invalid."
	MsgDuplicateEntry         = "The entry already exists."
	MsgResourceLocked         = "The resource is locked."
	MsgUnsupportedMedia       = "The media type is not supported."
	MsgExpiredToken           = "The token has expired."
	MsgInvalidCredentials     = "The credentials provided are invalid."
	MsgMissingParameter       = "A required parameter is missing."
	MsgValidationError        = "Validation failed for the provided input."
	MsgSessionExpired         = "Your session has expired."
	MsgAccountDisabled        = "The account is disabled."
	MsgPermissionDenied       = "Permission denied."
	MsgConfirmPasswordInvalid = "Password and confirmation do not match."
	MsgInvalidPassword        = "Incorrect password. Please check your credentials."

	// Server Error Messages
	MsgDatabaseError       = "A database error occurred."
	MsgCacheError          = "A cache error occurred."
	MsgTimeout             = "The operation timed out."
	MsgDependencyFailure   = "A dependency failed to respond."
	MsgNetworkError        = "A network error occurred."
	MsgConfigurationError  = "There is an issue with the configuration."
	MsgUnknownError        = "An unknown error occurred."
	MsgOperationFailed     = "The operation failed."
	MsgDataIntegrityError  = "Data integrity validation failed."
	MsgFileProcessingError = "There was an error processing the file."
	MsgServiceError        = "A service error occurred."

	// Resource-Related Messages
	MsgResourceNotAvailable  = "The requested resource is not available."
	MsgResourceQuotaExceeded = "The resource quota has been exceeded."
	MsgResourceAlreadyExists = "The resource already exists."
	MsgResourceDeleted       = "The resource has been deleted."
	MsgInvalidResourceState  = "The resource is in an invalid state."

	// Authentication/Authorization Messages
	MsgInvalidToken          = "The token is invalid."
	MsgAccessTokenExpired    = "The access token has expired."
	MsgRefreshTokenExpired   = "The refresh token has expired."
	MsgInvalidSession        = "The session is invalid."
	MsgIPBlocked             = "Access is blocked from your IP address."
	MsgMFARequired           = "Multi-factor authentication is required."
	MsgMFAVerificationFailed = "Multi-factor authentication verification failed."

	// File and Upload Error Messages
	MsgFileTooLarge      = "The file is too large."
	MsgInvalidFileFormat = "The file format is invalid."
	MsgFileUploadFailed  = "The file upload failed."
	MsgFileNotFound      = "The file was not found."

	// Rate Limiting/Quota Messages
	MsgRateLimitExceeded = "Rate limit exceeded. Try again later."
	MsgQuotaExceeded     = "Quota exceeded. Please upgrade or wait for reset."

	// Payment and Billing Error Messages
	MsgPaymentRequired     = "Payment is required to proceed."
	MsgPaymentFailed       = "The payment process failed."
	MsgCardDeclined        = "The card was declined."
	MsgInsufficientFunds   = "Insufficient funds to complete the transaction."
	MsgSubscriptionExpired = "The subscription has expired."
	MsgPaymentNotFound     = "No payment method found"

	// Pagination Error Messages
	// Input validation
	MsgInvalidPageNumber = "The page number provided is invalid."
	MsgInvalidPageSize   = "The page size provided is invalid."

	// Out-of-range checks
	MsgPageOutOfRange = "The requested page is out of range."
	MsgNoMorePages    = "There are no more pages available."

	// Internal or system-related
	MsgPaginationError   = "An error occurred while processing pagination."
	MsgInvalidOffset     = "The calculated offset is invalid or exceeds the dataset."
	MsgNegativePageValue = "Page number or size must be a positive integer."
)

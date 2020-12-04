package ftp

const (
	restartMarkerReply         = "110 Restart marker reply."
	readyInMinutes             = "120 Service ready in nnn minutes."
	alreadyOpen                = "125 Data connection already open; transfer starting."
	fileStatusOk               = "150 File status okay; about to open data connection."
	ok                         = "200 Command okay."
	notImplementedAtThisSite   = "202 Command not implemented, superfluous at this site."
	systemStatus               = "211 System status, or system help reply."
	directoryStatus            = "212 Directory status."
	fileStatus                 = "213 File status."
	helpMessage                = "214 Help message."
	nameSystemType             = "215 %s system type."
	readyForNewUser            = "220 Service ready for new user."
	closingControlConnection   = "221 Service closing control connection."
	connectionOpen             = "225 Data connection open; no transfer in progress."
	closingDataConnection      = "226 Closing data connection."
	enteringPassiveMode        = "227 Entering Passive Mode (%s)."
	userLoggedIn               = "230 User logged in, proceed."
	fileActionOk               = "250 Requested file action okay, completed."
	created                    = "257 %q created."
	needPassword               = "331 User name okay, need password."
	needAccountForLogin        = "332 Need account for login."
	peding                     = "350 Requested file action pending further information."
	localError                 = "421 Service not available, closing control connection."
	notAvailable               = "425 Can't open data connection."
	cantOpenConnection         = "426 Connection closed; transfer aborted."
	connectionClosed           = "450 Requested file action not taken."
	localErrorrInProcessing    = "451 Requested action aborted: local error in processing."
	unavailableFile            = "452 Requested action not taken."
	wrongCommand               = "500 Syntax error, command unrecognized."
	wrongArguments             = "501 Syntax error in parameters or arguments."
	notImplemented             = "502 Command not implemented."
	badSequence                = "503 Bad sequence of commands."
	notImplementedForParameter = "504 Command not implemented for that parameter."
	notLoggedIn                = "530 Not logged in."
	needAccountForStoringFiles = "532 Need account for storing files."
	fileNotFound               = "550 Requested action not taken."
	unknownPageType            = "551 Requested action aborted: page type unknown."
	notEnoughSpace             = "552 Requested file action aborted."
	disallowedFileName         = "553 Requested action not taken."
)

type response struct {
	message string
	data    string
	hasData bool
	closing bool
}

func newResponse(message string) response {
	return response{message: message}
}

func (res *response) SetData(data string) {
	res.hasData = true
	res.data = data
}

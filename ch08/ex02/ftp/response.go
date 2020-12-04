package ftp

import (
	"fmt"
	"strings"
)

type responseCode int

const (
	restartmarkerReplay        responseCode = 110
	readyInMinutes             responseCode = 120
	alreadyOpen                responseCode = 125
	fileStatusOk               responseCode = 150
	ok                         responseCode = 200
	notImplementedAtThisSite   responseCode = 202
	systemStatus               responseCode = 211
	directoryStatus            responseCode = 212
	fileStatus                 responseCode = 213
	helpMessage                responseCode = 214
	nameSystemType             responseCode = 215
	readyForNewUser            responseCode = 220
	closingControlConnection   responseCode = 221
	connectionOpen             responseCode = 225
	closingDataConnection      responseCode = 226
	enteringPassiveMode        responseCode = 227
	userLoggedIn               responseCode = 230
	fileActionOk               responseCode = 250
	created                    responseCode = 257
	needPassword               responseCode = 331
	needAccountForLogin        responseCode = 332
	peding                     responseCode = 350
	localError                 responseCode = 351
	notAvailable               responseCode = 421
	cantOpenConnection         responseCode = 425
	connectionClosed           responseCode = 426
	unavailableFile            responseCode = 450
	wrongCommand               responseCode = 500
	wrongArguments             responseCode = 501
	notImplemented             responseCode = 502
	badSequence                responseCode = 503
	notImplementedForParameter responseCode = 504
	notLoggedIn                responseCode = 530
	needAccountForStoringFiles responseCode = 532
	fileNotFound               responseCode = 550
	unknownPageType            responseCode = 551
	notEnoughSpace             responseCode = 552
	disallowedFileName         responseCode = 553
)

type response struct {
	code    responseCode
	message string
}

func newResponse(code responseCode, messages ...string) response {
	return response{
		code:    code,
		message: strings.Join(messages, " "),
	}
}

func (res *response) String() string {
	if res.message == "" {
		return fmt.Sprint(res.code)
	}
	return fmt.Sprintf("%d %s", res.code, res.message)
}

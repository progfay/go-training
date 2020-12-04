package ftp

import (
	"fmt"
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
	data    string
	hasData bool
}

func newResponse(code responseCode, message string) response {
	return response{
		code:    code,
		message: message,
	}
}

func (res *response) SetData(data string) {
	res.hasData = true
	res.data = data
}

func (res *response) Send(conn ftpConn) {
	if res.message == "" {
		fmt.Fprintf(conn.ctrlConn, "%d\n", res.code)
	} else {
		fmt.Fprintf(conn.ctrlConn, "%d %s\n", res.code, res.message)
	}

	if res.hasData {
		fmt.Fprintf(conn.dataConn, "%s\r\n", res.data)
		conn.dataConn.Close()
		fmt.Fprintf(conn.ctrlConn, "%d\n", closingDataConnection)
	}
}

package ftp

type state struct {
	name      string
	port      string
	printType string
	mode      string
	stru      string
}

func newState() state {
	return state{
		name:      "anonymous",
		port:      "",
		printType: "ASCII Non-print",
		mode:      "stream",
		stru:      "file",
	}
}

func (s *state) handle(req request) response {
	switch req.command {
	case "USER":
		s.name = req.message
		return newResponse(userLoggedIn, "User logged in")

	// case "QUIT":

	case "PORT":
		s.port = req.message
		return newResponse(ok, "Okay")

	case "LIST":
		return newResponse(fileStatusOk, "hoge", "fuga")

	// case "TYPE":
	// case "MODE":
	// case "STRU":
	// case "RETR":
	// case "STOR":
	// case "NOOP":

	default:
		return newResponse(notImplementedAtThisSite)
	}
}

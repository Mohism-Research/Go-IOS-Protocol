package iostdb

type Request struct {
	Time    int64
	From    string
	To      string
	ReqType int
	Body    []byte
}

type Response struct {
	Time        int64
	From        string
	To          string
	Code        int
	Description string
}

//go:generate mockgen -destination mocks/mock_network.go -package iosbase_mock -source network.go -imports .=github.com/iost-official/Go-IOS-Protocol/iosbase

type Net interface {
	Send(req Request) chan Response
	Listen(port uint16) (chan Request, chan Response, error)
	Close(port uint16) error
}

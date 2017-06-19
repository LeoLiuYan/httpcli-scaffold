package httpcli

import (
	"httpcli-scaffold/lib/encoding"
	"net/http"
)

var defaultErrorMapper = ErrorMapperFunc(FromResponse)

type requestOpt func(*http.Request)

type doFunc func(r *http.Request) (*http.Response, error)

type ErrorMapperFunc func(response *http.Response) error

type Client struct {
	url                  string
	methor               string
	header               http.Header
	requestOpts          []requestOpt
	do                   doFunc
	codec                encoding.Codec
	errorMapper          ErrorMapperFunc
	buildRequestWithBody func(encoding.Marshaler, ...requestOpt) (*http.Request, error)
	buildRequest         func(...requestOpt) (*http.Request, error)
}

func (c *Client) BuildRequestWithBody(marshaler encoding.Marshaler, opt ...requestOpt) (*http.Request, error) {

}

func (c *Client) BuildRequest(opt ...requestOpt) (*http.Request, error) {

}

func New(opts ...Opt) *Client {
	cli := &Client{
		header:      http.Header{},
		codec:       encoding.ProtobufCodec,
		errorMapper: defaultErrorMapper,
	}
	cli.buildRequestWithBody = cli.BuildRequestWithBody
	cli.buildRequest = cli.BuildRequest
}

package httpcli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Code int

func (c Code) String() string {
	return strconv.Itoa(int(c))
}

func (c Code) IsError() bool {
	return c >= 300
}

type Error struct {
	Code    Code
	Details string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Details)
}

func FromResponse(response *http.Response) error {
	if response == nil {
		return nil
	}

	c := Code(response.StatusCode)
	if !c.IsError() {
		return nil
	}

	e := &Error{
		Code: c,
	}
	if response.Body != nil {
		res, _ := ioutil.ReadAll(response.Body)
		e.Details = string(res)
		response.Body.Close()
	}

	return e
}

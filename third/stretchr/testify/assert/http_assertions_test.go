package assert

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func httpOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func httpRedirect(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func httpError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func TestHTTPSuccess(t *testing.T) {
	assert := New(t)

	mockT1 := new(testing.T)
	Equal(HTTPSuccess(mockT1, httpOK, "GET", "/", nil), true)
	False(mockT1.Failed())

	mockT2 := new(testing.T)
	Equal(HTTPSuccess(mockT2, httpRedirect, "GET", "/", nil), false)
	True(mockT2.Failed())

	mockT3 := new(testing.T)
	Equal(HTTPSuccess(mockT3, httpError, "GET", "/", nil), false)
	True(mockT3.Failed())
}

func TestHTTPRedirect(t *testing.T) {
	assert := New(t)

	mockT1 := new(testing.T)
	Equal(HTTPRedirect(mockT1, httpOK, "GET", "/", nil), false)
	True(mockT1.Failed())

	mockT2 := new(testing.T)
	Equal(HTTPRedirect(mockT2, httpRedirect, "GET", "/", nil), true)
	False(mockT2.Failed())

	mockT3 := new(testing.T)
	Equal(HTTPRedirect(mockT3, httpError, "GET", "/", nil), false)
	True(mockT3.Failed())
}

func TestHTTPError(t *testing.T) {
	assert := New(t)

	mockT1 := new(testing.T)
	Equal(HTTPError(mockT1, httpOK, "GET", "/", nil), false)
	True(mockT1.Failed())

	mockT2 := new(testing.T)
	Equal(HTTPError(mockT2, httpRedirect, "GET", "/", nil), false)
	True(mockT2.Failed())

	mockT3 := new(testing.T)
	Equal(HTTPError(mockT3, httpError, "GET", "/", nil), true)
	False(mockT3.Failed())
}

func TestHTTPStatusesWrapper(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	Equal(HTTPSuccess(httpOK, "GET", "/", nil), true)
	Equal(HTTPSuccess(httpRedirect, "GET", "/", nil), false)
	Equal(HTTPSuccess(httpError, "GET", "/", nil), false)

	Equal(HTTPRedirect(httpOK, "GET", "/", nil), false)
	Equal(HTTPRedirect(httpRedirect, "GET", "/", nil), true)
	Equal(HTTPRedirect(httpError, "GET", "/", nil), false)

	Equal(HTTPError(httpOK, "GET", "/", nil), false)
	Equal(HTTPError(httpRedirect, "GET", "/", nil), false)
	Equal(HTTPError(httpError, "GET", "/", nil), true)
}

func httpHelloName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func TestHTTPRequestWithNoParams(t *testing.T) {
	var got *http.Request
	handler := func(w http.ResponseWriter, r *http.Request) {
		got = r
		w.WriteHeader(http.StatusOK)
	}

	True(t, HTTPSuccess(t, handler, "GET", "/url", nil))

	Empty(t, got.URL.Query())
	Equal(t, "/url", got.URL.RequestURI())
}

func TestHTTPRequestWithParams(t *testing.T) {
	var got *http.Request
	handler := func(w http.ResponseWriter, r *http.Request) {
		got = r
		w.WriteHeader(http.StatusOK)
	}
	params := url.Values{}
	params.Add("id", "12345")

	True(t, HTTPSuccess(t, handler, "GET", "/url", params))

	Equal(t, url.Values{"id": []string{"12345"}}, got.URL.Query())
	Equal(t, "/url?id=12345", got.URL.String())
	Equal(t, "/url?id=12345", got.URL.RequestURI())
}

func TestHttpBody(t *testing.T) {
	assert := New(t)
	mockT := new(testing.T)

	True(HTTPBodyContains(mockT, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	True(HTTPBodyContains(mockT, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World"))
	False(HTTPBodyContains(mockT, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))

	False(HTTPBodyNotContains(mockT, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	False(HTTPBodyNotContains(mockT, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World"))
	True(HTTPBodyNotContains(mockT, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))
}

func TestHttpBodyWrappers(t *testing.T) {
	assert := New(t)
	mockAssert := New(new(testing.T))

	True(HTTPBodyContains(httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	True(HTTPBodyContains(httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World"))
	False(HTTPBodyContains(httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))

	False(HTTPBodyNotContains(httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	False(HTTPBodyNotContains(httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World"))
	True(HTTPBodyNotContains(httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))

}

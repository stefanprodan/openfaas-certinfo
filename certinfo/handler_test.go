package function

import (
	"regexp"
	"testing"
)

func TestHandleReturnsCorrectResponse(t *testing.T) {
	expected := "Google Internet Authority"
	resp := Handle([]byte("www.google.com/about/"))

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}

func TestHandleReturnsMultiSanResponse(t *testing.T) {
	expected := ".stefanprodan.com"
	resp := Handle([]byte("stefanprodan.com"))

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}

func TestHandleReturnsErrorResponse(t *testing.T) {
	expected := "connection refused"
	resp := Handle([]byte("http://fscked.org"))

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}

func TestHandleTimeoutResponse(t *testing.T) {
	expected := "timeout"
	resp := Handle([]byte("alexellis.io"))

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}

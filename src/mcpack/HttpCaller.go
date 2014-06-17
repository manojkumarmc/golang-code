package mcpack

import (
	"net/http"
	"io/ioutil"
	"log"
	"errors"
)

func Call(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Unable to call the method")
	}

	respo, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return "", errors.New("Unable to read the body")
	}
	defer resp.Body.Close()

	return string(respo[:]), nil // byte to string
}


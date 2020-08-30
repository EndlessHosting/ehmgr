package ehmgr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) doRequest(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", "ehmgr-go")
	r.Header.Set("ApiKey", c.ApiKey)
	httpCli := &http.Client{}
	return httpCli.Do(r)
}

func (c *Client) ListPackages() (PackageList, error) {
	uri := fmt.Sprintf("%s/api/package/list", c.Host)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accepts", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // Ensure the body is closed when we're done
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return UnmarshalPackageList(body)
}

func (c *Client) CreateUser(nu *NewUser) (*AcctResponse, error) {
	uri := fmt.Sprintf("%s/api/acct/create", c.Host)

	formData := url.Values {
		"username": {nu.Username},
		"password": {nu.Password},
		"email": {nu.Email},
		"domain": {nu.Domain},
		"packageName": {nu.Package},
		"ip": {nu.IP},
		"notify": {fmt.Sprintf( "%t" ,nu.Notify)},
	}
	req, err := http.NewRequest("POST", uri, strings.NewReader(formData.Encode())) // Beam me up Scotty
	if err != nil {
		return nil, err
	}
	err = req.ParseForm()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(formData.Encode())))
	req.Header.Set("Accepts", "application/json")


	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // Ensure the body is closed when we're done
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return UnmarshalAcctResponse(body) // Nothing went wrong
}
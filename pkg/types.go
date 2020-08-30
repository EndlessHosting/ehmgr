package ehmgr

import (
	"bytes"
	"encoding/json"
)

type NewUser struct {
	Username string`json:"username"`
	Password string`json:"password"`
	Email    string`json:"email"`
	Domain   string`json:"domain"`
	Package  string`json:"packageName"`
	IP		 string`json:"ip"`
	Notify   bool  `json:"notify"`
}

func CreateNewUser(username, email, password, pack, baseUserDomain, ip string) *NewUser {
	return &NewUser{
		Username: username,
		Password: password,
		Email: email,
		Package: pack,
		Domain: username + "." + baseUserDomain,
		IP: ip,
		Notify: false,
	}
}

func (nu *NewUser) Marshal() ([]byte, error) {
	return json.Marshal(nu)
}

type PackageList []string

func UnmarshalPackageList(j []byte) (PackageList, error) {
	pl := new(PackageList)
	decoder := json.NewDecoder(bytes.NewBuffer(j))
	err := decoder.Decode(&pl)
	if err != nil {
		return nil, err
	}

	return *pl, nil
}

type Package struct {
	Name  string
	Value string
}

type Client struct {
	Host 	   string
	ApiKey     string
}

func NewClient(host, apiKey string) *Client {
	return &Client{
		host,
		apiKey,
	}
}

type AcctResponse struct {
	Error bool`json:"error"`
	Text string`json:"text"`
	Details string`json:"details"`
	OriginalResponse string`json:"originalResponse"`
}

func UnmarshalAcctResponse(j []byte) (*AcctResponse, error) {
	resp := new(AcctResponse)
	decoder := json.NewDecoder(bytes.NewBuffer(j))
	err := decoder.Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
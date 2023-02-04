package solowaysdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	Username    string
	Password    string
	tr          http.Client
	xSid        string
	AccountInfo AccountInfo
}

func NewClient(username string, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
		tr:       http.Client{},
	}
}

func (c *Client) Login() (err error) {
	param := make(map[string]string)
	param["username"] = c.Username
	param["password"] = c.Password
	body, err := buildBody(param)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, Host+string(Login), body)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.tr.Do(req)
	if err != nil {
		return err
	}
	c.xSid = resp.Header["X-Sid"][0]
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := ReqUserInfo{}
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return err
	}
	return nil
}

type ReqUserInfo struct {
	Username string `json:"username"`
}

func (c *Client) Whoami() (err error) {
	if !checkSeed(c.xSid) {
		return fmt.Errorf("авторизация не пройдена")
	}
	req, err := http.NewRequest(http.MethodGet, Host+string(Whoami), nil)
	if err != nil {
		return err
	}
	buildHeader(req, c.xSid)
	resp, err := c.tr.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := AccountInfo{}
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return err
	}
	c.AccountInfo = data
	return nil
}

func (c *Client) GetPlacements() (placements PlacementsInfo, err error) {
	if !checkSeed(c.xSid) {
		return PlacementsInfo{}, fmt.Errorf("авторизация не пройдена")
	}
	if c.AccountInfo.Username == "" {
		return PlacementsInfo{}, fmt.Errorf("инфо об аккаунте не получено")
	}
	req, err := http.NewRequest(http.MethodGet, Host+"/api/clients/"+c.AccountInfo.Client.Guid+"/placements", nil)
	if err != nil {
		return PlacementsInfo{}, err
	}
	buildHeader(req, c.xSid)
	resp, err := c.tr.Do(req)
	if err != nil {
		return PlacementsInfo{}, err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PlacementsInfo{}, err
	}
	data := PlacementsInfo{}
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return PlacementsInfo{}, err
	}
	return data, nil
}

func (c *Client) GetPlacementsStat(placementIds []string, startDate time.Time, stopDate time.Time, withArchived bool) (err error) {
	if !checkSeed(c.xSid) {
		return fmt.Errorf("авторизация не пройдена")
	}
	reqParams := ReqPlacementsStat{
		PlacementIds: placementIds,
		StartDate:    startDate.Format("2006-01-02"),
		StopDate:     stopDate.Format("2006-01-02"),
	}
	if withArchived {
		reqParams.WithArchived = 1
	} else {
		reqParams.WithArchived = 0
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(reqParams)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, Host+string(PlacementsStat), &buf)
	if err != nil {
		return err
	}
	buildHeader(req, c.xSid)
	if err != nil {
		return err
	}
	resp, err := c.tr.Do(req)
	if err != nil {
		return err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := ""
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetPlacementStatByDay(placementGuid string, startDate time.Time, stopDate time.Time) (stat PlacementsStatByDay, err error) {
	if !checkSeed(c.xSid) {
		return PlacementsStatByDay{}, fmt.Errorf("авторизация не пройдена")
	}
	param := make(map[string]string)
	param["start_date"] = startDate.Format("2006-01-02")
	param["stop_date"] = stopDate.Format("2006-01-02")
	body, err := buildBody(param)
	req, err := http.NewRequest(http.MethodPost, Host+string(PlacementStatByDay)+"/"+placementGuid+"/stat", body)
	if err != nil {
		return PlacementsStatByDay{}, err
	}
	buildHeader(req, c.xSid)
	if err != nil {
		return PlacementsStatByDay{}, err
	}
	resp, err := c.tr.Do(req)
	if err != nil {
		return PlacementsStatByDay{}, err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PlacementsStatByDay{}, err
	}
	data := PlacementsStatByDay{}
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return PlacementsStatByDay{}, err
	}
	return data, nil
}

func buildBody(data map[string]string) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildBody json convert: %v", err)
	}
	return bytes.NewBuffer(b), nil
}

func buildHeader(req *http.Request, xSid string) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-sid", xSid)
}

func checkSeed(xSeed string) (status bool) {
	if xSeed == "" {
		return false
	} else {
		return true
	}
}

package solowaysdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var ErrAuthorize = errors.New("авторизация не пройдена")

type APIError struct {
	message string
}

func (e APIError) Error() string {
	return e.message
}

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

// Login sends login request.
func (c *Client) Login() error {
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

	buildHeader(req, nil)

	resp, err := c.tr.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return APIError{message: resp.Status}
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

	if data.Error != "" {
		return APIError{message: data.Error}
	}

	return nil
}

// ReqUserInfo информация о пользователе.
type ReqUserInfo struct {
	Username string `json:"username"`
	Error    string `json:"error"`
}

// Response Ответ от сервера.
func (c *Client) doRequest(ctx context.Context, method string, url string, body io.Reader) ([]byte, error) {
	if !checkSeed(c.xSid) {
		return nil, ErrAuthorize
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	buildHeader(req, &c.xSid)

	resp, err := c.tr.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, APIError{message: resp.Status}
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// Whoami Получение информации о пользователе.
func (c *Client) Whoami(ctx context.Context) error {
	responseBody, err := c.doRequest(ctx, http.MethodGet, Host+string(Whoami), nil)
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

// GetPlacements Получение списка площадок.
func (c *Client) GetPlacements(ctx context.Context) (PlacementsInfo, error) {
	if c.AccountInfo.Username == "" {
		return PlacementsInfo{}, fmt.Errorf("инфо об аккаунте не получено")
	}

	responseBody, err := c.doRequest(ctx, http.MethodGet, Host+"/api/clients/"+c.AccountInfo.Client.GUID+"/placements", nil)
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

// GetPlacementsStat Получение статистики площадок.
func (c *Client) GetPlacementsStat(ctx context.Context, placementIds []string, startDate time.Time, stopDate time.Time, withArchived bool) error {
	if c.AccountInfo.Username == "" {
		return fmt.Errorf("инфо об аккаунте не получено")
	}

	reqParams := ReqPlacementsStat{
		PlacementIDS: placementIds,
		StartDate:    startDate.Format("2006-01-02"),
		StopDate:     stopDate.Format("2006-01-02"),
	}

	if withArchived {
		reqParams.WithArchived = 1
	} else {
		reqParams.WithArchived = 0
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(reqParams)
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := c.doRequest(ctx, http.MethodPost, Host+string(PlacementsStat), &buf)
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

// GetPlacementStatByDay Получение статистики площадок в разрезе дней.
func (c *Client) GetPlacementStatByDay(ctx context.Context, placementGUID string, startDate time.Time, stopDate time.Time) (PlacementsStatByDay, error) {
	if !checkSeed(c.xSid) {
		return PlacementsStatByDay{}, ErrAuthorize
	}

	param := make(map[string]string)
	param["start_date"] = startDate.Format(time.DateOnly)
	param["stop_date"] = stopDate.Format(time.DateOnly)

	body, err := buildBody(param)
	if err != nil {
		return PlacementsStatByDay{}, fmt.Errorf("cant build body: %w", err)
	}

	responseBody, err := c.doRequest(ctx, http.MethodPost, Host+string(PlacementStatByDay)+"/"+placementGUID+"/stat", body)
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

// построение тела запроса.
func buildBody(data map[string]string) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildBody json convert: %w", err)
	}

	return bytes.NewBuffer(b), nil
}

// построение заголовка запроса.
func buildHeader(req *http.Request, xSid *string) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if xSid != nil {
		req.Header.Add("X-sid", *xSid)
	}
}

// проверка параметра seed.
func checkSeed(xSeed string) bool {
	if xSeed == "" {
		return false
	} else {
		return true
	}
}

package goalteaon

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	transport "github.com/ArthurHlt/basicauth-transport"
	"github.com/ArthurHlt/goalteon/beans"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	DefaultCount = 1024
)

type Client struct {
	Target              string
	httpClient          *http.Client
	locker              sync.Locker
	noAutoApplySaveSync bool
}

func NewClient(target, user, password string, opts ...clientOpt) *Client {
	if !strings.HasPrefix(target, "https://") && !strings.HasPrefix(target, "http://") {
		target = "https://" + target
	}
	target = strings.TrimSuffix(target, "/")
	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           http.DefaultTransport.(*http.Transport).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{},
	}

	httpClient := &http.Client{
		Transport: transport.NewBasicAuthTransport(
			user,
			password,
			tr,
		),
		Timeout: 1 * time.Minute,
	}
	c := &Client{
		Target:     target,
		httpClient: httpClient,
		locker:     &sync.Mutex{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

func (c *Client) NewRequest(method string, bean beans.Bean, urlParams url.Values) (*http.Request, error) {
	params := bean.GetParams()
	var body *bytes.Buffer = nil
	_, isBytesParams := params.(beans.BytesParams)
	if params != nil && !isBytesParams {
		b, err := json.Marshal(bean.GetParams())
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}
	if params != nil && isBytesParams {
		body = bytes.NewBuffer(bean.GetParams().(beans.BytesParams))
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.Target, bean.Path()), body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = urlParams.Encode()
	return req, nil
}

func (c *Client) List(bean beans.Bean, urlParams url.Values) ([]beans.BeanType, error) {
	if len(urlParams) == 0 {
		urlParams = url.Values{}
	}
	if urlParams.Get(UrlParamsKeyCount) == "" {
		urlParams.Set(UrlParamsKeyCount, fmt.Sprintf("%d", DefaultCount))
	}
	if urlParams.Has(UrlParamsKeyProp) {
		urlParams.Set(UrlParamsKeyProps, urlParams.Get(UrlParamsKeyProp))
		urlParams.Del(UrlParamsKeyProp)
	}
	req, err := c.NewRequest(http.MethodGet, bean, urlParams)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	list, err := UnmarshalListResponse(resp, bean)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) Get(bean beans.Bean, urlParams url.Values) (beans.BeanType, error) {
	req, err := c.NewRequest(http.MethodGet, bean, urlParams)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if _, ok := bean.(beans.ScalarBean); !ok {
		foundBeans, err := UnmarshalListResponse(resp, bean)
		if err != nil {
			return nil, err
		}
		if len(foundBeans) == 0 {
			return nil, nil
		}
		return foundBeans[0], nil
	}

	paramType := reflect.TypeOf(bean.GetParams())
	isPtr := false
	if paramType.Kind() == reflect.Ptr {
		isPtr = true
		paramType = paramType.Elem()
	}
	params := reflect.New(paramType).Interface().(beans.BeanType)
	err = UnmarshalResponse(resp, params)
	if err != nil {
		return nil, err
	}
	if !isPtr {
		return reflect.ValueOf(params).Elem().Interface().(beans.BeanType), nil
	}
	return params, nil
}

func (c *Client) Post(bean beans.Bean, urlParams url.Values) (*StatusResponse, error) {
	if bean.Path() == "" {
		return nil, fmt.Errorf("all index must be defined when creating")
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	_, err := c.RevertAll()
	if err != nil {
		return nil, fmt.Errorf("error when reverting all: %s", err.Error())
	}
	req, err := c.NewRequest(http.MethodPost, bean, urlParams)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	sr, err := UnmarshalStatusResponse(resp)
	if err != nil {
		return nil, err
	}
	if !c.noAutoApplySaveSync {
		_, err = c.ApplySaveSync()
		if err != nil {
			return nil, fmt.Errorf("error when applying save sync: %s", err.Error())
		}
	}
	return sr, nil
}

func (c *Client) Create(bean beans.Bean, urlParams url.Values) (*StatusResponse, error) {
	return c.Post(bean, urlParams)
}

func (c *Client) Put(bean beans.Bean, urlParams url.Values) (*StatusResponse, error) {
	if bean.Path() == "" {
		return nil, fmt.Errorf("all index must be defined when updating")
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	_, err := c.RevertAll()
	if err != nil {
		return nil, fmt.Errorf("error when reverting all: %s", err.Error())
	}
	req, err := c.NewRequest(http.MethodPut, bean, urlParams)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	sr, err := UnmarshalStatusResponse(resp)
	if err != nil {
		return nil, err
	}
	if !c.noAutoApplySaveSync {
		_, err = c.ApplySaveSync()
		if err != nil {
			return nil, fmt.Errorf("error when applying save sync: %s", err.Error())
		}
	}
	return sr, nil
}

func (c *Client) Update(bean beans.Bean, urlParams url.Values) (*StatusResponse, error) {
	return c.Put(bean, urlParams)
}

func (c *Client) Delete(bean beans.Bean, urlParams url.Values) (*StatusResponse, error) {
	if bean.Path() == "" {
		return nil, fmt.Errorf("all index must be defined when deleting")
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	_, err := c.RevertAll()
	if err != nil {
		return nil, fmt.Errorf("error when reverting all: %s", err.Error())
	}
	req, err := c.NewRequest(http.MethodDelete, bean, urlParams)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	sr, err := UnmarshalStatusResponse(resp)
	if err != nil {
		return nil, err
	}
	if !c.noAutoApplySaveSync {
		_, err = c.ApplySaveSync()
		if err != nil {
			return nil, fmt.Errorf("error when applying save sync: %s", err.Error())
		}
	}
	return sr, nil
}

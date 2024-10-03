package goalteon

import (
	"fmt"
	"github.com/ArthurHlt/goalteon/beans"
	"net/http"
	"net/url"
)

type BulkMethod uint

const (
	BulkMethodPost BulkMethod = iota
	BulkMethodCreate
	BulkMethodPut
	BulkMethodUpdate
	BulkMethodDelete
)

func (m BulkMethod) ToHttpMethod() string {
	switch m {
	case BulkMethodPost, BulkMethodCreate:
		return http.MethodPost
	case BulkMethodPut, BulkMethodUpdate:
		return http.MethodPut
	case BulkMethodDelete:
		return http.MethodDelete
	}
	return http.MethodPost
}

type BulkItem struct {
	Method       BulkMethod
	Bean         beans.Bean
	UrlParams    url.Values
	IgnoreErrors bool
}

func NewBulkItem(method BulkMethod, bean beans.Bean, urlParams url.Values) *BulkItem {
	return &BulkItem{
		Method:    method,
		Bean:      bean,
		UrlParams: urlParams,
	}
}

func (c *Client) Bulk(items []*BulkItem) ([]*StatusResponse, error) {
	if len(items) == 0 {
		return nil, nil
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	_, err := c.RevertAll()
	if err != nil {
		return nil, fmt.Errorf("error when reverting all: %s", err.Error())
	}
	srs := make([]*StatusResponse, 0)
	for i, item := range items {
		if item.Bean == nil {
			return nil, fmt.Errorf("bean cannot be nil")
		}
		if item.Bean.Path() == "" {
			return nil, fmt.Errorf("all index must be defined when creating")
		}
		req, err := c.NewRequest(item.Method.ToHttpMethod(), item.Bean, item.UrlParams)
		if err != nil {
			return nil, err
		}
		resp, err := c.Do(req)
		if err != nil {
			return nil, err
		}
		sr, err := UnmarshalStatusResponse(resp)
		if err != nil && !item.IgnoreErrors {
			return nil, fmt.Errorf("on item %d: %w", i, err)
		}
		srs = append(srs, sr)
	}
	if !c.noAutoApplySaveSync {
		_, err = c.ApplySaveSync()
		if err != nil {
			return nil, fmt.Errorf("error when applying save sync: %s", err.Error())
		}
	}
	return srs, nil
}

package goalteon

import (
	"encoding/json"
	"fmt"
	"github.com/ArthurHlt/goalteon/beans"
	"io"
	"net/http"
	"reflect"
)

type StatusResponse struct {
	StatusCode int    `json:"-"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func (s *StatusResponse) IsError() bool {
	return s.StatusCode > 299 || s.StatusCode < 200
}

func (s *StatusResponse) String() string {
	return fmt.Sprintf("status (Code %d): %s, message: %s", s.StatusCode, s.Status, s.Message)
}

func (s *StatusResponse) Error() string {
	return s.String()
}

func UnmarshalStatusResponse(resp *http.Response) (*StatusResponse, error) {
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		status, err := UnmarshalStatus(resp.StatusCode, b)
		if err != nil {
			status = &StatusResponse{
				StatusCode: resp.StatusCode,
				Status:     "error",
				Message:    string(b),
			}
		}
		status.StatusCode = resp.StatusCode
		return status, status
	}
	status, err := UnmarshalStatus(resp.StatusCode, b)
	if err != nil {
		return nil, err
	}
	if status.IsError() {
		return status, status
	}
	return status, nil
}

func UnmarshalStatus(statusCode int, data []byte) (*StatusResponse, error) {
	var statusResponse StatusResponse
	err := json.Unmarshal(data, &statusResponse)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal status response (%s), content: %s", err.Error(), string(data))
	}

	if statusResponse.Status == "error" || statusResponse.Status == "err" {
		statusCode = 500
	}
	statusResponse.StatusCode = statusCode
	return &statusResponse, nil
}

func UnmarshalResponse(resp *http.Response, params beans.BeanType) error {
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		status, err := UnmarshalStatus(resp.StatusCode, b)
		if err != nil {
			status = &StatusResponse{
				StatusCode: resp.StatusCode,
				Status:     "error",
				Message:    string(b),
			}
		}
		return status
	}
	if bytesParams, ok := params.(*beans.BytesParams); ok {
		*bytesParams = b
		return nil
	}

	err = json.Unmarshal(b, params)
	if err != nil {
		return fmt.Errorf("could not unmarshal data (%s), content: %s", err.Error(), string(b))
	}
	return nil
}

func UnmarshalListResponse(resp *http.Response, bean beans.Bean) ([]beans.BeanType, error) {
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		status, err := UnmarshalStatus(resp.StatusCode, b)
		if err != nil {
			status = &StatusResponse{
				StatusCode: resp.StatusCode,
				Status:     "error",
				Message:    string(b),
			}
		}
		return nil, status
	}
	mapVal := reflect.New(reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.SliceOf(bean.GetParamsType()))).Type())
	err = json.Unmarshal(b, mapVal.Interface())
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal data (%s), content: %s", err.Error(), string(b))
	}
	var res []beans.BeanType
	valSlice := mapVal.Elem().MapIndex(reflect.ValueOf(bean.Name()))
	for i := 0; i < valSlice.Len(); i++ {
		res = append(res, valSlice.Index(i).Interface().(beans.BeanType))
	}
	return res, nil
}

func Translate[T beans.BeanType](params beans.BeanType) T {
	return params.(T)
}

func TranslateList[T beans.BeanType](params []beans.BeanType) []T {
	if len(params) == 0 {
		return make([]T, 0)
	}
	var res []T
	for _, p := range params {
		res = append(res, p.(T))
	}
	return res
}

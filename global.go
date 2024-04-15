package goalteaon

import (
	"fmt"
	"github.com/ArthurHlt/goalteon/beans"
	"github.com/mitchellh/mapstructure"
	"io"
	"net/http"
	"time"
)

type AgSaveConfig int

const (
	AgSaveConfigSave         AgSaveConfig = 1
	AgSaveConfigIdle         AgSaveConfig = 2
	AgSaveConfigInProgress   AgSaveConfig = 3
	AgSaveConfigComplete     AgSaveConfig = 4
	AgSaveConfigFailed       AgSaveConfig = 5
	AgSaveConfigSaveNoBackup AgSaveConfig = 6
)

type AgApplyConfig int

const (
	AgApplyConfigApply      AgApplyConfig = 1
	AgApplyConfigIdle       AgApplyConfig = 2
	AgApplyConfigInProgress AgApplyConfig = 3
	AgApplyConfigComplete   AgApplyConfig = 4
	AgApplyConfigFailed     AgApplyConfig = 5
)

type AgSyncStatus int

const (
	AgSyncStatusIdle       AgSyncStatus = 0
	AgSyncStatusInProgress AgSyncStatus = 1
	AgSyncStatusSuccess    AgSyncStatus = 2
	AgSyncStatusFailure    AgSyncStatus = 3
)

type StatusGlobal struct {
	AgSaveConfig          AgSaveConfig `mapstructure:"agSaveConfig"`
	AgSaveLastError       error
	AgApplyConfig         AgApplyConfig `mapstructure:"agApplyConfig"`
	AgApplyLastError      error
	AgSyncStatus          AgSyncStatus `mapstructure:"agSyncStatus"`
	AgLastSyncErrorReason string       `mapstructure:"agLastSyncErrorReason"`
	AgLastSyncError       error
}

func (c *Client) StatusGlobal() (*StatusGlobal, error) {
	res, err := c.Get(
		&beans.ConfigBean{},
		UrlParamsProps(nil, "agSaveConfig", "agApplyConfig", "agSyncStatus", "agLastSyncErrorReason"),
	)
	if err != nil {
		return nil, err
	}
	mParams := Translate[beans.MapParams](res)
	status := &StatusGlobal{}
	err = mapstructure.Decode(mParams, status)
	if err != nil {
		return nil, err
	}
	if status.AgLastSyncErrorReason != "" {
		status.AgLastSyncError = fmt.Errorf(status.AgLastSyncErrorReason)
	}

	if status.AgSaveConfig == AgSaveConfigFailed {
		resSave, err := c.List(&beans.AgSaveTable{}, nil)
		if err != nil {
			return nil, err
		}
		resSaveError := TranslateList[*beans.AgSaveTableParams](resSave)
		if len(resSaveError) > 0 {
			status.AgSaveLastError = fmt.Errorf(resSaveError[0].StringVal)
		}
	}

	if status.AgApplyConfig == AgApplyConfigFailed {
		resApply, err := c.List(&beans.AgApplyTable{}, nil)
		if err != nil {
			return nil, err
		}
		resApplyError := TranslateList[*beans.AgApplyTableParams](resApply)
		if len(resApplyError) > 0 {
			status.AgApplyLastError = fmt.Errorf(resApplyError[0].StringVal)
		}
	}
	return status, nil
}

func (c *Client) Save() (*StatusResponse, error) {
	req, err := c.NewRequest(http.MethodPost, &beans.ConfigBean{}, UrlParamsAction(nil, "save"))
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return UnmarshalStatusResponse(resp)
}

func (c *Client) Apply() (*StatusResponse, error) {
	req, err := c.NewRequest(http.MethodPost, &beans.ConfigBean{}, UrlParamsAction(nil, "apply"))
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return UnmarshalStatusResponse(resp)
}

func (c *Client) Sync() (*StatusResponse, error) {
	req, err := c.NewRequest(http.MethodPost, &beans.ConfigBean{}, UrlParamsAction(nil, "sync"))
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return UnmarshalStatusResponse(resp)
}

func (c *Client) Revert() (*StatusResponse, error) {
	req, err := c.NewRequest(http.MethodPost, &beans.ConfigBean{}, UrlParamsAction(nil, "revert"))
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return UnmarshalStatusResponse(resp)
}

func (c *Client) RevertAll() (*StatusResponse, error) {
	req, err := c.NewRequest(http.MethodPost, &beans.ConfigBean{}, UrlParamsAction(nil, "revertapply"))
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return UnmarshalStatusResponse(resp)
}

func (c *Client) ShowDiff() (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/config/getdiff", c.Target), nil)
	if err != nil {
		return "", err
	}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Client) ShowDiffFlash() (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/config/getdiffflash", c.Target), nil)
	if err != nil {
		return "", err
	}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Client) ApplySaveSync() (st *StatusGlobal, err error) {
	_, err = c.Apply()
	if err != nil {
		return nil, fmt.Errorf("error when applying: %s", err.Error())
	}
	for {
		st, err = c.StatusGlobal()
		if err != nil {
			return st, fmt.Errorf("error when getting status: %s", err.Error())
		}
		if st.AgApplyLastError != nil {
			return st, fmt.Errorf("error when applying: %s", st.AgApplyLastError.Error())
		}
		if st.AgApplyConfig != AgApplyConfigInProgress {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	_, err = c.Save()
	if err != nil {
		return nil, fmt.Errorf("error when saving: %s", err.Error())
	}
	for {
		st, err = c.StatusGlobal()
		if err != nil {
			return st, fmt.Errorf("error when getting status: %s", err.Error())
		}
		if st.AgSaveLastError != nil {
			return st, fmt.Errorf("error when saving: %s", st.AgSaveLastError.Error())
		}
		if st.AgSaveConfig != AgSaveConfigInProgress {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	// do not check for errors here, as it is possible that the sync will fail it should not block anything
	_, err = c.Sync()
	if err != nil {
		return nil, fmt.Errorf("error when syncing: %s", err.Error())
	}
	return st, nil
}

func (c *Client) Logout() error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/config/logout", c.Target), nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 && resp.StatusCode < 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("error when logout (status code: %d): %s", resp.StatusCode, string(b))
	}

	return nil
}

package investsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Client struct {
	Host   string
	Client *http.Client
	Logger *zap.SugaredLogger
}

func (c *Client) post(ctx context.Context, path string, body interface{}, result interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return errors.Wrap(err, "json marshal body err")
	}
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/api/v1/%s", c.Host, strings.TrimLeft(path, "/")), bytes.NewReader(bodyBytes))
	req.Header["content-type"] = []string{"application/json"}
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "http post err")
	}
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Logger.Debugw("invest sdk post response", "path", path, "resp", string(respBodyBytes))
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("none ok status: %d", resp.StatusCode)
	}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return errors.Wrap(err, "json unmarshal result failed")
	}
	return nil
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.Client.Get(fmt.Sprintf("%s/api/v1/%s", c.Host, path))
	if err != nil {
		return errors.Wrap(err, "http post err")
	}
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Logger.Debugw("invest sdk post response", "path", path, "resp", string(respBodyBytes))
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("none ok status: %d", resp.StatusCode)
	}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return errors.Wrap(err, "json unmarshal result failed")
	}
	return nil
}

func (c *Client) Health(ctx context.Context) error {
	var ret string
	return c.get("/health", &ret)
}

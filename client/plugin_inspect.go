// +build experimental

package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

// PluginInspectWithRaw inspects an existing plugin
func (cli *Client) PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error) {
	resp, err := cli.get(ctx, "/plugins/"+name, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	defer ensureReaderClosed(resp)
	body, err := ioutil.ReadAll(resp.body)
	if err != nil {
		return nil, nil, err
	}
	var p types.Plugin
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&p)
	return &p, body, err
}

package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/ipfs/go-ipfs/core"
)

const (
	addAPI  = "/api/v0/add"
	statAPI = "/api/v0/files/stat"
)

// IPFS is the struct that contains all informations to access an IPFS node
type IPFS struct {
	client   *http.Client
	node     *core.IpfsNode
	endpoint string
}

// Response return the response of a file uploaded
type Response struct {
	Name string
	Hash string
	Size string
}

// New creates a new IPFS client method
func New() *IPFS {
	return &IPFS{
		client: &http.Client{},
		node: *core.NewNode(context.Background(), &core.BuildCfg{})
		endpoint: "http://core-ipfs:5001",
	}
}

// Add data from a reader to the IPFS node
func (ipfs *IPFS) Add(name string, reader io.Reader) (*Response, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, reader); err != nil {
		return nil, err
	}
	w.Close()

	req, err := http.NewRequest("POST", ipfs.endpoint+addAPI, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := ipfs.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", res.Status)
	}
	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Exists returns true if the file associated of the hash exists
func (ipfs *IPFS) Exists(hash string) (bool, error) {
	response, err := ipfs.client.Get(ipfs.endpoint + statAPI + "?arg" + hash)
	fmt.Println(response)
	if err != nil {
		return false, err
	}
	return response.StatusCode == 200, nil
}

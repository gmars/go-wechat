package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Request struct {
	reqUrl, method string
	body           io.Reader
	headers        *http.Header
}

// NewRequest 创建api请求
func NewRequest(reqUrl, method string) *Request {
	return &Request{
		reqUrl:  reqUrl,
		method:  method,
		headers: &http.Header{},
	}
}

// Do 发送网络请求
func (r *Request) Do(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, r.method, r.reqUrl, r.body)
	if err != nil {
		return nil, err
	}

	if r.headers != nil {
		req.Header = r.headers.Clone()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Printf("request Do body Close error:%s", err.Error())
		}
	}(resp.Body)
	return io.ReadAll(resp.Body)
}

// WithQuery 拼装url参数
func (r *Request) WithQuery(query map[string]string) error {
	if query == nil {
		return nil
	}

	u, err := url.Parse(r.reqUrl)
	if err != nil {
		return err
	}

	q := u.Query()
	for key, val := range query {
		q.Set(key, val)
	}
	u.RawQuery = q.Encode()
	r.reqUrl = u.String()
	return nil
}

// WithFormBody form payload with file
func (r *Request) WithFormBody(params map[string]string, files map[string]*multipart.FileHeader) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	defer func(writer *multipart.Writer) {
		err := writer.Close()
		if err != nil {
			fmt.Printf("request WithFormFile writer Close error:%s", err.Error())
		}
	}(writer)

	for key, val := range params {
		err := writer.WriteField(key, val)
		if err != nil {
			return err
		}
	}

	for key, file := range files {
		err := func(key string, file *multipart.FileHeader) error {
			fh, err := file.Open()
			if err != nil {
				return err
			}
			defer func(fh multipart.File) {
				err = fh.Close()
				if err != nil {
					fmt.Printf("request WithFormFile multipart.File Close error:%s", err.Error())
				}
			}(fh)

			part, err := writer.CreateFormFile(key, file.Filename)
			_, err = io.Copy(part, fh)
			if err != nil {
				return err
			}
			return nil
		}(key, file)

		if err != nil {
			return err
		}
	}

	r.body = payload
	r.headers.Set("Content-Type", writer.FormDataContentType())
	return nil
}

// WithJsonBody json payload
func (r *Request) WithJsonBody(data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r.headers.Set("Content-Type", "application/json")
	r.body = bytes.NewReader(dataBytes)
	return nil
}

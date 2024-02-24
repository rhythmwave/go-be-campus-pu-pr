package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/errors"
	"google.golang.org/grpc/codes"

	log "github.com/sirupsen/logrus"
)

func client() *http.Client {
	return &http.Client{
		// Timeout: time.Second * constants.HttpClientDefaultTimeout,
		Transport: &http.Transport{
			DisableCompression: true,
			DisableKeepAlives:  true,
			// Dial: (&net.Dialer{
			// 	Timeout: 60 * time.Second,
			// }).Dial,
			// TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func do(method string, endpoint string, params map[string]string, headers map[string]string, body io.Reader) ([]byte, *constants.ErrorResponse) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		log.Error(errors.Wrap(err))
		return nil, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}
	req.Close = true

	q := req.URL.Query()
	for i, v := range params {
		q.Add(i, v)
	}
	req.URL.RawQuery = q.Encode()

	for i, v := range headers {
		req.Header.Add(i, v)
	}

	netClient := client()
	resp, err := netClient.Do(req)
	if err != nil {
		log.Error(errors.Wrap(err))
		return nil, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errors.Wrap(err))
		return nil, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}
	if resp.StatusCode != 200 {
		log.Error(err, ConvertBytesToString(buf))
		return nil, ErrHttpClient(endpoint)
	}

	return buf, nil
}

// HttpClientDoMultipart function to hit 3rd party API with content-type multipart/form-data
// Params:
// method: http method
// params: http url parameter request
// headers: http header request
// body: http body request for multpart/form-data
// result: body response, should be struct pointer
// Returns *constants.ErrorResponse
func HttpClientDoMultipart(method string, endpoint string, params map[string]string, headers map[string]string, body map[string]io.Reader, result interface{}) *constants.ErrorResponse {
	var b bytes.Buffer
	var err error

	w := multipart.NewWriter(&b)

	for key, r := range body {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer func() {
				err := x.Close()
				if err != nil {
					log.Errorln(err)
				}
			}()
		}

		if x, ok := r.(*os.File); ok {
			fw, err = w.CreateFormFile(key, x.Name())
			if err != nil {
				log.Error(errors.Wrap(err))
				return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
			}
		} else {
			fw, err = w.CreateFormField(key)
			if err != nil {
				log.Error(errors.Wrap(err))
				return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
			}
		}
		_, err = io.Copy(fw, r)
		if err != nil {
			log.Error(errors.Wrap(err))
			return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
		}
	}

	err = w.Close()
	if err != nil {
		log.Error(errors.Wrap(err))
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	if headers != nil {
		headers["Content-Type"] = w.FormDataContentType()
	} else {
		headers = map[string]string{
			"Content-Type": w.FormDataContentType(),
			"Accept":       "application/json",
		}
	}

	buf, errs := do(method, endpoint, params, headers, &b)
	if errs != nil {
		return errs
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		log.Error(errors.Wrap(err))
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return nil
}

// HttpClientDoJson function to hit 3rd party API with content-type application/json
// Params:
// method: http method
// params: http url parameter request
// headers: http header request
// body: http body request for application/json
// result: body response, should be struct pointer
// Returns *constants.ErrorResponse
func HttpClientDoJson(method string, endpoint string, params map[string]string, headers map[string]string, body interface{}, result interface{}) *constants.ErrorResponse {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Error(errors.Wrap(err))
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	reqBody := bytes.NewBuffer(jsonBody)
	buf, errs := do(method, endpoint, params, headers, reqBody)
	if errs != nil {
		log.Error("HttpClientDoJson 1", errs.Err)
		return errs
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		log.Error("HttpClientDoJson", errors.Wrap(err))
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return nil
}

// HttpClientDo function to hit 3rd party API without body data
// Params:
// method: http method
// params: http url parameter request
// headers: http header request
// result: body response, should be struct pointer
// Returns *constants.ErrorResponse
func HttpClientDo(method string, endpoint string, params map[string]string, headers map[string]string, result interface{}) *constants.ErrorResponse {
	buf, errs := do(method, endpoint, params, headers, nil)
	if errs != nil {
		return errs
	}

	err := json.Unmarshal(buf, result)
	if err != nil {
		log.Error(errors.Wrap(err))
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return nil
}

// HttpClientDoUrlEncoded function to hit 3rd party API with content-type application/x-www-form-urlencoded
// Params:
// method: http method
// params: http url parameter request
// headers: http header request
// body: http body request for application/x-www-form-urlencoded
// result: body response, should be struct pointer
// Returns *constants.ErrorResponse
func HttpClientDoUrlEncoded(method string, endpoint string, params map[string]string, headers map[string]string, body url.Values, result interface{}) *constants.ErrorResponse {
	reqBody := strings.NewReader(body.Encode())
	buf, errs := do(method, endpoint, params, headers, reqBody)
	if errs != nil {
		return errs
	}

	err := json.Unmarshal(buf, &result)
	if err != nil {
		log.Error(errors.Wrap(err))
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return nil
}

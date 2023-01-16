package main

import (
	"bytes"
	"fmt"
	"github.com/phin1x/go-ipp"
	"io"
	"net/http"
	"strconv"
)

func loadPrinters(parseMode string) (int, error) {
	req, err := http.NewRequest(http.MethodPost, cupsEndpoint, bytes.NewReader(cupsRequestBody))
	if err != nil {
		return -1, fmt.Errorf("cannot create request: %v", err)
	}

	req.Header.Set("Content-Type", ipp.ContentTypeIPP)
	req.Header.Set("Content-Length", strconv.Itoa(len(cupsRequestBody)))

	hResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, fmt.Errorf("cannot exec http request: %v", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(hResp.Body)

	var r io.Reader

	if parseMode == directParseMode {
		r = hResp.Body
	} else {
		bytez, err := io.ReadAll(hResp.Body)
		if err != nil {
			return -1, fmt.Errorf("cannot read bytes from CUPS response: %v", err)
		}

		r = bytes.NewReader(bytez)
	}

	var buf bytes.Buffer

	resp, err := ipp.NewResponseDecoder(r).Decode(&buf)
	if err != nil {
		return -1, fmt.Errorf("cannot decode response: %v", err)
	}

	return len(resp.PrinterAttributes), nil
}

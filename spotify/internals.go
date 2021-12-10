package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func getAccessToken() (string, error) {
	reqUrl, _ := url.Parse("https://open.spotify.com/")
	req := &http.Request{
		Method: "GET",
		URL:    reqUrl,
		Header: map[string][]string{
			"User-Agent": {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36 RuxitSynthetic/1.0 v2873991242106302242 t8068951106021062059"},
		},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while retrieving anonymous access token: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	atRes := AccessTokenResponse{}
	if err = json.Unmarshal(resBody, &atRes); err != nil {
		return "", err
	}

	if atRes.AccessToken == "" {
		return "", errors.New("empty access token")
	}
	return atRes.AccessToken, nil
}

func getInternalApiKey() (string, error) {
	reqUrl, _ := url.Parse("https://open.spotify.com/")
	req := &http.Request{
		Method: "GET",
		URL:    reqUrl,
		Header: map[string][]string{
			"User-Agent": {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36 RuxitSynthetic/1.0 v2873991242106302242 t8068951106021062059"},
		},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while retrieving spotify.com: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	matches := regexp.MustCompile(`<script.*>.*"accessToken":"(.+?)",.*<\/script>`).FindSubmatch(resBody)
	if len(matches) <= 1 {
		fmt.Print(string(resBody))
		return "", errors.New("internal api key not found in response body")
	}
	return string(matches[1]), nil
}

package goapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	myClient = &http.Client{}
)

// getJson used only for beauty - reduce the number of lines
// It is assumed that I will use this method more often, than once
func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// GetPublishedArticles used to get a list of articles
//
// "Articles" are all the posts that users create on DEV that typically show up in the feed
// They can be a blog post, a discussion question, a help thread etc. but is referred to as article within the code
//
// By default it will return featured, published articles ordered by descending popularity
//
// Each page will contain 30 articles
//
// Responses, according to the combination of params, are cached for 24 hours
func GetPublishedArticles(q QueryArticle) (response Articles, err error) {
	query := url.Values{}

	if q.Page != 0 {
		query.Add("page", strconv.Itoa(int(q.Page)))
	}

	if q.Tag != "" {
		query.Add("tag", q.Tag)
	}

	if q.Username != "" {
		query.Add("username", q.Username)
	}

	if q.Tag != "" {
		query.Add("tag", q.Tag)
	}

	if q.Top != 0 {
		query.Add("top", strconv.Itoa(int(q.Top)))
	}

	request := fmt.Sprintf("%s%s", PublishedRequest, query.Encode())

	err = getJson(request, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateNewArticle used to create a new article
//
// RATE LIMITING: There is a limit of 10 articles created each 30 seconds by the same user
func CreateNewArticle(p Payload, AccessToken string) (response Article, err error) {
	if AccessToken == "" {
		err = errors.New("AccessToken is empty")
		return Article{}, err
	}

	payloadBytes, err := json.Marshal(p)
	if err != nil {
		return Article{}, err
	}
	body := bytes.NewReader(payloadBytes)

	r, err := http.NewRequest("POST", "https://dev.to/api/articles", body)
	if err != nil {
		return Article{}, err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Api-Key", AccessToken)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return Article{}, err
	}

	if resp.StatusCode != 201 {
		_ = errors.New(fmt.Sprintf("article is not created. Status code: %d", resp.StatusCode))
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Article{}, err
	}

	if response.ID == 0 {
		err = errors.New("the id, that dev.to returned is invalid. Probably, this post is already exists.")
		return response, err
	}

	return response, nil
}

// GetPublishedArticle used for retrieve a single published article given its id
func GetPublishedArticle(id int32) (response Article, err error) {
	idstring := strconv.Itoa(int(id))

	if id == 0 {
		_ = errors.New("invalid id")
	}

	request := fmt.Sprintf("%s%s", "https://dev.to/api/articles/", idstring)

	err = getJson(request, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// UpdateExistingArticle updates the article
func UpdateExistingArticle(p Payload, id int32, AccessToken string) (response Article, err error) {

	// Struct ->  JSON payload
	payloadJSON, err := json.Marshal(p)
	if err != nil {
		return Article{}, err
	}

	body := strings.NewReader(string(payloadJSON))

	// making URL string
	urlWithID := strings.Builder{}
	urlWithID.WriteString("https://dev.to/api/articles/")
	urlWithID.WriteString(strconv.Itoa(int(id)))

	req, err := http.NewRequest("PUT", urlWithID.String(), body)
	if err != nil {
		return Article{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Article{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Article{}, err
	}

	return response, nil
}

// Error:  parsing time """" as ""2006-01-02T15:04:05Z07:00"": cannot parse """ as "2006"
// TODO: make custom unmarshal
//func GetPublishedArticlesByMe(page int32, perPage int32, AccessToken string) (response Articles, err error) {
//
//	if perPage > 1000 {
//		err = errors.New("maximum page size is 1000")
//		return nil, err
//	}
//
//	req, err := http.NewRequest("GET", "https://dev.to/api/articles/me/all", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Api-Key", AccessToken)
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//
//	err = json.NewDecoder(resp.Body).Decode(&response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}

package goapi

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var AccessToken = os.Getenv("AccessToken")

func TestGetPublishedArticles(t *testing.T) {
	response, err := GetPublishedArticles(QueryArticle{
		Page:     0,
		Tag:      "",
		Username: "",
		State:    "",
		Top:      0,
	})

	if err != nil {
		t.Error(err)
	}

	if response[0].TypeOf != "article" {
		t.Errorf("GetPublishedArticles returns not an article in first position. [0].TypeOf = %s", response[0].TypeOf)
	}
}

func TestGetPublishedArticle(t *testing.T) {
	response, err := GetPublishedArticle(666)

	if err != nil {
		t.Error(err)
	}

	if response.ID != 666 {
		t.Errorf("Invalid ID. Expected %d, response %d.", 666, response.ID)
	}

	if response.TypeOf != "article" {
		t.Errorf("GetPublishedArticle returns not an article in first position. [0].TypeOf = %s", response.TypeOf)
	}
}

func TestCreateNewArticle(t *testing.T) {
	// making pseudorandom more random
	rand.Seed(time.Now().UnixNano())
	// random int with uint = MAX(uint32)
	// because dev.to uses uint32 as ID
	rndint := rand.Intn(int(^uint32(0)))

	// Random title
	title := strings.Builder{}
	title.WriteString("Random title with ID = ")
	title.WriteString(strconv.Itoa(rndint))

	response, err := CreateNewArticle(Payload{Article: NewArticle{
		Title:          title.String(),
		Description:    "Testing description API",
		BodyMarkdown:   "#BodyMarkdown",
		Published:      false,
		Series:         "TestingGoAPI",
		MainImage:      "",
		CanonicalURL:   "",
		Tags:           []string{"go", "api"},
		OrganizationID: 0,
	}}, AccessToken)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("--- URL of the post: %s\n", response.URL)
}

//func TestUpdateExistingArticle(t *testing.T) {
//
//}

//func TestGetPublishedArticlesByMe(t *testing.T) {
//	_, err := GetPublishedArticlesByMe(1, 30,  AccessToken); if err != nil {
//		t.Error(err)
//	}
//
//}

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/manoj-JU/web-page-scrapper/utils"
)

// Scrape amazon's amazonProduct page
func ScrapeAmazonProductPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var url URL
	_ = json.NewDecoder(r.Body).Decode(&url)
	collector := colly.NewCollector()

	amazonProduct := ScrapeAmazonProductPageHelper(collector)

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit(url.URL)

	if amazonProduct.Title == "" {
		utils.HandleError(errors.New("Invalid webPage"), http.StatusBadRequest, w)
		return
	}

	err := CallDBPersistAPI(amazonProduct)
	if err != nil {
		utils.HandleError(err, http.StatusInternalServerError, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&amazonProduct)

}

//calls db_persistor_service api to persist amazonProduct
func CallDBPersistAPI(amazonProduct *AmazonProduct) error {
	fmt.Println("inside callDBPersistAPI")
	postBody, _ := json.Marshal(amazonProduct)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://db_persistor_service:8080/amazon/products", "application/json", responseBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &amazonProduct)
	if err != nil {
		return err
	}

	mp := make(map[string]interface{})
	err = json.Unmarshal(body, &mp)
	if err != nil {
		return err
	}

	if val, ok := mp["CreatedAt"]; ok {
		amazonProduct.CreatedAt = val.(string)
	}
	if val, ok := mp["UpdatedAt"]; ok {
		amazonProduct.UpdatedAt = val.(string)
	}
	return nil
}

//scrapes the amazon product details web page
func ScrapeAmazonProductPageHelper(collector *colly.Collector) *AmazonProduct {
	amazonProduct := AmazonProduct{}

	//get title
	collector.OnHTML("#productTitle", func(element *colly.HTMLElement) {
		title := strings.TrimSpace(element.Text)
		amazonProduct.Title = title
	})

	//get imageURL
	collector.OnHTML("#imgTagWrapperId img", func(element *colly.HTMLElement) {
		imageURL := strings.TrimSpace(element.Attr("src"))
		amazonProduct.ImageURL = imageURL
	})

	amazonProduct.Description = ""
	//get Description
	collector.OnHTML("#feature-bullets ul li:not(#replacementPartsFitmentBullet)", func(element *colly.HTMLElement) {
		amazonProduct.Description += strings.TrimSpace(element.Text)
	})

	//get price
	collector.OnHTML("#priceblock_ourprice", func(element *colly.HTMLElement) {
		amazonProduct.Price += strings.TrimSpace(element.Text)
	})

	//get total reviews
	collector.OnHTML("#acrCustomerReviewText", func(element *colly.HTMLElement) {
		re := regexp.MustCompile("[0-9]+")
		s := re.FindAllString(element.Text, -1)
		totalreviews := ""
		for _, str := range s {
			totalreviews += str
		}
		amazonProduct.TotalReviews = totalreviews
	})

	return &amazonProduct
}

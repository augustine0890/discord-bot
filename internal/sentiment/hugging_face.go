package sentiment

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TextRequest struct {
	Text string `json:"text"`
}

type Results struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

func HuggingFaceSentiment(request TextRequest, task string) (*Results, error) {
	c := http.Client{Timeout: time.Duration(5) * time.Minute}
	var results *Results

	textRequestJson, err := json.Marshal(request)
	if err != nil {
		log.Printf("Text Request Marshal: %s", err)
		return nil, err
	}

	endpoint := "http://localhost:8000/" + task
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(textRequestJson))
	if err != nil {
		log.Printf("HuggingFace Sentiment Endpoint: %s", err)
		return nil, err
	}

	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("Error http: %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error Reading Body: %s", err)
		return nil, err
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Printf("Error Unmarshaling: %s", err)
		return nil, err
	}
	return results, nil
}

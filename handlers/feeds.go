package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/vineethtw/vishu/requests"
	"github.com/vineethtw/vishu/services"
)

/*
Create is used to create a new card
*/
func Create(feedsService services.FeedService) http.Handler {
	feedRequest := requests.NewFeedRequest{}
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		if err := json.Unmarshal(body, &feedRequest); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		if feedRequest.Payload == "" {
			http.Error(writer, errors.New("payload cannot be empty").Error(), http.StatusBadRequest)
			return
		}

		feedsService.CreateNew("invoice", feedRequest.Payload)
		writer.WriteHeader(http.StatusOK)
	})
}

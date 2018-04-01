package services

/*
FeedService deals with all operations involving Feeds
*/
type FeedService interface {
	CreateNew(eventType string, payload string)
}

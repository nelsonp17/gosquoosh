package appwrite

import (
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/storage"
)

// New creates a new appwrite client and storage service
func New(endpoint, projectId, apiSecret string) (client.Client, *storage.Storage) {

	var _endpoint, _apiSecret string
	if endpoint == "" && apiSecret == "" {
		_endpoint = Endpoint
		_apiSecret = ApiSecret
	} else {
		_endpoint = endpoint
		_apiSecret = apiSecret
	}

	_client := appwrite.NewClient(
		appwrite.WithEndpoint(_endpoint),
		appwrite.WithProject(projectId),
		appwrite.WithKey(_apiSecret),
	)
	service := storage.New(_client)
	return _client, service
}

func NewInstanceService(endpoint, projectId, apiSecret string) *storage.Storage {
	var service *storage.Storage
	_, service = New(endpoint, projectId, apiSecret)

	return service
}

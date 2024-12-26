package appwrite

import (
	"fmt"
)

var ApiSecret string
var Endpoint string

func GetPathFileFull(endpoint, projectId, bucketId, fileId string) string {
	if endpoint == "" {
		endpoint = Endpoint
	}
	if projectId == "" {
		projectId = "project"
	}
	return fmt.Sprintf("%s/storage/buckets/%s/files/%s/view?project=%s", endpoint, bucketId, fileId, projectId)
}

func Set(appwriteEndpoint, appwriteApiSecret string) {
	//fmt.Println("Set Appwrite")
	//fmt.Println("Endpoint", appwriteEndpoint)
	//fmt.Println("ApiSecret", appwriteApiSecret)

	Endpoint = appwriteEndpoint
	ApiSecret = appwriteApiSecret
}

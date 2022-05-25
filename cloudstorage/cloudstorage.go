package cloudstorage

import (
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

// var (
// 	storageClient *storage.Client
// )

//HandleFileUploadToBucket uploads file to bucket
func HandleFileUploadToBucket(c *gin.Context) {
	bucket := "test_golang_api"

	var err error

	ctx := appengine.NewContext(c.Request)

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("key/test-golang-sa.json"))
	if err != nil {
		internalServerErrorHandler(c, err)
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		internalServerErrorHandler(c, err)
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		internalServerErrorHandler(c, err)
	}

	if err := sw.Close(); err != nil {
		internalServerErrorHandler(c, err)
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	if err != nil {
		internalServerErrorHandler(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file " + sw.Attrs().Name + " uploaded succesfully",
		"pathname": u.EscapedPath(),
	})

}

func internalServerErrorHandler(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
	return
}

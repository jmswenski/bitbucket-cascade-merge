package main

import (
	"bitbucket-cascade-merge/internal"
	"github.com/gin-gonic/gin"
	"github.com/ktrysmt/go-bitbucket"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	username := os.Getenv("BITBUCKET_USERNAME")
	password := os.Getenv("BITBUCKET_PASSWORD")
	releaseBranchPrefix := os.Getenv("RELEASE_BRANCH_PREFIX")
	developmentBranchName := os.Getenv("DEVELOPMENT_BRANCH_NAME")
	bitbucketSharedKey := os.Getenv("BITBUCKET_SHARED_KEY")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	if username == "" {
		log.Fatal("$BITBUCKET_USERNAME must be set. See README.md")
	}
	if password == "" {
		log.Fatal("$BITBUCKET_PASSWORD must be set. See README.md")
	}
	if releaseBranchPrefix == "" {
		log.Fatal("RELEASE_BRANCH_PREFIX must be set. See README.md")
	}
	if developmentBranchName == "" {
		log.Fatal("DEVELOPMENT_BRANCH_NAME must be set. See README.md")
	}
	if bitbucketSharedKey == "" {
		log.Fatal("BITBUCKET_SHARED_KEY must be set. See README.md")
	}

	bitbucketClient := bitbucket.NewBasicAuth(username, password)
	bitbucketService := internal.NewBitbucketService(bitbucketClient, releaseBranchPrefix, developmentBranchName)
	bitbucketController := internal.NewBitbucketController(bitbucketService, bitbucketSharedKey)

	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/", bitbucketController.Webhook)
	router.GET("/", func(c *gin.Context){
		c.JSON(200, nil)
	})

	_ = router.Run(":" + port)
}




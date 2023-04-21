// Package main is the entrypoint for azstorageping.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const usage = "usage: azstorageping [account] [key] [container]"

// pingAccount connects to a storage account, listing container contents.
func pingAccount(account string, key string, container string) error {

	kc, err := azblob.NewSharedKeyCredential(account, key)
	if err != nil {
		return fmt.Errorf("azstorageping.pingAccount: error creating credential key: %w", err)

	}

	url := fmt.Sprintf("https://%s.blob.core.windows.net", account)
	client, err := azblob.NewClientWithSharedKeyCredential(url, kc, nil)
	if err != nil {
		return fmt.Errorf("azstorageping.pingAccount: error creating client: %w", err)
	}

	log.Printf("azstorageping.pingAccount: connected to %s\n", client.URL())
	log.Printf("azstorageping.pingAccount: checking access to container %s\n", container)

	pager := client.NewListBlobsFlatPager(container, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	log.Println("azstorageping.pingAccount: listing blobs")

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			return fmt.Errorf("azstorageping.pingAccount? error listing container contents: %w", err)
		}

		for _, blob := range resp.Segment.BlobItems {
			log.Println(*blob.Name)
		}
	}

	return nil

}

func main() {

	if len(os.Args) < 4 {
		log.Fatalf("azstorageping.main: Invaild arguments\n%s", usage)
	}

	account := os.Args[1]
	key := os.Args[2]
	container := os.Args[3]

	err := pingAccount(account, key, container)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("azstorageping.main: ping complete")
}

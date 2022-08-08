package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

/**
Function to read from Google Cloud Storage. Generic implementation for Demo purposes:

- read public cloud storage with no authentication
- Currently using NEXRAD publicly accessible bucket & object as its freely available

Once we have a designated bucket & associated object, We can add specific parse logic.
**/

func readFromStorage() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		log.Fatal(err)
	}
	rc, err := client.Bucket("gcp-public-data-nexrad-l2/1991/06/05/KTLX").Object("NWS_NEXRAD_NXL2LG_KTLX_19910605160000_19910605235959.tar").NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	slurp, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file contents:", slurp)
}

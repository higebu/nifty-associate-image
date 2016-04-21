package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/higebu/go-niftycloud/compute"
	"github.com/higebu/go-niftycloud/niftycloud"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [OPTIONS] [IMAGE ID]\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	public := flag.Bool("public", false, "distribute the image to the public")
	redistribute := flag.Bool("redistribute", false, "enable to redistribute the image")
	distributionIds := flag.String("distribution-id", "", "commma-separated distribution ids")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	imageId := args[0]

	var ids []string
	if *distributionIds != "" {
		ids = strings.Split(*distributionIds, ",")
	}

	auth, err := niftycloud.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := compute.New(auth, niftycloud.JPEast)
	opts := compute.NiftyAssociateImage{
		ImageId:         imageId, // Test image.
		IsPublic:        *public,
		IsRedistribute:  *redistribute,
		DistributionIds: ids,
	}
	resp, err := client.NiftyAssociateImage(&opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Return)
}

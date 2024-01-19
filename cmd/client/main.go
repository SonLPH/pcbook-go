package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SonLPH/pcbook-go/client"
	"github.com/SonLPH/pcbook-go/pb"
	"github.com/SonLPH/pcbook-go/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testCreateLaptop(laptopClient *client.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
}

func testSearchLaptop(laptopClient *client.LaptopClient) {
	for i := 0; i < 10; i++ {
		laptopClient.CreateLaptop(sample.NewLaptop())
	}
	filter := &pb.Filter{
		MaxPriceUsd:        3000,
		MinCpuCores:        4,
		MinCpuFrequencyGhz: 2.5,
		MinRam:             &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}

	laptopClient.SearchLaptop(filter)
}

func testUploadImage(laptopClient *client.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	laptopClient.UploadImage(laptop.GetId(), "tmp/laptop.jpg")
}

func testRateLapop(laptopClient *client.LaptopClient) {
	n := 3
	laptopIDs := make([]string, n)

	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopClient.CreateLaptop(laptop)
		laptopIDs[i] = laptop.GetId()
	}

	scores := make([]float64, n)

	for {
		fmt.Print("rate laptop (y/n)?")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal("cannot rate laptop: ", err)
		}
	}
}

const (
	username         = "user1"
	password         = "secret"
	refreshDuaration = 30 * time.Second
)

func authMethods() map[string]bool {
	const laptopServicePath = "/pcbook.LaptopService/"

	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("Dial server: %s", *serverAddress)

	cc1, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClinet := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClinet, authMethods(), refreshDuaration)

	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	cc2, err := grpc.Dial(
		*serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)

	laptopClient := client.NewLaptopClient(cc2)
	testRateLapop(laptopClient)
}

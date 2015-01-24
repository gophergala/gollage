package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"strconv"
	"strings"
)

var accessKey, secretKey string
var auth aws.Auth

func init() {
	key, err := ioutil.ReadFile("accessKey")
	if err != nil {
		fmt.Println("Lol have fun uploading images without an AWS key", err)
	}
	accessKey = strings.TrimSpace(string(key))

	key, err = ioutil.ReadFile("secretKey")
	if err != nil {
		fmt.Println("Lol have fun uploading images without an AWS key", err)
	}
	secretKey = strings.TrimSpace(string(key))

	auth = aws.Auth{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func AddImageToBucket(wall Wall, wallName, imageName string, imageData io.Reader, length int64) {
	connection := s3.New(auth, aws.USEast)
	picBucket := connection.Bucket("gollage/" + wallName + "/rawImages")
	name := strings.Split(imageName, ".")[0]
	name = name + strconv.Itoa(len(wall.Images)) + ".png"
	picBucket.PutReader(name, imageData, length, "image/png", s3.PublicRead)
}

func NewWallBucket(name string) error {
	connection := s3.New(auth, aws.USEast)
	picBucket := connection.Bucket("gollage/" + name)
	return picBucket.PutBucket(s3.PublicRead)
}

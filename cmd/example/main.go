package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhashkevych/auth/config"
	"github.com/zhashkevych/auth/pkg/parser"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	usernamePtr := flag.String("username", "test_user6", "username")
	passwordPtr := flag.String("pass", "qwerty", "password")

	flag.Parse()

	if err := createUser(*usernamePtr, *passwordPtr); err != nil {
		log.Fatal(err)
	}

	token, err := authorize(*usernamePtr, *passwordPtr)
	if err != nil {
		log.Fatal(err)
	}

	user, err := parser.ParseToken(token, []byte(viper.GetString("auth.signing_key")))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created and authorized user: %+v", user)
}

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func createUser(username, password string) error {
	reqBody := &userRequest{
		Username: username,
		Password: password,
	}

	resp := new(response)
	return request(reqBody, resp, "http://localhost:8001/auth/sign-up")
}

func authorize(username, password string) (string, error) {
	reqBody := &userRequest{
		Username: username,
		Password: password,
	}

	resp := new(response)
	if err := request(reqBody, resp, "http://localhost:8001/auth/sign-in"); err != nil {
		log.Fatal(err)
	}

	return resp.Token, nil
}

func request(req *userRequest, res *response, endpoint string) error {
	reqBodyBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		endpoint,
		"application/json",
		bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, res); err != nil {
		return err
	}

	if res.Status == "error" {
		return errors.New(fmt.Sprintf("error occured on user creation: %s", res.Message))
	}

	return nil
}

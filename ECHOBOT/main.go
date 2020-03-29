package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)	

//Input point
func main() {
	botToken := "1108473075:AAEL-qG3bFybXIcaYsARalDW20C_ig4v_S8"
	//https://api.telegram.org/bot<token>/METHOD_NAME
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	for ;   ; {
		updates,err := getUpdates(botUrl) {
			if err != nil {
				log.Println("Smth went wrong: ", err.Error())
			}
			fmt.Println(updates)
		}
	}
}

//request to update
func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result,nil
}

//response to updates
// func respond() {

// }

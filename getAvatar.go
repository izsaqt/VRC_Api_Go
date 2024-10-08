package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getAvatar(avatarID string) {
	authToken, err := loadAuthToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	getAvatarUrl := "https://vrchat.com/api/1/avatars/" + avatarID
	fmt.Println(getAvatarUrl)

	req, err := http.NewRequest("GET", getAvatarUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "auth="+authToken)
	req.Header.Add("User-Agent", "golang Client")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to send get avatar Status Code:", resp.StatusCode)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	formattedData := formatJSON(respBody)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response body:", formattedData)

}

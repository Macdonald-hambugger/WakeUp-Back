package service

import (
	"WakeUp-Back/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetFriends(c *gin.Context) {
	accessToken := entity.AccessToken
	apiUrl := "https://kapi.kakao.com/v1/api/talk/friends"

	data := url.Values{}
	data.Set("limit", "100")
	data.Set("order", "asc")
	data.Set("friend_order", "nickname")
	data.Set("offset", "0")

	req, err := http.NewRequest("GET", apiUrl+"?"+data.Encode(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// JSON 파일로 저장
	err = ioutil.WriteFile("response.json", body, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Response saved to response.json"})
}

func LoadFriends(c *gin.Context) {
	jsonf, err := ioutil.ReadFile("response.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var jsonResult interface{}
	err = json.Unmarshal(jsonf, &jsonResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jsonResult)
}

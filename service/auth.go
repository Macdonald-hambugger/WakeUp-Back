package service

import (
	"WakeUp-Back/entity"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

func requestUserInfo() {
	resp, err := http.Get("https://kapi.kakao.com/v2/user/me")
	if err != nil {
		fmt.Printf("failed to request user information: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v\n", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	fmt.Println(result)
}

func GetKakaoUserInfo(accessToken string) (*entity.User, error) {
	req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo entity.User
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func SignUp(db *gorm.DB, c *gin.Context) {
	accessToken := c.Param("id")
	userInfo, err := GetKakaoUserInfo(accessToken)
	if err != nil {
		panic("카카오 유저 정보를 받아오는데 실패하였습니다.")
	}

	user := &entity.User{
		Nickname: userInfo.Nickname,
		Email:    userInfo.Email,
		Profile:  userInfo.Profile,
		Uuid:     userInfo.Uuid,
	}

	if err := db.Create(user).Error; err != nil {
		panic(err)
		return
	}

	c.JSON(200, user)
}

func Login(db *gorm.DB, c *gin.Context) {
	accessToken := c.Param("id")
	userInfo, err := GetKakaoUserInfo(accessToken)
	if err != nil {
		panic("카카오 유저 정보를 받아오는데 실패하였습니다.")
	}

	user := &entity.User{
		Nickname: userInfo.Nickname,
		Email:    userInfo.Email,
		Profile:  userInfo.Profile,
		Uuid:     userInfo.Uuid,
	}

	if err := db.Create(user).Error; err != nil {
		panic(err)
		return
	}

	if err := db.Where("uuid = ?", user.Uuid).First(&user).Error; err != nil {
		err := db.Create(&user).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	} else {
		c.SetCookie("user_uuid", user.Uuid, 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{"message": "User logged in successfully"})
	}
}

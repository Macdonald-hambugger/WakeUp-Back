package service

import (
	"WakeUp-Back/entity"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateToken(userId int64) (string, error) {
	var err error

	secretKey := "asdooinvzxcubuwebdcs" // 이 키는 보안을 위해 안전하게 보관해야 합니다.

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 토큰 유효 시간은 24시간으로 설정했습니다. 필요에 따라 변경하세요.
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetKakaoUserInfo(accessToken string) (*entity.KakaoUserInfoDTO, error) {
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

	var userInfo entity.KakaoUserInfoDTO
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info from Kakao"})
		return
	}

	c.JSON(200, gin.H{"userInfo": userInfo})
}

func Login(db *gorm.DB, c *gin.Context) {
	accessToken := c.Param("id")

	userInfo, err := GetKakaoUserInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info from Kakao"})
		return
	}

	user := &entity.User{
		ID:           userInfo.ID,
		Nickname:     userInfo.KakaoAccount.Profile.Nickname,
		ProfileImage: userInfo.KakaoAccount.Profile.ProfileImageURL,
	}

	if err := db.Where("id = ?", user.ID).First(&user).Error; err != nil {
		err := db.Create(&user).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	} else {
		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

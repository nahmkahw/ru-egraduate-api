package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
)

var ctx = context.Background()

func GoogleAuth(c *gin.Context) {

	const BEARER_SCHEMA = "Bearer "
	AUTH_HEADER := c.GetHeader("Authorization")
	ID_TOKEN := AUTH_HEADER[len(BEARER_SCHEMA):]

	if len(AUTH_HEADER) == 0 {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "authorization key in header not found"})
		c.Abort()
		return
	}

	googleAuthVerify, err := verifyGoogleAuth(ID_TOKEN)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "Google is not authorized"})
		c.Abort()
		return
	}

	log.Println("LOG จาก Google authorized ไป generate token ต่อ...", googleAuthVerify)
	c.Next()

}

func verifyGoogleAuth(id_token string) (*oauth2.Tokeninfo, error) {

	// ตั้งเวลาการรอคอย หรือกำหนดการหยุดเชื่อมต่อกรณีที่ไม่ได้ Respose จาก Google
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout}

	oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(id_token)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	trip "github.com/nvsces/flw-server-go"
	"github.com/sirupsen/logrus"
)


func (h *Handler) login(c *gin.Context) {
	var input trip.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	access_token:=os.Getenv("access_token");

	client_secret:=os.Getenv("client_secret");

	urlReq:="https://api.vk.com/method/secure.checkToken?v=5.81&token="+input.AccessToken+"&client_secret="+client_secret+"&access_token="+access_token;

	resp, err := http.Post(urlReq, "application/json", nil)

    if err != nil {
        logrus.Print(err)
    }

	body, err := ioutil.ReadAll(resp.Body)

	var info Info
	if err := json.Unmarshal(body, &info); err != nil {
	}
		if input.UserVkId!=info.Key.UserId{
			newErrorResponse(c, http.StatusInternalServerError, "Верефикация ВК не пройдена")
			return
	}

	// проверка на нового пользователя

	id, err := h.services.Authorization.GetUser(input.UserVkId);
	fmt.Println(id)
	fmt.Println(err)
	if err != nil {
		fmt.Println("User not found")
		_, err2 := h.services.Authorization.CreateUser(input)
		if err2 != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	token, err := h.services.Authorization.GenerateToken(input.UserVkId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"id":id,
	})
}


type Client struct {
	Date     int   `json:"date"`
	Expire   int   `json:"expire"`
	Success  int   `json:"success"`
	UserId   int   `json:"user_id"`
}

type Info struct {
	Key Client `json:"response"`
}
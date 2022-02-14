package handlers

import (
	"github.com/pusher/pusher-http-go"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (repo *DBRepo) PusherAuth(w http.ResponseWriter, r *http.Request) {
	// user 會在 app.Session.Put中放入 Context中
	userId := repo.App.Session.GetInt(r.Context(), "userID")
	u, _ := repo.DB.GetUserById(userId)
	//params, _ := ioutil.ReadAll(r.Body)  // go 1.16之後不建議 ioutil
	params, _ := io.ReadAll(r.Body)

	presenceData := pusher.MemberData{
		UserID: strconv.Itoa(userId),
		UserInfo: map[string]string{
			"name": u.FirstName,
			"id":   strconv.Itoa(userId),
		},
	}

	response, err := app.WsClient.AuthenticatePresenceChannel(params, presenceData)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println("PusherAuth 寫出... ")
	_, _ = w.Write(response)
}

//func (repo *DBRepo) TestPusher(w http.ResponseWriter, r *http.Request) {
//	data := make(map[string]string)
//	data["message"] = "Hello, world"
//
//	err := repo.App.WsClient.Trigger("public-channel", "test-event", data)
//	if err != nil {
//		log.Println(err)
//	}
//}

package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"encoding/json"
	"strings"

	"dluweb/api"
)

// {
//  "Id":"2312577",
//  "FirstName":null,
//  "LastName":null,
//  "FullName":"Trần Nguyễn Tuấn Anh",
//  "Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjIzMTI1NzciLCJOYW1lIjoiVHLhuqduIE5ndXnhu4VuIFR14bqlbiBBbmgiLCJSb2xlIjoiU1YiLCJuYmYiOjE3MDY1MzEwMDgsImV4cCI6MTcwNjUzODIwOCwiaWF0IjoxNzA2NTMxMDA4LCJpc3MiOiJQU0NVSVNBcGkiLCJhdWQiOiJ2aHUifQ.VYtpXP5gviw1uIBKFJFmpl905LGEViLtP7LYfAbFKwE",
//  "Role":"SV",
//  "GraduateLevel":null,
//  "IsLogin":true,
//  "Message":"",
//  "Expire":"2024-01-29T21:23:28.2465334+07:00"
// }

type LoginForm struct {
	Id	string	`json:"Id"`
	FirstName	string	`json:"FirstName"`
	LastName	string	`json'"LastName"`
	FullName	string	`json:"FullName"`
	Token	string	`json:"Token"`
	Role	string	`json:"Role"`
	GraduateLevel	string	`json:"GraduateLevel"`
	IsLogin	bool	`json:"IsLogin"`
	Message	string	`json:"Message"`
	Expire	string	`json:"Expire"`
	Avatar	string	`json:"data"`
}

var (
	Response	LoginForm
	// Avatar	AvatarBase64
)

func Login(w http.ResponseWriter, r *http.Request) {
	if Response.IsLogin {
		fmt.Println("Đã đăng nhập, chuyển hướng về /")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		// Lấy dữ liệu từ form
		username := r.FormValue("email")
		password := r.FormValue("password")

		// Gọi API để đăng nhập và nhận JSON response
		jsonData := api.ApiLogin(username, password)

		// Parse JSON response vào biến LoginForm
		err := json.Unmarshal([]byte(jsonData), &Response)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
			return
		}

		jsonAvatar := api.ApiAvatar(Response.Token)
		err = json.Unmarshal([]byte(jsonAvatar), &Response)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
			return
		}

		base64head := "data:jpg;base64,"
		fullBase64 := strings.Replace(Response.Avatar, base64head, "", 1)
		Response.Avatar = fullBase64

		// Kiểm tra trạng thái đăng nhập
		if Response.IsLogin {
			fmt.Println("Đăng nhập thành công !")
			if len(Response.Avatar) != 0 {
				fmt.Println("Đã lấy avatar thành công !")
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Nếu đăng nhập không thành công, in ra thông báo lỗi
		fmt.Println("Tài khoản hoặc mật khẩu không đúng!")
	}

	// Nếu là GET request hoặc sau khi xử lý POST, hiển thị trang login
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "login.html"),
		path.Join("views", "pages.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

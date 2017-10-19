package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	//"strconv"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/golab.com.htm")
	if err != nil {
		fmt.Println("templateParseFile:", err)
		fmt.Fprintf(w, "Server error") //ここでwに書き込まれたものがクライアントに出力されます。
	} else {
		t.Execute(w, nil)
	}
	//
	r.ParseForm() //urlが渡すオプションを解析します。POSTに対してはレスポンスパケットのボディを解析します（request body）
	//注意：もしParseFormメソッドがコールされなければ、以下でフォームのデータを取得することができません。
	fmt.Println(r.Form) //これらのデータはサーバのプリント情報に出力されます
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//	fmt.Fprintf(w, "Hello astaxie!") //ここでwに書き込まれたものがクライアントに出力されます。
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/login.gtpl")
		if err != nil {
			fmt.Println("templateParseFile:", err)
			fmt.Fprintf(w, "Server error") //ここでwに書き込まれたものがクライアントに出力されます。
		} else {
			t.Execute(w, nil)
		}
		fmt.Println("[path]", r.URL.Path)
	} else {
		r.ParseForm()       //urlが渡すオプションを解析します。POSTに対してはレスポンスパケットのボディを解析します（request body）
		fmt.Println(r.Form) //これらのデータはサーバのプリント情報に出力されます
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
			fmt.Println("val2:", v)
		}
		//ログインデータがリクエストされ、ログインのロジック判断が実行されます。
		fmt.Println("username:", r.Form["UserName"])
		fmt.Println("password:", r.Form["Password"])
		if len(r.Form["UserName"][0]) < 6 || len(r.Form["Password"][0]) < 6 {
			t, _ := template.ParseFiles("template/login.gtpl")
			t.Execute(w, nil)
		} else {
			t, _ := template.ParseFiles("template/userinfo.html")
			t.Execute(w, nil)
		}
	}
}

func uinfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/userinfo.html")
		if err != nil {
			fmt.Println("templateParseFile:", err)
			fmt.Fprintf(w, "Server error") //ここでwに書き込まれたものがクライアントに出力されます。
		} else {
			t.Execute(w, nil)
		}
		fmt.Println("[path]", r.URL.Path)
	} else {
		r.ParseForm()       //urlが渡すオプションを解析します。POSTに対してはレスポンスパケットのボディを解析します（request body）
		fmt.Println(r.Form) //これらのデータはサーバのプリント情報に出力されます
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
		//
		// 입력데이타 검사
		//
		// 이름
		if m, _ := regexp.MatchString("^[가-힣]+$", r.Form.Get("hangulname")); !m {
			fmt.Println("한글이름Error : ", r.Form.Get("hangulname"))
			t, _ := template.ParseFiles("template/userinfo.html")
			t.Execute(w, nil)
			return
		}
		// name
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
			fmt.Println("영어이름Error : ", r.Form.Get("engname"))
			t, _ := template.ParseFiles("template/userinfo.html")
			t.Execute(w, nil)
			return
		}
		//전화번호
		//telno, err := strconv.Atoi(r.Form.Get("phone"))
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("phone")); !m {
			//if err != nil {
			//数の変換でエラーが発生。つまり、数字ではありません
			fmt.Println("전화번호Error : ", r.Form.Get("phone"))
			t, _ := template.ParseFiles("template/userinfo.html")
			t.Execute(w, nil)
			return
		}
		//휴대폰번호
		if m, _ := regexp.MatchString(`^(010[0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			//数の変換でエラーが発生。つまり、数字ではありません
			fmt.Println("휴대폰번호Error : ", r.Form.Get("mobile"))
			t, _ := template.ParseFiles("template/userinfo.html")
			t.Execute(w, nil)
			return
		}
		//e-mail
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("e-mail type Error : ", r.Form.Get("email"))
			t, _ := template.ParseFiles("template/userinfo.html")
			t.Execute(w, nil)
			return
		}
		// Pull-Down Menu
		var b bool
		slice := []string{"bykatalk", "bymail", "bymessage"}
		for _, v := range slice {
			if v == r.Form.Get("osirase") {
				b = true
			}
		}
		if !b {
			fmt.Println("Pull-Down menu Error : ", r.Form.Get("email"))
			t, _ := template.ParseFiles("template/userinfo.html")
			t.Execute(w, nil)
			return
		}
		// 체크박스 검증
		/*
			slice = []string{"economic", "science", "sports", "entertainment"}
			a := Slice_diff(r.Form["interest"], slice)
			if a != nil {
				fmt.Println("Check-Box menu Error : ", r.Form.Get("email"))
				t, _ := template.ParseFiles("template/userinfo.html")
				t.Execute(w, nil)
				return
			}
		*/
		//
		fmt.Println("- 이상 -")

	}
}

//
func main() {
	http.HandleFunc("/", sayhelloName)       //アクセスのルーティングを設定します
	http.HandleFunc("/login", login)         //アクセスのルーティングを設定します
	http.HandleFunc("/userinfo", uinfo)      //アクセスのルーティングを設定します
	err := http.ListenAndServe(":8080", nil) //監視するポートを設定します
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

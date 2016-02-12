package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Status struct has the variables holding the status and status description
type Status struct {
	StatusDesc1 string
	StatusDesc2 string
	StatusDesc3 string
	Status1     string
	Status2     string
	Status3     string
}

var (
	serverStatus Status
	lightStatus  string
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error in parsing the form.")
	}
	username := template.HTMLEscapeString(r.Form.Get("username"))
	password := template.HTMLEscapeString(r.Form.Get("password"))
	if username == "admin" && password == "admin" {
		fmt.Fprintln(w, "Success")
	} else if username == "admin" {
		fmt.Fprintln(w, "Wrong password")
	} else if password == "admin" {
		fmt.Fprintln(w, "Wrong username")
	} else {
		fmt.Fprintln(w, "Wrong credentials")
	}
}

func handleSetLightStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error in parsing the form.")
	}
	lightStatus = template.HTMLEscapeString(r.Form.Get("status"))
	fmt.Fprintln(w, "Light is switched "+lightStatus+".")
}

func handleGetLightStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Light is switched "+lightStatus+".")
}

func handleGetServerStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error in parsing the form.")
	}
	data, err := json.MarshalIndent(serverStatus, "", " ")
	if err != nil {
		log.Fatal("Error in parsing json response at handleServerStatus", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(data))
}

func handleDemoPageView(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("demo").ParseFiles("demopage.html"))
	t.Execute(w, serverStatus)
}

func handleSetServerStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error in parsing the form.")
	}
	serverStatus.StatusDesc1 = template.HTMLEscapeString(r.Form.Get("StatusDesc1"))
	serverStatus.StatusDesc2 = template.HTMLEscapeString(r.Form.Get("StatusDesc2"))
	serverStatus.StatusDesc3 = template.HTMLEscapeString(r.Form.Get("StatusDesc3"))
	serverStatus.Status1 = template.HTMLEscapeString(r.Form.Get("Status1"))
	serverStatus.Status2 = template.HTMLEscapeString(r.Form.Get("Status2"))
	serverStatus.Status3 = template.HTMLEscapeString(r.Form.Get("Status3"))

	t := template.Must(template.New("demo").ParseFiles("demopage.html"))
	t.Execute(w, serverStatus)
}

func main() {
	serverStatus.StatusDesc1 = "Power consumption"
	serverStatus.StatusDesc2 = "Total power remaining"
	serverStatus.StatusDesc3 = "Estimated time left"
	serverStatus.Status1 = "0 KWh"
	serverStatus.Status2 = "0 KWh"
	serverStatus.Status3 = "00:00"
	lightStatus = "Off"

	http.HandleFunc("/HomeAutomation/Login", handleLogin)
	http.HandleFunc("/HomeAutomation/SetLightStatus", handleSetLightStatus)
	http.HandleFunc("/HomeAutomation/GetLightStatus", handleGetLightStatus)
	http.HandleFunc("/HomeAutomation/GetServerStatus", handleGetServerStatus)
	http.HandleFunc("/HomeAutomation/", handleDemoPageView)
	http.HandleFunc("/HomeAutomation/SetServerStatus", handleSetServerStatus)

	fmt.Println("Starting the server at port 9090")
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

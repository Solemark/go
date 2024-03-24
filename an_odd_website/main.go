package main

import (
	"an_odd_website/router"
	clientRouter "an_odd_website/router/clients"
	employeeRouter "an_odd_website/router/employees"
	settingRouter "an_odd_website/router/settings"
	"log"
	"net/http"
)

func setSettingRoutes() {
	http.HandleFunc("/data/settings", settingRouter.GetSettingData)
	http.HandleFunc("/data/settings/update", settingRouter.UpdateSetting)
}

//func setAccountingRoutes() {}

func setEmployeeRoutes() {
	http.HandleFunc("/data/employees", employeeRouter.GetEmployeeData)
	http.HandleFunc("/data/employees/new", employeeRouter.NewEmployee)
	http.HandleFunc("/data/employees/update", employeeRouter.UpdateEmployee)
	http.HandleFunc("/data/employees/remove", employeeRouter.RemoveEmployee)
}

func setClientRoutes() {
	http.HandleFunc("/data/clients", clientRouter.GetClientData)
	http.HandleFunc("/data/clients/new", clientRouter.NewClient)
	http.HandleFunc("/data/clients/update", clientRouter.UpdateClient)
	http.HandleFunc("/data/clients/remove", clientRouter.RemoveClient)
}

func setRoutes() {
	http.HandleFunc("/", router.GetWebPage)
	http.HandleFunc("/styles/{style}", router.GetStyles)
	http.HandleFunc("/scripts/{script}", router.GetScripts)
	http.HandleFunc("/favicon.ico", router.GetIcon)
	setClientRoutes()
	setEmployeeRoutes()
	//setAccountingRoutes()
	setSettingRoutes()
}

func main() {
	setRoutes()
	log.Println("Starting server at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Yega-Noragami/go-chaos/pkg/models"
	"github.com/Yega-Noragami/go-chaos/pkg/utils"
)

var NewList models.List

func GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /lists | %s \n", time.Now().String())
	newList := models.GetAllList()
	res, _ := json.Marshal(newList)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST /lists | %s \n", time.Now().String())
	list := &models.List{}
	utils.ParseBody(r, list)
	l := list.CreateList()
	res, _ := json.Marshal(l)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetListByKey(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("GET /lists/{listkey} | %s \n", time.Now().String())
// 	vars := mux.Vars(r)
// 	listKey := vars["Tempkey"]
// 	listDetails, _ := models.GetListByKey(listKey)

// 	res, _ := json.Marshal(listDetails)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteList(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	listKey := vars["Tempkey"]
// 	list := models.DeleteList(listKey)
// 	res, _ := json.Marshal(list)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateList(w http.ResponseWriter, r *http.Request) {
// 	var updateList = &models.List{}
// 	utils.ParseBody(r, updateList)
// 	vars := mux.Vars(r)
// 	listKey := vars["Tempkey"]

// 	//check is item exists in list
// 	listDetails, db := models.GetListByKey(listKey)
// 	if updateList.TempKey != "" {
// 		listDetails.TempKey = updateList.TempKey
// 	}
// 	if updateList.TempValue != "" {
// 		listDetails.TempValue = updateList.TempValue
// 	}
// 	// if updateList.Timestamp != "" {
// 	// 	listDetails.Timestamp = updateList.Timestamp
// 	// }
// 	db.Save(&listDetails)
// 	res, _ := json.Marshal(listDetails)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

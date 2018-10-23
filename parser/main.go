package main

//func parse(w http.ResponseWriter, r *http.Request) {  // single handler
//	var js []byte
//	var status int
//
//	switch r.Method {
//	case http.MethodGet:
//		js, _ = json.Marshal(map[string]string{"status": "ok"})
//		status = http.StatusOK
//	default:
//		js, _ = json.Marshal(map[string]string{"status": "error"})
//		status = http.StatusBadRequest
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	w.Write(js)
//}
//
//func main() {
//	http.HandleFunc("/parse", parse)
//	err :=  http.ListenAndServe(":8000", nil)
//
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

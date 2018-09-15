package main
import ("fmt"
		"encoding/json"
		"log"
		"net/http"
		_ "math/rand"
		_ "strconv"
	)
	
// http Basic auth method
// https://golang.org/pkg/net/http/#Request.BasicAuth
func basicAuth(w http.ResponseWriter, r *http.Request) {
        
    w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

    username, password, authOK := r.BasicAuth()
    if authOK == false {
        http.Error(w, "Not authorized", 401)
        return
    }

    if username != "appointy"  || password != "codestudio" {
        http.Error(w, "Not authorized", 401)
        return
    }
    
}

// Todolists and Todos as slice
//var todolists []Todolists
var todos []Todos

// model
type Todos struct{
    operationId string   `json:"operationId"`
    responses int        `json:"responses"`       
    parameter *parameter `json:"parameter"`
    tags string          `json:"tags"` 
}

type parameter struct{
    name string `json:"name"`
    in string `json:"in"`
    required string `json:"required"`
}
	
func addTodoList(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "hello")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}	

func deleteTodoList(w http.ResponseWriter, r *http.Request){
    
}

func editTodoListName(w http.ResponseWriter, r *http.Request){

}

func addTodoItem(w http.ResponseWriter, r *http.Request){

}

func deleteTodoListItem(w http.ResponseWriter, r *http.Request){

}

func getTodoListItem(w http.ResponseWriter, r *http.Request){

}

func updateTodoItem(w http.ResponseWriter, r *http.Request){

}
	
func main() {
	
	//Route handler. Endpoints
	http.HandleFunc("/todolist", addTodoList)
	http.HandleFunc("/todolist/{id}", deleteTodoList)
	//panic
	http.HandleFunc("/todolist/{id}--", editTodoListName)
	http.HandleFunc("/todolist:additem", addTodoItem)
	http.HandleFunc("/todolist:deleteitem/{id}", deleteTodoListItem)
	http.HandleFunc("/todolist:getItem/{id}", getTodoListItem)
	http.HandleFunc("/todolist/updateItem", updateTodoItem)
	
	http.HandleFunc("/", basicAuth)
	
	//Serving port
	log.Fatal(http.ListenAndServe(":8000", nil))
}

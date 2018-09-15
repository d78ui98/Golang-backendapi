package main

import ("github.com/lib/pq")

func connect(){
    opts := &bg.Options{
        Users: "postgresql",
        Password: "postgresql"
        Addr: "localhost:8000"        
    }
    var db *pg.DB := pg.Connect(opts)
    
    if db == nil {
        log.Printf("Failed to connect to database.\n")
        os.exit(1)
    }
    log.Printf("Connection to database failed")

    closeErr := db.Close()
    if closeErr != nil {
        log.Printf("Err")
    }
}

func main(){
    db.Connect()
}


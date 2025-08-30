package main

import "encoding/json"

func main() {
	bytes, err := json.Marshal(cm1)

	//desr
	cm2:=Coloe{}
	json.Unmarshal(bytes,cm2)
}

type Color struct {
	Id   uint `json:"id"`
	Code string
}

//insert and get operation of usern(id,name,email,status)

//when u insert a record it should be logged into a file lo infor(date and time,action("insert/get"))


// the log should be goroutine and also channel

// its hsould be file based insert,id should be be randome numbers


for true{
	1.insert 
	enter name
	enter email
//succesfully saved
	2.get

	Enter email fetch the data searching in a file
}

//struct
//wriruing a file
//loop by users enter

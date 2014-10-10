package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type jsonStructure struct {
	IntParam1, IntParam2       int
	StringParam1, StringParam2 string
}

type CandidateStructure struct {
	Name, Dob, Address, City, State, Degree, Gender, Occupation string
	Age                                                         int
}

func main() {

	rtr := mux.NewRouter()
	rtr.Headers("Content-Type", "application/json",
		"X-Requested-With", "XMLHttpRequest")
	rtr.Methods("GET,POST")
	rtr.HandleFunc("/hello", HelloWorld)
	rtr.HandleFunc("/helloJson", JsonGetAndSet)
	rtr.HandleFunc("/helloDb", DbConnector)
	rtr.HandleFunc("/StarPyramidGenerator/{TotalRows}", StarPyramidGenerator)
	rtr.HandleFunc("/CreateCandidate", CreateCandidate)
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello senthil!")
}

func JsonGetAndSet(w http.ResponseWriter, req *http.Request) {
	jsonRequestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData := jsonStructure{}
	err = json.Unmarshal(jsonRequestBody, &jsonData)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(jsonData)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(jsonResponse)
	}

}

func DbConnector(w http.ResponseWriter, req *http.Request) {
	conn, err := connectDb()
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	// Create table if not exists
	CreateTableIfNotExists(conn)

	jsonRequestBody, err1 := ioutil.ReadAll(req.Body)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData := jsonStructure{}
	err = json.Unmarshal(jsonRequestBody, &jsonData)
	if err != nil {
		fmt.Println(err.Error())
	}
	//Prepare insert statement...
	ins, err1 := conn.Prepare("insert into mytable (integer1,integer2,string1,string2) values (?,?,?,?)")
	if err1 != nil {
		fmt.Println(err.Error())
	}

	//Begining a new transaction...
	tr, err1 := conn.Begin()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	tr_ins := tr.Do(ins)

	//Performing inserts...
	_, err = tr_ins.Run(jsonData.IntParam1, jsonData.IntParam2, jsonData.StringParam1, jsonData.StringParam2)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Commit the transaction...
	tr.Commit()

	conn.Close()
	fmt.Fprintln(w, "Inserted Successfully")
}

func StarPyramidGenerator(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	totalrows := vars["TotalRows"]
	//	totalrows := req.FormValue("TotalRows")
	rows, err := strconv.ParseInt(totalrows, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	var i, space, k int64
	k = 0
	for i = 1; i <= rows; i++ {
		for space = 1; space <= rows-i; space++ {
			fmt.Fprint(w, " ")
		}
		for k != 2*i-1 {
			fmt.Fprint(w, "* ")
			k++
		}
		k = 0
		fmt.Fprint(w, "\n")
	}
	return
}

func CreateCandidate(w http.ResponseWriter, req *http.Request) {
	jsonRequestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData := CandidateStructure{}
	err = json.Unmarshal(jsonRequestBody, &jsonData)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Body.Close()

	conn, err := connectDb()
	if err != nil {
		fmt.Println(err.Error())
	}
	CreateCandidateTableIfNotExists(conn)
	//Prepare insert statement...
	ins, err1 := conn.Prepare("insert into candidate (address,age,city,degree,dob,gender,name,occupation,state) values (?,?,?,?,?,?,?,?,?)")
	if err1 != nil {
		fmt.Println(err.Error())
	}

	//Begining a new transaction...
	tr, err1 := conn.Begin()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	tr_ins := tr.Do(ins)

	//Performing inserts...
	_, err = tr_ins.Run(jsonData.Address, jsonData.Age, jsonData.City, jsonData.Degree, jsonData.Dob, jsonData.Gender, jsonData.Name, jsonData.Occupation, jsonData.State)
	if err != nil {
		fmt.Println(w, err.Error())
	}

	//Commit the transaction...
	tr.Commit()

	conn.Close()
	fmt.Fprintln(w, "Inserted Successfully")

}
func connectDb() (mysql.Conn, error) {
	conn := mysql.New("tcp", "", "localhost:3306", "root", "data", "test")
	err := conn.Connect()
	if err != nil {
		return nil, err
	}
	return conn, nil

}

func CreateTableIfNotExists(conn mysql.Conn) {

	_, err := conn.Start("create table  IF NOT EXISTS mytable(`id` int(11) NOT NULL AUTO_INCREMENT, integer1 int(11), integer2 int(11),string1 text,string2 text, PRIMARY KEY (`id`))")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreateCandidateTableIfNotExists(conn mysql.Conn) {
	_, err := conn.Start("create table  IF NOT EXISTS candidate(`id` int(11) NOT NULL AUTO_INCREMENT, address varchar(100), age int(11), city varchar(50),  degree varchar(50),dob datetime, gender varchar(10),name varchar(50),  occupation varchar(50),  state varchar(50), PRIMARY KEY (id))")
	if err != nil {
		fmt.Println(err.Error())
	}
}

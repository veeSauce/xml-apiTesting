package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Bookstore struct {
	XMLName xml.Name `xml:"bookstore"`
	Text    string   `xml:",chardata"`
	Book    []struct {
		Text     string `xml:",chardata"`
		Category string `xml:"category,attr"`
		Title    struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"title"`
		Author string `xml:"author"`
		Year   string `xml:"year"`
		Price  string `xml:"price"`
	} `xml:"book"`
}


func main(){

	router := mux.NewRouter()

	router.HandleFunc("/doSomething", DoSomething).Methods("GET")

	fmt.Printf("running server")

	log.Fatal(http.ListenAndServe(":8080", router))


	//bookstore := Bookstore{};
	//
	//byteXml, err := xml.Marshal(bookstore)
	//if err != nil{
	//	fmt.Errorf(err.Error())
	//}
	//
	//client := &http.Client{}
	//
	//url := "http://localhost:8080/doSomething"
	//
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteXml))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	//// now POST it
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(resp)


}

func DoSomething(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	for _, item := range params {
		fmt.Printf("these is the request sent throught the post request: " + item)
	}

	json.NewEncoder(w).Encode("This is the response")

}

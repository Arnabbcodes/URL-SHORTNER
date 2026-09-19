package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)
type URL struct{
	ID string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL string `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}
var urlDB=make(map[string]URL);
func GenerateShortURL(OriginalURL string)string{
	hasher:=md5.New();
	hasher.Write([]byte(OriginalURL));//convert url to bytes of slice
    fmt.Println("Hasher:",hasher);
	data:=hasher.Sum(nil);
	fmt.Println("Hasher data:",data);
	hash:=hex.EncodeToString(data);
	fmt.Println("Encoded hash:",hash);
	fmt.Println("Final String:",hash[:8]);
	return hash[:8];
}
func CreateURL(originalURL string)string{
	shortURL:=GenerateShortURL(originalURL);
	id:=shortURL;
	urlDB[id]=URL{
		ID:id,
        OriginalURL: originalURL,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}
	return shortURL;
}
//IF A USER ENTERS THE SHORT URL AND WANTS TO GET THE ORG URL--
func GetURL(id string)(URL,error){
	url, ok:=urlDB[id];
	if !ok{
		return URL{},errors.New("OOPS! URL not found.")
	}
	return url,nil;
}
func RootPageURL(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello World!");
}
//shortening url
func ShortURLHandler(w http.ResponseWriter,r *http.Request){
	var data struct{
		URL string `json:"url"`
	}
	error:=json.NewDecoder(r.Body).Decode(&data);
	if error!=nil{
		http.Error(w,"Invalid request",http.StatusBadRequest)
		return
	}
	shortURL_:=CreateURL(data.URL);
	//fmt.Fprintf(w,shortURL);
	response:=struct{
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}
	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(response);
	
}
func RedirectURLHandler(w http.ResponseWriter,r *http.Request){
	id:=r.URL.Path[len("/redirect/"):];
	url, err:=GetURL(id);
	if err!=nil{
		http.Error(w,"Invalid request",http.StatusNotFound)
	}
	http.Redirect(w,r,url.OriginalURL,http.StatusFound);
    
}
func main(){
	fmt.Println("Shortening the URL..");
	OriginalURL:="https://jsonplaceholder.typicode.com/todos";
	GenerateShortURL(OriginalURL);
	// // //handler function call
	http.HandleFunc("/",RootPageURL);
	http.HandleFunc("/shorten",ShortURLHandler);
	http.HandleFunc("/redirect/",RedirectURLHandler);
	//forming server
	fmt.Println("Starting your server on port 5000......");
	err:=http.ListenAndServe(":5000",nil);
	if err!=nil{
		fmt.Println("Error while creating your server",err);
	}
}

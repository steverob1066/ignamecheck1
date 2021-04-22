package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main(){
	file, ferr := os.Open("usernames.txt")
	if ferr != nil{
		log.Fatal(ferr)
	}
	// Create a new buffer to read the file into
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		usernameText := scanner.Text()
		username := strings.ToLower(usernameText)
		wg.Add(1)
		go checkName(username)
		wg.Wait()
	}
}

func checkName(name string){
	checkName, err := http.Get("https://www.instagram.com/"+name+"/")
	if err != nil{
		log.Fatal(err)
	}
	checkPageBytes, err := ioutil.ReadAll(checkName.Body)
	checkPageStr := string(checkPageBytes)
	if strings.Contains(checkPageStr, "@"+name+""){
		fmt.Println(""+name+" Is Taken")
	} else{
		fmt.Println(""+name+" Is Free")
	}
	wg.Done()
}

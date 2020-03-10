package main

import ("fmt"
		"net/http")

func main(){

	client := &http.Client {}

		url := "https://github.com/arlon021"
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Errorf("deu ruim")
		}

		response, err := client.Do(Request)
		if err != nil{
			fmt.Errorf("deu ruim2")
		}
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil{
			fmt.Errorf("deu ruim2")
		}
		var Usuario User
		err := json.Unmarshal(bytes, &usuario)
		if err != nil{
			fmt.Errorf("deu ruim3")
		}


	
}
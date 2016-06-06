package main

import "fmt"
import "os"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "github.com/mscrypto/mscryptotest3/data_types"


const baseURL string = "api.softlayer.com/rest/v3/"

func main() {
	username := os.Getenv("SL_USERNAME")
	apiKey := os.Getenv("SL_API_KEY")
	loadBalancerService := "SoftLayer_Account"
	loadBalancerFunction := "getAdcLoadBalancers.json"
	request := "https://"+username+":"+apiKey+"@"+baseURL+loadBalancerService+"/"+loadBalancerFunction
	resp, err := http.Get(request)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	stringBody := string(body)

	fmt.Println(stringBody)
	var resLoadBalancers []data_types.ResLoadBalancer
	json.Unmarshal(body, &resLoadBalancers)
	res,_ := json.Marshal(resLoadBalancers[0])
	fmt.Println(string(res))
}

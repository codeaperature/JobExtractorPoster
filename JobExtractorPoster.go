package main

import (
	"fmt"
	"net/http"

//	"github.com/shirou/gopsutil/mem"
//	"github.com/shirou/gopsutil/cpu"
	"bytes"
	"io/ioutil"
//	"net/url"
	"time"
	"encoding/json"
)


type requestNodeType struct {
	Urls 		[]string `json:"urls"`
}


var req requestNodeType

func init() {

	req.Urls = []string{
		"http://www.indeed.com/viewjob?jk=9388f66529358f6a",
		"http://www.indeed.com/viewjob?jk=53e937ef73c0c808",
		"http://www.indeed.com/viewjob?jk=65f38db1e507757b",
		"http://www.indeed.com/viewjob?jk=096a4e930ada10a8",
		"http://www.indeed.com/viewjob?jk=df685f1132a82dbc",
		"http://www.indeed.com/viewjob?jk=16e9e11bd0c0c29a",
		"http://www.indeed.com/viewjob?jk=9920f716603991f9",
		"http://www.indeed.com/viewjob?jk=f7f7998a6c11bd33",
		"http://www.indeed.com/viewjob?jk=42413713c332b92a",
		"http://www.indeed.com/viewjob?jk=5e1b910a2748eb4e",
		"http://www.indeed.com/viewjob?jk=2e73d41c79dfafd1",
		"http://www.indeed.com/viewjob?jk=069786d5d0d1a1d8",
		"http://www.indeed.com/viewjob?jk=df58f4394fbe3d48",
		"http://www.indeed.com/viewjob?jk=3175c76cfeca92dc",
		"http://www.indeed.com/viewjob?jk=7ac008854406191e",
		"http://www.indeed.com/viewjob?jk=8ac4d4be02dfe265",
		"http://www.indeed.com/viewjob?jk=92756ecfcbbeefb3",
		"http://www.indeed.com/viewjob?jk=e2484787b59d6ae2",
		"http://www.indeed.com/viewjob?jk=d60f61b65192ab8d",
		"http://www.indeed.com/viewjob?jk=e7fb235841447ec7",
		"http://www.indeed.com/viewjob?jk=34f1014de0e57be8",
		"http://www.indeed.com/viewjob?jk=8a1fecc87b15a8dc",
		"http://www.indeed.com/viewjob?jk=0538a626b9c676ee",
		"http://www.indeed.com/viewjob?jk=f366ff05d9ed17b3",
		"http://www.indeed.com/viewjob?jk=b892a5a9e9607cd8",
		"http://www.indeed.com/viewjob?jk=146f98c644a4f431",
		"http://www.indeed.com/viewjob?jk=a079ced2ba220bb5",
		"http://www.indeed.com/viewjob?jk=147c3d6c1c252cc3",
		"http://www.indeed.com/viewjob?jk=2bbe433d8b9c4451",
		"http://www.indeed.com/viewjob?jk=a05fbfa8c33ba1b4",
		"http://www.indeed.com/viewjob?jk=04e4aff184ce968a",
		"http://www.indeed.com/viewjob?jk=37747988137bef9e",
		"http://www.indeed.com/viewjob?jk=32dea5a856ca5f81",
		"http://www.indeed.com/viewjob?jk=500f03947e63e32c",
		"http://www.indeed.com/viewjob?jk=64b0c049d1419b78",
		"http://www.indeed.com/viewjob?jk=257474c0b2d8f0d3",
		"http://www.indeed.com/viewjob?jk=8cdd20b6e3016e41",
		"http://www.indeed.com/viewjob?jk=8d3afae0c14e141d",
		"http://www.indeed.com/viewjob?jk=1dc1399281fa659d",
		"http://www.indeed.com/viewjob?jk=6a9e355c9810e92d",
		"http://www.indeed.com/viewjob?jk=923c58da3d9d189f",
		"http://www.indeed.com/viewjob?jk=fff77f0ffa29f0d5",
		"http://www.indeed.com/viewjob?jk=9312e02e691b58d2",
		"http://www.indeed.com/viewjob?jk=4a7e006151aeda01",
		"http://www.indeed.com/viewjob?jk=975ef15d5eb01cb5",
		"http://www.indeed.com/viewjob?jk=c53db9d9a5e36981",
		"http://www.indeed.com/viewjob?jk=2ab1cd544a31bd65",
		"http://www.indeed.com/viewjob?jk=872c204c1c4ac087",
		"http://www.indeed.com/viewjob?jk=2871144dcff52445",
		"http://www.indeed.com/viewjob?jk=b5f135a0ea266ee3",
		"http://www.indeed.com/viewjob?jk=8aeff41fe5c08440",
		"http://www.indeed.com/viewjob?jk=31cbd932ea1dfbe7",
		"http://www.indeed.com/viewjob?jk=47a925be180132d5",
		"http://www.indeed.com/viewjob?jk=a466e20697a0696d",
		"http://www.indeed.com/viewjob?jk=240a7aef6b40ca73",
		"http://www.indeed.com/viewjob?jk=57cc284f7d989690",
		"http://www.indeed.com/viewjob?jk=428d970e5ef528dc",
		"http://www.indeed.com/viewjob?jk=abf087cfdb95229b",
		"http://www.indeed.com/viewjob?jk=76972d49a8e4678a",
		"http://www.indeed.com/viewjob?jk=debaabfd1324cab7",
		"http://www.indeed.com/viewjob?jk=75902fbad462a726",
		"http://www.indeed.com/viewjob?jk=8f564470bced0079",
		"http://www.indeed.com/viewjob?jk=6d5266667bd59e4c",
		"http://www.indeed.com/viewjob?jk=7ffd55af12b10bbd",
		"http://www.indeed.com/viewjob?jk=a20f2faa0c1205c7",
		"http://www.indeed.com/viewjob?jk=47e9258f2f2a31c0",
		"http://www.indeed.com/viewjob?jk=56b518ddb1e96408",
		"http://www.indeed.com/viewjob?jk=31311ffecee2e57e",
		"http://www.indeed.com/viewjob?jk=5cde34b5e32d1c13",
		"http://www.indeed.com/viewjob?jk=d0b98235c5d1b0d3",
		"http://www.indeed.com/viewjob?jk=2a92bc35e8d20800",
		"http://www.indeed.com/viewjob?jk=c8efbe17f865ea3c",
		"http://www.indeed.com/viewjob?jk=37b3fe15c163d732",
		"http://www.indeed.com/viewjob?jk=e0404df4ec114768",
		"http://www.indeed.com/viewjob?jk=0b91b9be2f82353a",
		"http://www.indeed.com/viewjob?jk=b74b533fc63d3ca5",
		"http://www.indeed.com/viewjob?jk=aab69f9b8f6974b8",
		"http://www.indeed.com/viewjob?jk=f9c6d2189ca38e5a",
		"http://www.indeed.com/viewjob?jk=62e7e4c2352db3b0",
		"http://www.indeed.com/viewjob?jk=6a1cdbc6104307c9",
		"http://www.indeed.com/viewjob?jk=a79fcb55870a3d63",
		"http://www.indeed.com/viewjob?jk=5e26b99b60b90cdb",
		"http://www.indeed.com/viewjob?jk=8b370760e8c792bd",
		"http://www.indeed.com/viewjob?jk=5f3947f2cd4987ce",
		"http://www.indeed.com/viewjob?jk=8b370760e8c792bd",
		"http://www.indeed.com/viewjob?jk=c6657ac06c0c613f",
		"http://www.indeed.com/viewjob?jk=60436ab165786c2d",
		"http://www.indeed.com/viewjob?jk=5796062b217b5ee5",
		"http://www.indeed.com/viewjob?jk=006be650f92b58b7",
		"http://www.indeed.com/viewjob?jk=1f938d51463b681a",
		"http://www.indeed.com/viewjob?jk=5bc8cd072f8f720a",
		"http://www.indeed.com/viewjob?jk=677da2257ac4de24",
		"http://www.indeed.com/viewjob?jk=2bb631c1c859c646",
		"http://www.indeed.com/viewjob?jk=7c8ebfca47d59121",
		"http://www.indeed.com/viewjob?jk=ea79eff9b23893d6",
		"http://www.indeed.com/viewjob?jk=7b56a44f77817173",
		"http://www.indeed.com/viewjob?jk=1f59e45b958363d6",
		"http://www.indeed.com/viewjob?jk=95e615115f1d00ff",
		"http://www.indeed.com/viewjob?jk=42e87a5537816dba",
		"http://www.indeed.com/viewjob?jk=b6d1c01f8adac186",
		"http://www.indeed.com/viewjob?jk=75c5135e31233d35",
		"http://www.indeed.com/viewjob?jk=f59e7d45a3e06d24",
		"http://www.indeed.com/viewjob?jk=5f5425ad1085dcde",
		"http://www.indeed.com/viewjob?jk=2d9e176a4b27ca80",
		"http://www.indeed.com/viewjob?jk=0ec5c209217714a3",
		"http://www.indeed.com/viewjob?jk=77ba306c4179c682",
		"http://www.indeed.com/viewjob?jk=2af3c51e806f4ec0",
		"http://www.indeed.com/viewjob?jk=1607de2893f23ef0",
		"http://www.indeed.com/viewjob?jk=0acfed0e5762bb47",
		"http://www.indeed.com/viewjob?jk=b6c1a2f3158c8633",
		"http://www.indeed.com/viewjob?jk=d6740029f6e7e436",
		"http://www.indeed.com/viewjob?jk=305bc9029b2dead2",
		"http://www.indeed.com/viewjob?jk=d016eac3b074254f",
		"http://www.indeed.com/viewjob?jk=29e330615071fe0d",
		"http://www.indeed.com/viewjob?jk=66b66c5fbb5627c5",
		"http://www.indeed.com/viewjob?jk=068811486851426d",
		"http://www.indeed.com/viewjob?jk=96c3f2f0eb0c1682",
		"http://www.indeed.com/viewjob?jk=4b67e0a5234d1816",
		"http://www.indeed.com/viewjob?jk=efc2f258b5c24de5",
		"http://www.indeed.com/viewjob?jk=cd9ee776538b50e0",
		"http://www.indeed.com/viewjob?jk=1ccc744266956f63",
		"http://www.indeed.com/viewjob?jk=aa7037034dc10ff7",
		"http://www.indeed.com/viewjob?jk=55e83f1578e2fb8c",
		"http://www.indeed.com/viewjob?jk=6059a1e30a600d37",
		"http://www.indeed.com/viewjob?jk=5756008a96bba24f",
		"http://www.indeed.com/viewjob?jk=8c7405c42436afb0",
		"http://www.indeed.com/viewjob?jk=04b35e7720a430bc",
		"http://www.indeed.com/viewjob?jk=194b4d2b83c86068",
		"http://www.indeed.com/viewjob?jk=262a999f99c9e458",
		"http://www.indeed.com/viewjob?jk=4bfacadadab37f82",
		"http://www.indeed.com/viewjob?jk=f2dfa857dc57d15f",
	}
}



func main() {

	fmt.Println("Start Collecting Agent!\n")

	

	for ;; {
	
	    url := "http://localhost:8888/get_jobs"
	    fmt.Println("URL:>", url)
	
	    jsonStr,err := json.Marshal(req)
	    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	    req.Header.Set("X-Custom-Header", "myvalue")
	    req.Header.Set("Content-Type", "application/json")
	
	    client := &http.Client{}
	    resp, err := client.Do(req)
	    if err != nil {
	        panic(err)
	    }
	    defer resp.Body.Close()
	
	    fmt.Println("response Status:", resp.Status)
	    fmt.Println("response Headers:", resp.Header)
	    body, _ := ioutil.ReadAll(resp.Body)
	    fmt.Println("response Body:", string(body))

		// delay every second
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

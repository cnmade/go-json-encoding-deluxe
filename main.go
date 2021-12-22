// You can edit this code!
// Click here and start typing.
package main

import (
	"bytes"
	"fmt"
	"github.com/cnmade/goencode/json"
	"io/ioutil"
	"net/http"
)

func main() {
	hc, err := http.NewRequest("GET", "https://gitlab.com/api/v4/templates/gitignores", bytes.NewBuffer([]byte{}))
	if err != nil {
		panic(err)
	}

	clt := http.Client{}
	response, err := clt.Do(hc)
	json.UnsupportedBehaviour = json.UnsupportedBehaviourWithSprintf
	jsonStr, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("error: %#v\n", err.Error())
		panic(err)
	}
	fmt.Printf("Response: '%+v' \n", string(jsonStr))

	fmt.Printf("Response raw: '%+v'\n", response)

	bodyAll, err := ioutil.ReadAll(response.Body)

	fmt.Println("response status: " + response.Status)
	fmt.Println("body " + string(bodyAll))
}

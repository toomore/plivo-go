// plivo
/*
	func main() {
		p := Plivo{
			Host:     "https://api.plivo.com/v1/Account",
			User:     "",
			Password: "",
		}

		data := map[string]string{
			"dst":  "",
			"src":  "",
			"text": "",
		}

		fmt.Println(p)
		p.Send(data)

		fmt.Println(p.RenderPath("/Message/"))
		fmt.Println(p.RenderPath("/Message"))
	}
*/
package plivo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Plivo struct
type Plivo struct {
	Host     string
	User     string
	Password string
}

func (p Plivo) Send(data map[string]string) {
	jsonData, _ := json.Marshal(data)
	a, _ := http.NewRequest("POST", p.renderPath("/Message/"), bytes.NewReader(jsonData))
	a.URL.User = url.UserPassword(p.User, p.Password)

	header := http.Header{}
	header.Add("Content-Type", "application/json")
	a.Header = header

	if false {
		var c http.Client
		resp, err := c.Do(a)
		if err != nil {
			fmt.Printf("Error >>> %s \n", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Resp >>> \n%s\n", body)
		}
	} else {
		fmt.Println(a, p)
	}
}

func (p Plivo) renderPath(urlpath string) string {
	URLPath, _ := url.ParseRequestURI(p.Host)
	URLPath.Path = path.Join(URLPath.Path, p.User, urlpath)
	if strings.LastIndex(urlpath, "/") >= 0 {
		return fmt.Sprintf("%s/", URLPath.String())
	}
	return URLPath.String()
}

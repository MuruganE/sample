// Sample project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"strconv"
	"time"
)

var userHistories []URLHistory

type URLHistory struct {
	Carrier     string       `json:"carrier"`
	UserResults []UserResult `json:"data"`
}

type UserResult struct {
	ID    string      `json:"id"`
	Start int64       `json:"date_start"`
	End   int64       `json:"date_end"`
	Count int         `json:"count"`
	URLS  []URLResult `json:"urls"`
}

type URLResult struct {
	URL    string `json:"url"`
	Count  int    `json:"count"`
	Reason string `json:"reason"`
}

func main() {
	fmt.Println("Hello World!", time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println("Now started at : ", i)
		/*for i1 := 0; i1 < 3; i1++ {
			for i2 := 0; i2 < 3; i2++ {
				setHistory("car:"+strconv.Itoa(i), "id:"+strconv.Itoa(i1), "url:"+strconv.Itoa(i2), "reputation")
			}
		}*/
	}

}

func setHistory(cR string, Id string, dm string, res string) {
	defaultConfigFile := "testConfig.json"
	nowT := time.Now()
	carOkay := false

	if len(userHistories) > 0 {
		for k, v := range userHistories {
			idOkay := false
			if v.Carrier == cR {
				carOkay = true
				for k1, v1 := range v.UserResults {
					urlOkday := false
					if v1.ID == Id {
						idOkay = true
						for k2, v2 := range v1.URLS {
							if v2.URL == dm {
								urlOkday = true
								v2.Count++
								v2.Reason = res
								v1.URLS[k2] = v2
								v.UserResults[k1] = v1
								userHistories[k] = v
								break
							}
						}
						if !urlOkday {
							var tmpURL = URLResult{
								URL:    dm,
								Count:  1,
								Reason: res,
							}
							v1.URLS = append(v1.URLS, tmpURL)
							v.UserResults[k1] = v1
							userHistories[k] = v
							break
						}
					}
				}
				if !idOkay {
					var tmpID = UserResult{
						ID:    Id,
						Count: 1,
						Start: nowT.Unix(),
						End:   (nowT.Add(15 * time.Minute)).Unix(),
						URLS: []URLResult{
							{URL: dm,
								Count:  1,
								Reason: res}},
					}
					v.UserResults = append(v.UserResults, tmpID)
					userHistories[k] = v
					break
				}
			}
		}
	}
	if !carOkay {

		var tmpTest URLHistory

		tmpTest.Carrier = cR
		var tmpP = URLResult{
			URL:    dm,
			Count:  1,
			Reason: res,
		}

		var tmpU UserResult
		tmpU.ID = Id
		tmpU.Count = 1
		nowT = time.Now()
		tmpU.Start = nowT.Unix()
		tmpU.End = (nowT.Add(10 * time.Minute)).Unix()
		tmpU.URLS = append(tmpU.URLS, tmpP)
		tmpTest.UserResults = append(tmpTest.UserResults, tmpU)
		fmt.Printf("Result is : \n %+v \n ", tmpTest)
		userHistories = append(userHistories, tmpTest)
	}

	b, _ := json.Marshal(userHistories)
	if err := ioutil.WriteFile(defaultConfigFile, b, 0644); err == nil {
		//fmt.Println("Update Done!")
	} else {
		fmt.Println("Update Failed due to ", err)
	}
}

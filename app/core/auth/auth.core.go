package core

import (
	constant "app/core/constant"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const zitadelPAT = "Bearer 6qhk8jJ5QONpAIiIOA93j4wvkr9y7oKIPHLXCYTBsUBUFcI_QGxKCzSetgrtjAzdbUQ05Kg"

var client = &http.Client{}

type createProjectResponse struct {
	ID string `json:"id"`
}

func CreateProject() {
	fmt.Println("Creating Pocketlink project on Zitadel...")
	body := []byte(`{
		"name": "Pocketlink",
		"projectRoleAssertion": true,
		"projectRoleCheck": true,
		"hasProjectCheck": true
	}`)
	req, reqErr := http.NewRequest("POST", constant.ZITADEL_HOST+"/management/v1/projects", bytes.NewBuffer(body))
	req.Header.Add("x-zitadel-orgid", constant.ZITADEL_ORGANIZATION_ID)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", zitadelPAT)
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	res, resErr := client.Do(req)
	if resErr != nil {
		log.Fatal(resErr)
	}

	var decodedBody createProjectResponse
	jsonErr := json.NewDecoder(res.Body).Decode(&decodedBody)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	if decodedBody.ID != "" {
		fmt.Println("Pocketlink project is created on Zitadel")
		CreateProjectRoles(decodedBody.ID)
	}

	defer res.Body.Close()
}

func CreateProjectRoles(appId string) {
	for _, roleId := range constant.ROLE_IDS {
		body := []byte(fmt.Sprintf(`{
			"roleKey": "%s",
			"displayName": "%s"
		}`, roleId, cases.Title(language.English).String(roleId)))

		req, reqErr := http.NewRequest("POST", fmt.Sprintf(constant.ZITADEL_HOST+"/management/v1/projects/%s/roles", appId), bytes.NewBuffer(body))
		req.Header.Add("x-zitadel-orgid", constant.ZITADEL_ORGANIZATION_ID)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", zitadelPAT)
		if reqErr != nil {
			log.Fatal(reqErr)
		}

		res, resErr := client.Do(req)
		if resErr != nil {
			log.Fatal(resErr)
		}
		defer res.Body.Close()
	}
	fmt.Println("Pocketlink project roles is created on Zitadel.")
}

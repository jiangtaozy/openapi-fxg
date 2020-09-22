/*
 * Maintained by jemo from 2020.9.10
 * Created by jemo on 2020.9.10 17:00:25
 * Access Token
 */

package openapiFxg

import (
  "log"
  "encoding/json"
  "net/http"
)

func AccessToken(appId string, appSecret string) map[string]interface{} {

  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    "https://openapi-fxg.jinritemai.com/oauth2/access_token",
    nil,
  )
  if err != nil {
    log.Println("access-token-new-request-error: ", err)
  }
  query := req.URL.Query()
  query.Add("app_id", appId)
  query.Add("app_secret", appSecret)
  query.Add("grant_type", "authorization_self")
  req.URL.RawQuery = query.Encode()
  resp, err := client.Do(req)
  if err != nil {
    log.Println("access-token-client-do-error: ", err)
  }
  var body map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&body)
  return body
}

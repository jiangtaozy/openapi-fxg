/*
 * Maintained by jemo from 2020.9.22 to now
 * Created by jemo on 2020.9.22 16:57:06
 * Refresh Access Token
 */

package openapiFxg

import (
  "log"
  "encoding/json"
  "net/http"
)

func RefreshAccessToken(appId string, appSecret string, refreshToken string) map[string]interface{} {

  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    "https://openapi-fxg.jinritemai.com/oauth2/refresh_token",
    nil,
  )
  if err != nil {
    log.Println("refresh-access-token-new-request-error: ", err)
  }
  query := req.URL.Query()
  query.Add("app_id", appId)
  query.Add("app_secret", appSecret)
  query.Add("grant_type", "refresh_token")
  query.Add("refresh_token", refreshToken)
  req.URL.RawQuery = query.Encode()
  resp, err := client.Do(req)
  if err != nil {
    log.Println("refresh-access-token-client-do-error: ", err)
  }
  var body map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&body)
  return body
}

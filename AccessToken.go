/*
 * Maintained by jemo from 2020.9.10
 * Created by jemo on 2020.9.10 17:00:25
 * Access Token
 */

package openapi-fxg

import (
  "log"
)

func AccessToken(appId string, appSecret string) map[string]interface{} {

  log.Println("appId: ", appId)
  log.Println("appSecret: ", appSecret)

  return map[string]interface{} {
    "access_token": "access_token",
  }
}

/*
 * Maintained by jemo from 2020.12.3 to now
 * Created by jemo on 2020.12.3 14:43:05
 * Request API
 * 发起API请求
 */

package openapiFxg

import (
  "log"
  "time"
  "errors"
  "net/http"
  "encoding/json"
)

func requestApi(appKey string, appSecret string, accessToken string, param map[string]interface{}, method string, path string) (map[string]interface{}, error) {
  paramByte, err := json.Marshal(param)
  if err != nil {
    log.Println("request-api-json-marshal-error: ", err)
    return nil, err
  }
  paramJson := string(paramByte)
  timestamp := time.Now().Format("2006-01-02 15:04:05")
  str := appSecret + "app_key" + appKey + "method" + method + "param_json" + paramJson + "timestamp" + timestamp + "v2" + appSecret
  sign, err := getMd5String(str)
  if err != nil {
    return nil, err
  }

  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    "https://openapi-fxg.jinritemai.com" + path,
    nil,
  )
  if err != nil {
    log.Println("request-api-new-request-error: ", err)
    return nil, err
  }
  query := req.URL.Query()
  query.Add("method", method)
  query.Add("app_key", appKey)
  query.Add("access_token", accessToken)
  query.Add("param_json", paramJson)
  query.Add("timestamp", timestamp)
  query.Add("v", "2")
  query.Add("sign", sign)
  query.Add("sign_method", "md5")
  req.URL.RawQuery = query.Encode()
  resp, err := client.Do(req)
  if err != nil {
    log.Println("request-api-client-do-error: ", err)
    return nil, err
  }
  var body map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&body)
  errNo := body["err_no"].(float64)
  message := body["message"].(string)
  if errNo != 0 {
    log.Println("request-api-server-error: ", message)
    return nil, errors.New(message)
  }
  data := body["data"].(map[string]interface{})
  return data, nil
}

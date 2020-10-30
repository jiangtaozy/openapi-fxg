/*
 * Maintained by jemo from 2020.9.10 to now
 * Created by jemo on 2020.9.10 16:38:40
 * Product List
 * 获取商品列表
 */

package openapiFxg

import (
  "io"
  "log"
  "fmt"
  "time"
  "net/http"
  "crypto/md5"
  "encoding/json"
)

func ProductList(appKey string, appSecret string, accessToken string) map[string]interface{} {

  method := "product.list"
  param := map[string]interface{}{
    "page": "0",
    "size": "10",
    "status": "0",
    "check_status": "3",
  }
  paramByte, err := json.Marshal(param)
  if err != nil {
    log.Println("product-list-json-marshal-error: ", err)
  }
  paramJson := string(paramByte)
  timestamp := time.Now().Format("2006-01-02 15:04:05")
  str := appSecret + "app_key" + appKey + "method" + method + "param_json" + paramJson + "timestamp" + timestamp + "v2" + appSecret
  sign := getMd5String1(str)

  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    "https://openapi-fxg.jinritemai.com/product/list",
    nil,
  )
  if err != nil {
    log.Println("product-list-new-request-error: ", err)
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
    log.Println("product-list-client-do-error: ", err)
  }
  var body map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&body)
  return body
}

func getMd5String1(str string) string {
  m := md5.New()
  _, err := io.WriteString(m, str)
  if err != nil {
    log.Fatal(err)
  }
  arr := m.Sum(nil)
  return fmt.Sprintf("%x", arr)
}

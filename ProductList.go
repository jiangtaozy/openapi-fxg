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
  "errors"
  "net/http"
  "crypto/md5"
  "encoding/json"
)

func ProductList(appKey string, appSecret string, accessToken string, param map[string]interface{}) (map[string]interface{}, error) {

  method := "product.list"
  paramByte, err := json.Marshal(param)
  if err != nil {
    log.Println("product-list-json-marshal-error: ", err)
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
    "https://openapi-fxg.jinritemai.com/product/list",
    nil,
  )
  if err != nil {
    log.Println("product-list-new-request-error: ", err)
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
    log.Println("product-list-client-do-error: ", err)
    return nil, err
  }
  var body map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&body)
  errNo := body["err_no"].(float64)
  message := body["message"].(string)
  if errNo != 0 {
    log.Println("product-list-server-error: ", message)
    return nil, errors.New(message)
  }
  data := body["data"].(map[string]interface{})
  return data, nil
}

func getMd5String(str string) (string, error) {
  m := md5.New()
  _, err := io.WriteString(m, str)
  if err != nil {
    log.Println("product-list-get-md5-string-error: ", err)
    return "", err
  }
  arr := m.Sum(nil)
  return fmt.Sprintf("%x", arr), nil
}

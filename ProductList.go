/*
 * Maintained by jemo from 2020.9.10 to now
 * Created by jemo on 2020.9.10 16:38:40
 * Product List
 * 获取商品列表
 */

package openapiFxg

func ProductList(appKey string, appSecret string, accessToken string, param map[string]interface{}) (map[string]interface{}, error) {

  method := "product.list"
  path := "/product/list"
  return requestApi(appKey, appSecret, accessToken, param, method, path)
}

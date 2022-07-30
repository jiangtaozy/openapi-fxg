/*
 * Maintained by jemo from 2020.12.3 to now
 * Created by jemo on 2020.12.3 14:30:19
 * Order List
 * 获取订单列表
 */

package openapiFxg

func OrderList(appKey string, appSecret string, accessToken string, param map[string]interface{}) (map[string]interface{}, error) {

  method := "order.list"
  path := "/order/list"
  return requestApi(appKey, appSecret, accessToken, param, method, path)
}

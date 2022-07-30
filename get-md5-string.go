/*
 * Maintained by jemo from 2020.12.3 to now
 * Created by jemo on 2020.12.3 14:47:44
 * Get MD5 String
 */

package openapiFxg

import (
  "io"
  "log"
  "fmt"
  "crypto/md5"
)

func getMd5String(str string) (string, error) {
  m := md5.New()
  _, err := io.WriteString(m, str)
  if err != nil {
    log.Println("get-md5-string-error: ", err)
    return "", err
  }
  arr := m.Sum(nil)
  return fmt.Sprintf("%x", arr), nil
}

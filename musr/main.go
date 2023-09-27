// import (
// 	"crypto/sha256"
// 	"fmt"
// 	"time"

// 	"golang.org/x/crypto/bcrypt"
// )

// func main() {
// 	// f := "jasurbeksuyunov"
// 	// d, _ := GenerateHash(f)
// 	// fmt.Println(d)
// 	// var w string
// 	// fmt.Println(" parolni kiriting :")
// 	// fmt.Scan(&w)
// 	// r := CheckPasswordIfMatchs("$2a$10$I2Yjq0Vt/rRT2bDT5OFpru.76ub2jTqDWvna0N.dTHvTlyJEaaP0i", w)
// 	// fmt.Println(".....", r)
// 	f :=GenerateToken("4c0f9a3a-46d8-4038-91ba-54e7d3c776b0")
// 	fmt.Println(f)
// }


// func GenerateHash(password string) (string, error) {
// 	password_bytes := []byte(password)

// 	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(password_bytes, bcrypt.DefaultCost)

// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedPasswordBytes), nil
// }

// // Comparing the password with the hash
// func CheckPasswordIfMatchs(hashedPassword, currPassword string) bool {
// 	err := bcrypt.CompareHashAndPassword(
// 		[]byte(hashedPassword), []byte(currPassword))
// 	return err == nil
// }
// func GenerateToken(user_id string) string {
// 	secret_key := "AssalomuAlaykumXushKelibsizUstoz"
// 	tim := time.Now().String()
// 	x := secret_key + tim + user_id + "J"
// 	h := sha256.New()
// 	h.Write([]byte(x))
// 	hash := h.Sum(nil)
// 	authTokenStep := fmt.Sprintf("%x", hash)
// 	return authTokenStep
// }

// package main

// import (
//   "fmt"
//   "strings"
//   "net/http"
//   "io/ioutil"
// )

// func main() {

//   url := "http://207.248.62.69:8090/APIPortalProveedores/api/Maestro/ObtenerMaestrosFacturaManual/"
//   method := "GET"

//   payload := strings.NewReader("Token=5149f0d8-819d-4d10-a9f9-cbe56df4e3d3&RFCProveedor=RDI841003QJ4&RFCSociedad=PLE8708273Q0")

//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, payload)

//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

//   res, err := client.Do(req)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer res.Body.Close()

//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   fmt.Println(string(body))
// }

package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://my.soliq.uz/services/lgota-bindings/getbytin?tin=499719765&lang=ru"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Basic ZmFjdHVyYXV6Ojc5NWM0bm0wMGF4MDEyYyQ=")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
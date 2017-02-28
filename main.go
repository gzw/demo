/*
* @Author: guozhiwei
* @Date:   2017-02-28 11:01:33
* @Last Modified by:   guozw
* @Last Modified time: 2017-02-28 14:52:26
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloworld"))
}

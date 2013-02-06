package main

import (
    "fmt"
    "os"
    "bytes"
    "io"
    "io/ioutil"
    "net/http"
)
const (
  PORT = ":8080"
  URL = "http://server4.operamini.com/"
)

func handler(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil && err != io.EOF {
        fmt.Println(err)
        return
    }
    req, err := http.NewRequest("POST",URL,bytes.NewReader(body))
    if err != nil && err != io.EOF {
        fmt.Println(err)
        return
    }
    req.Header.Add("content-type","application/xml")
    resp, err := http.DefaultClient.Do(req)
    if err != nil && err != io.EOF {
        fmt.Println(err)
        return
    }
    w.WriteHeader(resp.StatusCode)
    result, err := ioutil.ReadAll(resp.Body)
    if err != nil && err != io.EOF {
        fmt.Println(err)
        return
    }
    w.Write(result)
    return
}

func main() {
    http.HandleFunc("/",handler)
    http.ListenAndServe(PORT,nil)
    os.Exit(0)
}

package main

import (
    "crypto/tls"
    elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
    "log"
    "net"
    "net/http"
    "strings"
    "time"
)

func main() {
    cfg := elasticsearch7.Config{
        Addresses: []string{
            "http://localhost:9200",
        },
        Username: "foo",
        Password: "bar",
        Transport: &http.Transport{
            MaxIdleConnsPerHost:   10,
            ResponseHeaderTimeout: time.Second,
            DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
            TLSClientConfig: &tls.Config{
                MinVersion:         tls.VersionTLS11,
            },
        },
    }
    es, _ := elasticsearch7.NewClient(cfg)
    log.Println(elasticsearch7.Version)
    log.Println(es.Info())
    log.Println("OK")

    /// write index
    res, err := es.Index(
        "test",                                  // Index name
        strings.NewReader(`{"title" : "Test"}`), // Document body
        es.Index.WithDocumentID("1"),            // Document ID
        es.Index.WithRefresh("true"),            // Refresh
    )
    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }
    defer res.Body.Close()
    log.Println(res)
    // => [201 Created] {"_index":"test","_type":"_doc","_id":"1" ...
}
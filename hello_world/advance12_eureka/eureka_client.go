package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/ArthurHlt/go-eureka-client/eureka"
    "log"
    "strconv"
    "time"
)

func showJson(v interface{}) {
    b, err := json.Marshal(v)
    if err != nil {
        fmt.Printf("%+v\n", v)
    } else {
        var out bytes.Buffer
        err = json.Indent(&out, b, "", "    ")
        if err != nil {
            fmt.Printf("%+v", v)
        } else {
            fmt.Println("\n" + out.String())
        }
    }
}

func main() {

    client := eureka.NewClient([]string{
        "http://192.168.0.216:9761/eureka", //From a spring boot based eureka server
        // add others servers here
    })
    appId := "golang-demo"
    localIp := "192.168.5.134"
    hostName := localIp
    bindPort := 6002
    var ttl uint = 30

    instance := eureka.NewInstanceInfo(hostName, appId, localIp, bindPort, ttl, false) //Create a new instance to register
    instance.Metadata = &eureka.MetaData{
        Map: make(map[string]string),
    }
    instance.Metadata.Map["management.port"] = strconv.Itoa(bindPort) //add metadata for example
    err := client.RegisterInstance(appId, instance)                   // Register new instance in your eureka(s)
    if err != nil {
        fmt.Println("register error == ", err)
    }

    time.Sleep(1 * time.Second)
    apps, err := client.GetApplications() // Retrieves all applications from eureka server(s)
    if err != nil {
        fmt.Println("get error == ", err)
    }
    _, _ = client.GetApplication(instance.App)                               // retrieve the application "test"
    instanceInfo, _ := client.GetInstance(instance.App, instance.InstanceId) // retrieve the instance from "test.com" inside "test"" app

    showJson(*apps)
    showJson(*instanceInfo)

    for {
        err := client.SendHeartbeat(instance.App, instance.InstanceId) // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
        if err != nil {
            log.Print("err for heartbeat..")
        } else {
            log.Printf("heartbeat...%s", instance.App)
        }
        time.Sleep(20 * time.Second)
    }
}

# Goalteon

Goalteon is a golang client for [Alteon](https://www.radware.com/products/alteon/) made by [Radware](https://www.radware.com/).

It handles by default concurrency call on Alteon.

## Installation

run `go get github.com/ArthurHlt/goalteon` to install the package.

## Usage

```go
package main

import (
    "fmt"
    "github.com/ArthurHlt/goalteon"
)

func main() {
    client := goalteon.NewClient("https://alteon-ip", "username", "password", goalteaon.WithInsecureSkipVerify(true))
    
	// List all virtual server for example
	res, err := client.List(&beans.SlbNewCfgEnhVirtServerTable{}, nil)
    if err != nil {
        panic(err)
    }
	
	// translate to params list for ease usage
	// you can also use goalteaon.Translate to translate only one object
    resTr := goalteaon.TranslateList[*beans.SlbNewCfgEnhVirtServerTableParams](res)
    b, err := json.MarshalIndent(resTr, "", "  ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
	
	// List all virtual server with only some properties
	res, err = client.List(&beans.SlbNewCfgEnhVirtServerTable{}, goalteaon.UrlParamsProps(nil, "VirtServerIpAddress", "VirtServerBwmContract"))
    if err != nil {
        panic(err)
    }
	fmt.Println(res)
	
	// Get one virtual server
	res, err = client.Get(&beans.SlbNewCfgEnhVirtServerTable{
        SlbNewCfgEnhVirtServerIndex: "foo",
    }, nil)
    if err != nil {
        panic(err)
    }
    resTrOne := goalteaon.Translate[*beans.SlbNewCfgEnhVirtServerTableParams](res)
	fmt.Println(resTrOne)
	
	// Create a virtual server
	// by default: one modifying request can be made at the same time with a mutex lock
	// it also revert all not saved change before applying changes
	// and finally apply, save and sync changes
    status, err := client.Create(&beans.SlbNewCfgEnhVirtServerTable{
        SlbNewCfgEnhVirtServerIndex: "foo",
        VirtServerIpAddress: "172.16.0.2",
    })
    if err != nil {
        panic(err)
    }
	fmt.Println(status)
	
	// Bulk changes example
	// It let you make multiple changes in one session and apply, save and sync all those changes
	statuses, err := client.Bulk([]*goalteaon.BulkItem{
        {
            Method: goalteaon.BulkMethodCreate,
            Bean: &beans.SlbNewCfgEnhVirtServerTable{
                SlbNewCfgEnhVirtServerIndex: "foo",
                Params:  				  &beans.SlbNewCfgEnhVirtServerTableParams{
                    VirtServerIpAddress:      "172.16.0.2",
                },
            },
        },
        {
            Method: goalteaon.BulkMethodCreate,
            Bean: &beans.SlbNewCfgEnhVirtServerTable{
                SlbNewCfgEnhVirtServerIndex: "foo2",
                Params:  				  &beans.SlbNewCfgEnhVirtServerTableParams{
                    VirtServerIpAddress:      "172.16.0.3",
                },
            },
        },
    })
	if err != nil {
        panic(err)
    }
	fmt.Println(statuses)
	
	// import/export
	
	// AppShape
    data, err := client.ExportAppShape(&goalteaon.ImpExpAppShapeParams{
        ID: "APM_script",
    })
    if err != nil {
        log.Println(err)
    }
    fmt.Println(string(data))

    _, err = client.ImportAppShape([]byte(`when CLIENT_ACCEPTED {
       TCP::collect
    }
    when CLIENT_DATA {
       TCP::payload replace 0 0 "PROXY TCP[IP::version] [IP::remote_addr] [IP::local_addr] [TCP::remote_port] [TCP::local_port]\r\n"
       TCP::release
    }
    -----END`),
        &goalteaon.ImpExpAppShapeParams{
            ID: "proxy_protocol",
        },
    )
    if err != nil {
        log.Println(err)
    }
	
	// Config
	data, err := client.ExportConfig(&goalteaon.ImpExpConfigParams{
        Pkey:       true,
        Passphrase: "radware",
    })
    if err != nil {
        log.Println(err)
    }
    fmt.Println(string(data))

    _, err = client.ImportConfig(data,
        &goalteaon.ImpExpConfigParams{
            Pkey:       true,
            Passphrase: "radware",
        },
    )
    if err != nil {
        log.Println(err)
    }
	
	// SSL
	data, err := client.ExportSsl(&goalteaon.ImpExpSslParams{
        ID: "WebManagementCert",
    })
    if err != nil {
        log.Println(err)
    }
    fmt.Println(string(data))

    _, err = client.ImportSsl(data,
        &goalteaon.ImpExpSslParams{
            ID: "WebManagementCert",
        },
    )
    if err != nil {
        log.Println(err)
    }
}
```


## Advanced usage

```go
package main

import (
        goalteaon "github.com/ArhurHlt/goalteon"
"github.com/ArthurHlt/goalteon""sync"

)

// CustomLocker is dummy locker for show how to make a custom locker
// *sync.Mutex is already a valid Locker and by default goalteon use *sync.Mutex as locker
type CustomLocker struct {
	mutex *sync.Mutex
}

func (c *CustomLocker) Lock() {
	c.mutex.Lock()
}

func (c *CustomLocker) Unlock() {
	c.mutex.Unlock()
}

func main() {
	client := goalteon.NewClient(
		"https://alteon-ip", "username", "password", 
		goalteaon.WithInsecureSkipVerify(true),
		// disable lock between change request
		goalteaon.WithNoLock(),
		// disable auto apply save sync after a change(s) bean made by create/update/delete/bulk
		goalteaon.WithNoAutoApplySaveSync(),
    )
	
	clientWithCustomLocker := goalteon.NewClient(
        "https://alteon-ip", "username", "password", 
        goalteaon.WithInsecureSkipVerify(true),
        goalteaon.WithLocker(&CustomLocker{
            mutex: &sync.Mutex{},
        }),
   )
}

```

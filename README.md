[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/archaron/lbapi/blob/main/README.md)
[![ru](https://img.shields.io/badge/lang-ru-green.svg)](https://github.com/archaron/lbapi/blob/main/README.ru.md)

[![license](https://img.shields.io/github/license/archaron/lbapi.svg)](https://github.com/archaron/lbapi/blob/main/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/archaron/lbapi)](https://pkg.go.dev/mod/github.com/archaron/lbapi)
[![GitHub Workflow Status](https://github.com/archaron/lbapi/actions/workflows/go.yml/badge.svg)](https://github.com/archaron/lbapi/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/archaron/lbapi)](https://goreportcard.com/report/github.com/archaron/lbapi)
![Go version](https://img.shields.io/github/go-mod/go-version/archaron/lbapi?style=flat&label=Go%20%3E%3D)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Farcharon%2Flbapi.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Farcharon%2Flbapi?ref=badge_shield)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Farcharon%2Flbapi.svg?type=shield&issueType=security)](https://app.fossa.com/projects/git%2Bgithub.com%2Farcharon%2Flbapi?ref=badge_shield)

LBApi
========
Implementation of interaction with the LANBilling billing system via the JSON-RPC protocol.
This protocol is used by the system to communicate between agents/administrative interface and the system core.
Tested on LANBilling 2.0 build 2.0.39.

Requirements
-------------------
To work with the API, you need to create a manager on whose behalf you will log in to the system and add the host from
which you will connect to the core to "trusted hosts" in the tab `Options -> Trusted hosts` and restart LBcore.

API connection
-------------------
The client configuration is specified by the `ClientConfig` structure:
```go
	config := lbapi.ClientConfig{
		Address:         "127.0.0.1:1502", // Core endpoint IP and Port
		Username:        "admin",          // User name
		Password:        "",               // Password
		MaxFails:        5,                // Maximal allowed ping misses before reconnection
		Timeout:         5 * time.Second,  // Connection timeout
		ReconnectPeriod: 10 * time.Second, // Periodical ping and reconnect period
	}
```
`MaxFails`, `Timeout`, `ReconnectPeriod` are used to maintain connection with the billing. If `ReconnectPeriod` is set,
a periodic request for the server version is launched and if the server does not respond within `TimeOut`, it is considered unavailable.
If the number of times the server was unavailable exceeds `MaxFails`, the connection will be automatically re-established.
If successful, Login will be performed automatically and the events the client was subscribed to will be re-subscribed.
This is done for cases when the core was suddenly restarted, but we do not want to "drop" the entire service and lose events.

```go
    log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelDebug,
	}))

    lb = lbapi.NewClient(config, log)


    err = lb.ConnectAndLogin(ctx)

    if err != nil {
        panic(err)
    }

    defer lb.Close()
```

Once connected, you can make requests to the API as follows:
```go
	options, err := lb.GetSyncOptions(context.Background(), lbapi.GetSyncOptionsRequest{
		AgentID: 1,
	})

	if err != nil {
		panic(err)
	}
```

Events subscription
-------------------
When simulating the agent's work, you can subscribe to "events":
```go
    // Subscribe to specified events
    ok, err := lb.Subscribe(context.Background(), []events.LBEvent{events.ChangeAgentEvent, events.BlockVgEvent})
	if err != nil {
		panic(err)
	}

	if !ok {
        // Server was unable to subscibe
    }

    // Obtain the events channel
    ch := lb.GetEventsChannel()

    for {
        e := <-ch:
        spew.Dump(e)
    }
```

Events arrive in the channel obtained via the `GetEventsChannel` method. For testing purposes, all events known to me are added to `events.AllLBEvents`




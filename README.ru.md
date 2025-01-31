[![en](https://img.shields.io/badge/lang-en-green.svg)](https://github.com/archaron/lbapi/blob/main/README.md)
[![ru](https://img.shields.io/badge/lang-ru-red.svg)](https://github.com/archaron/lbapi/blob/main/README.ru.md)

[![license](https://img.shields.io/github/license/archaron/lbapi.svg)](https://github.com/archaron/lbapi/blob/main/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/archaron/lbapi)](https://pkg.go.dev/mod/github.com/archaron/lbapi)
[![GitHub Workflow Status](https://github.com/archaron/lbapi/actions/workflows/go.yml/badge.svg)](https://github.com/archaron/lbapi/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/archaron/lbapi)](https://goreportcard.com/report/github.com/archaron/lbapi)
![Go version](https://img.shields.io/github/go-mod/go-version/archaron/lbapi?style=flat&label=Go%20%3E%3D)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Farcharon%2Flbapi.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Farcharon%2Flbapi?ref=badge_shield)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Farcharon%2Flbapi.svg?type=shield&issueType=security)](https://app.fossa.com/projects/git%2Bgithub.com%2Farcharon%2Flbapi?ref=badge_shield)


[![Coverage Status](https://coveralls.io/repos/github/archaron/lbapi/badge.svg?branch=main)](https://coveralls.io/github/archaron/lbapi?branch=main)
[![GitHub tag](https://img.shields.io/github/tag/archaron/lbapi.svg?maxAge=86400)](https://github.com/archaron/lbapi)


LBApi
========
Реализация взаимодействия с биллинговой системой LANBilling через протокол JSON-RPC.
Данный протокол используется системой для общения между агентами/административным интерфейсом и ядром системы.
Тестировалось на LANBilling 2.0 build 2.0.39.

Подготовка к использованию
-------------------
Для работы с API требуется создать менеджера, от имени которого будет происходить вход в систему и добавить хост с
которого будет осуществляться подключение к ядру в "доверенные хосты" во вкладке `Опции -> Доверенные хосты` и перезапустить LBcore.

Подключение к API
-------------------
Конфигурация клиента задаётся структурой `ClientConfig`:
```go
	config := lbapi.ClientConfig{
		Address:         "127.0.0.1:1502", // Адрес и порт подключения
		Username:        "admin",          // Имя пользователя
		Password:        "",               // Пароль
		MaxFails:        5,                // Максимальное количество пропусков "пинга" перед переподключением
		Timeout:         5 * time.Second,  // Таймаут подключения
		ReconnectPeriod: 10 * time.Second, // Период проверки связи и переподключения
	}
```
`MaxFails`, `Timeout`, `ReconnectPeriod` используются для поддержания связи с биллингом. Если установлен `ReconnectPeriod`,
запускается периодический запрос версии сервера и в случае, если сервер не отвечает в течение `TimeOut`, он считается недоступным.
Если количество раз, когда сервер был не доступен, превысит `MaxFails`, соединение автоматически будет переустановлено.
В случае успеха, автоматически осуществится Login и события, на которые был подписан клиент переподпишутся.
Сделано это для случаев, когда внезапно перезапустили ядро, но мы не хотим "ронять" весь сервис и терять события.


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

После соединения можно осуществлять запросы к API следующим образом:
```go
	options, err := lb.GetSyncOptions(context.Background(), lbapi.GetSyncOptionsRequest{
		AgentID: 1,
	})

	if err != nil {
		panic(err)
	}
```

Подписка на события
-------------------
При эмуляции работы агента можно подписаться на "события":
```go
    // Подписываемся на нужные события
    ok, err := lb.Subscribe(context.Background(), []events.LBEvent{events.ChangeAgentEvent, events.BlockVgEvent})
	if err != nil {
		panic(err)
	}

	if !ok {
        // Сервер отказал в подписке
    }

    // Получаем канал событий
    ch := lb.GetEventsChannel()

    for {
        e := <-ch:
        spew.Dump(e)
    }
```

События приходят в канал, полученный через метод `GetEventsChannel`. Для тестов все известные мне события добавлены в `events.AllLBEvents`


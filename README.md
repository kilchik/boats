Для запуска приложения
`make run`

Note: для сборки нужен Go 1.13

Структура
`/cmd/boats/main.go` - входная точка приложения
`/internal/app/boats` - ручки
`/clients/nausys` - клиент для Nausys
`/pkg/storage` - вся работа с базой
`/pkg/syncer` - синхронизация базы с Nausys
`/static` - SPA на Vue

TODO:
- unit tests for server
- навесить индексы?
- преттифай html
- parallel data retrieval
- benchmarks
- swagger
- selenium test

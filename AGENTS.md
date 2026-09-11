# AGENTS.md

Go monorepo of three independent `package main` programs (no shared library). Each top-level dir is its own binary/server. There is no CI, no Makefile, and no linter/formatter config; use the standard Go toolchain.

## Layout & entrypoints

- `WeatherParser/` — HTTP server on `:8080`, route `/weatherdata/`. Scrapes Yandex Weather (`WeatherParser/utils.go` `yaURL`), regex-parses the HTML, and POSTs the result to `dbService`.
- `dbService/` — HTTP server on `:8080`, route `/dbservice/weatherdata` (POST only, gorilla/mux). Writes each reading to Firestore.
- `utility/` — CLI, not a server: `go run ./utility -mode backupData|restoreData`.

Run one with e.g. `go run ./WeatherParser`. Each dir also has its own `Dockerfile` (pinned to `golang:1.15`).

## Build / verify

- **`CGO_ENABLED=0` is required in this environment.** The Firestore SDK pulls in cgo and the C headers are missing, so plain `go build ./...` / `go vet ./...` / `go test ./...` fail with `pthread.h: No such file` until you disable cgo.
  - `CGO_ENABLED=0 go build ./...`
  - `CGO_ENABLED=0 go vet ./...`
  - `CGO_ENABLED=0 go test ./WeatherParser/`
- `go.mod` says `go 1.16`; code uses deprecated `io/ioutil` throughout — match the existing style, don't "modernize" imports opportunistically.
- Single test: `CGO_ENABLED=0 go test ./WeatherParser/ -run TestGetTemperature -v`.

## Testing quirks

Tests live only in `WeatherParser/`. Two distinct kinds:

- **Offline/safe:** `tempParser_test.go`, `humidityParser_test.go`, `airPressParser_test.go`, `weathCondParser_test.go`, `windDir_test.go`, `windSpeedParser_test.go` — table tests over inline HTML snippets. These pass without network and are the ones to rely on.
- **Live-network:** `parser_test.go`'s `TestParserFull` and `TestYaPageAvailable` fetch `https://yandex.ru/pogoda/` at test time. They need internet and break whenever Yandex changes its markup — treat failures there as environmental, not necessarily your bug.
- **Benchmarks** (`BenchmarkSyncParser`, `BenchmarkAsyncParser`) read `WeatherParser/yaPage.txt`, which is **not committed and not produced by any test** (`TestParserFull` writes a differently-named `yaPageTest.txt`). Create `yaPage.txt` first (the unused helper `readAndSaveYaPage()` does exactly this) or benchmarks parse an empty page.

## Cross-service coupling (edit both sides)

Types are duplicated by hand across packages and kept in sync manually — there is no shared module:

- JSON contract: `WeatherParser/weatherData.go` `WeatherData` and `dbService/weatherData.go` `weatherData` must keep identical JSON tags (`temp`, `windSpd`, `windDir`, `weatherCond`, `humidity`, `airPress`, `timeStamp`). Change one, change the other.
- Firestore model: `dbService/fireStoreWeatherData.go` and `utility/fireStoreWeatherData.go` both define `fsWeatherData`, `weatherDataCollectionName`, `save()`, `key()`. Keep them consistent.

## Runtime / config gotchas

- **Hardcoded Cloud Run URL:** `WeatherParser/weatherData.go` POSTs to `https://fire-store-service-7dw6xhm35q-ey.a.run.app`, not `localhost`. Running the parser locally writes to the *deployed* dbService/Firestore.
- **Firestore:** project `weather-calend-fires` is hardcoded; collection `weather-data`; doc key format `Day.Month.Year-Hour` (`fsWeatherData.key()`), i.e. one document per hour — same-hour writes overwrite.
- **Credentials:** `dbService` and `utility` need `GOOGLE_APPLICATION_CREDENTIALS` pointing at a service-account key JSON (see `dbService/Dockerfile` and `utility/backupData.ps1`). The key file is not in the repo.
- `WeatherParser` sets `TZ=Europe/Moscow` in its Dockerfile; timestamps are timezone-sensitive.

## Parser fragility

`WeatherParser/parser.go` scrapes via regexes matched to specific **Russian** Yandex markup (e.g. `Текущая температура`, wind `title="Ветер: ..."`). Note it normalizes Yandex's Unicode minus `−` (U+2212) to ASCII `-`. Any Yandex HTML change silently breaks extraction (getters return a `errXNotFound` sentinel and a zero value rather than failing loudly).

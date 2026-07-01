<div align="center">

# device-manager

[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](LICENSE)

</div>

Realistic Telegram client device profiles for [mtgo](https://github.com/mtgo-labs/mtgo) — generate device identities that Telegram sees during `initConnection`.

## Features

- **10 device types**: Android, Android-X, iOS, macOS, Windows, Linux, Desktop, Web Z, Web K, Webogram
- **Static presets**: `TelegramDesktop()`, `TelegramAndroid()`, …
- **Deterministic generation**: same uniqueID always yields the same model/version — stable session identities across restarts
- **Direct Config injection**: `profile.Apply(&cfg)` sets `cfg.Device` in one call
- **Race-free**: all device lists use `sync.Once` lazy init

## Install

```bash
go get github.com/mtgo-labs/device-manager
```

## Quick start

```go
import (
    "github.com/mtgo-labs/mtgo/telegram"
    "github.com/mtgo-labs/device-manager"
)

cfg := telegram.DefaultConfig

// Deterministic profile from a session ID
device.Android.Generate("my-session").Apply(&cfg)

client, err := telegram.NewClient(apiID, apiHash, &cfg)
```

### Device types

| Type | Description |
|------|-------------|
| `device.Android` | Official Telegram Android |
| `device.AndroidX` | Telegram-X for Android |
| `device.IOS` | Official Telegram iOS |
| `device.MacOS` | Official Telegram macOS |
| `device.Windows` | Telegram Desktop (Windows) |
| `device.Linux` | Telegram Desktop (Linux) |
| `device.Desktop` | Telegram Desktop (random OS) |
| `device.WebZ` | Telegram Web Z (React) |
| `device.WebK` | Telegram Web K |
| `device.Webogram` | Telegram Webogram |

## License

Apache License 2.0

package device

import (
	"sync"
	"testing"

	"github.com/mtgo-labs/mtgo/telegram"
	"github.com/mtgo-labs/mtgo/telegram/types"
)

func TestTelegramDesktop(t *testing.T) {
	p := TelegramDesktop()
	if p.LangPack != "tdesktop" {
		t.Errorf("expected LangPack 'tdesktop', got %q", p.LangPack)
	}
	if p.Platform != types.ClientPlatformDesktop {
		t.Errorf("expected Platform %q, got %q", types.ClientPlatformDesktop, p.Platform)
	}
}

func TestTelegramAndroid(t *testing.T) {
	p := TelegramAndroid()
	if p.Platform != types.ClientPlatformAndroid {
		t.Errorf("expected Platform %q, got %q", types.ClientPlatformAndroid, p.Platform)
	}
	if p.LangPack != "android" {
		t.Errorf("expected LangPack 'android', got %q", p.LangPack)
	}
}

func TestTelegramIOS(t *testing.T) {
	p := TelegramIOS()
	if p.Platform != types.ClientPlatformIOS {
		t.Errorf("expected Platform %q, got %q", types.ClientPlatformIOS, p.Platform)
	}
}

func TestTelegramMacOS(t *testing.T) {
	p := TelegramMacOS()
	if p.LangPack != "macos" {
		t.Errorf("expected LangPack 'macos', got %q", p.LangPack)
	}
}

func TestTelegramWebZ(t *testing.T) {
	p := TelegramWebZ()
	if p.Platform != types.ClientPlatformWeb {
		t.Errorf("expected Platform %q, got %q", types.ClientPlatformWeb, p.Platform)
	}
}

func TestTelegramWebK(t *testing.T) {
	p := TelegramWebK()
	if p.LangPack != "macos" {
		t.Errorf("expected LangPack 'macos', got %q", p.LangPack)
	}
}

func TestDeviceGenerateAndroid(t *testing.T) {
	p := Android.Generate("test-session")
	if p.DeviceModel == "" {
		t.Error("expected non-empty DeviceModel")
	}
	if p.SystemVersion == "" {
		t.Error("expected non-empty SystemVersion")
	}
}

func TestDeviceGenerateIOS(t *testing.T) {
	p := IOS.Generate("test-session")
	if p.DeviceModel == "" {
		t.Error("expected non-empty DeviceModel")
	}
}

func TestDeviceGenerateWindows(t *testing.T) {
	p := Windows.Generate("test-session")
	if p.DeviceModel == "" {
		t.Error("expected non-empty DeviceModel")
	}
}

func TestDeviceGenerateLinux(t *testing.T) {
	p := Linux.Generate("test-session")
	if p.DeviceModel == "" {
		t.Error("expected non-empty DeviceModel")
	}
}

func TestDeviceGenerateMacOS(t *testing.T) {
	p := MacOS.Generate("test-session")
	if p.DeviceModel == "" {
		t.Error("expected non-empty DeviceModel")
	}
}

func TestDeviceGenerateDesktop(t *testing.T) {
	p := Desktop.Generate("test-session")
	if p.DeviceModel == "" {
		t.Error("expected non-empty DeviceModel")
	}
}

func TestDeviceGenerateWebZ(t *testing.T) {
	p := WebZ.Generate("")
	if p.AppVersion != "1.28.3 Z" {
		t.Errorf("expected AppVersion '1.28.3 Z', got %q", p.AppVersion)
	}
}

func TestDeviceGenerateUnknown(t *testing.T) {
	p := Device("unknown").Generate("test-session")
	if p.DeviceModel == "" {
		t.Error("expected fallback to generate non-empty DeviceModel")
	}
}

func TestDeviceGenerateDeterministic(t *testing.T) {
	p1 := Windows.Generate("same-id")
	p2 := Windows.Generate("same-id")
	if p1.DeviceModel != p2.DeviceModel {
		t.Errorf("expected same DeviceModel, got %q and %q", p1.DeviceModel, p2.DeviceModel)
	}
	if p1.SystemVersion != p2.SystemVersion {
		t.Errorf("expected same SystemVersion, got %q and %q", p1.SystemVersion, p2.SystemVersion)
	}
}

func TestProfileCopy(t *testing.T) {
	p := TelegramDesktop()
	cp := p.Copy()
	if cp.DeviceModel != p.DeviceModel {
		t.Error("copy should have same DeviceModel")
	}
	cp.DeviceModel = "modified"
	if p.DeviceModel == "modified" {
		t.Error("original should not be affected by copy modification")
	}
}

func TestProfileWithDevice(t *testing.T) {
	p := TelegramDesktop()
	modified := p.WithDevice("TestModel", "TestOS")
	if modified.DeviceModel != "TestModel" {
		t.Errorf("expected DeviceModel 'TestModel', got %q", modified.DeviceModel)
	}
	if modified.SystemVersion != "TestOS" {
		t.Errorf("expected SystemVersion 'TestOS', got %q", modified.SystemVersion)
	}
	if p.DeviceModel == "TestModel" {
		t.Error("original should not be modified")
	}
}

func TestProfileString(t *testing.T) {
	p := TelegramAndroid()
	s := p.String()
	if s == "" {
		t.Error("String() should not be empty")
	}
}

func TestProfileApply(t *testing.T) {
	p := TelegramAndroid()
	cfg := telegram.DefaultConfig
	p.Apply(&cfg)

	if cfg.Device.DeviceModel != p.DeviceModel {
		t.Errorf("expected DeviceModel %q, got %q", p.DeviceModel, cfg.Device.DeviceModel)
	}
	if cfg.Device.SystemVersion != p.SystemVersion {
		t.Errorf("expected SystemVersion %q, got %q", p.SystemVersion, cfg.Device.SystemVersion)
	}
	if cfg.Device.LangPack != p.LangPack {
		t.Errorf("expected LangPack %q, got %q", p.LangPack, cfg.Device.LangPack)
	}
	if cfg.Device.ClientPlatform != p.Platform {
		t.Errorf("expected ClientPlatform %q, got %q", p.Platform, cfg.Device.ClientPlatform)
	}
}

func TestProfileApplyPreservesOtherFields(t *testing.T) {
	p := TelegramIOS()
	cfg := telegram.DefaultConfig
	cfg.APIID = 12345
	cfg.APIHash = "my-hash"
	cfg.SessionName = "my-bot"
	cfg.InMemory = true
	p.Apply(&cfg)

	if cfg.APIID != 12345 {
		t.Error("Apply should not overwrite APIID")
	}
	if cfg.APIHash != "my-hash" {
		t.Error("Apply should not overwrite APIHash")
	}
	if cfg.SessionName != "my-bot" {
		t.Error("Apply should not overwrite SessionName")
	}
	if !cfg.InMemory {
		t.Error("Apply should not overwrite InMemory")
	}
}

func TestToDeviceConfig(t *testing.T) {
	p := TelegramMacOS()
	dc := p.ToDeviceConfig()

	if dc.DeviceModel != p.DeviceModel {
		t.Errorf("expected DeviceModel %q, got %q", p.DeviceModel, dc.DeviceModel)
	}
	if dc.AppVersion != p.AppVersion {
		t.Errorf("expected AppVersion %q, got %q", p.AppVersion, dc.AppVersion)
	}
	if dc.ClientPlatform != p.Platform {
		t.Errorf("expected ClientPlatform %q, got %q", p.Platform, dc.ClientPlatform)
	}
}

func TestGenerateAndroidNonEmpty(t *testing.T) {
	for i := range 100 {
		p := GenerateAndroid("session-" + string(rune('a'+i%26)))
		if p.DeviceModel == "" || p.SystemVersion == "" {
			t.Fatalf("iter %d: expected non-empty device fields", i)
		}
	}
}


func TestTelegramWebogram(t *testing.T) {
	p := TelegramWebogram()
	if p.AppVersion != "0.7.0" {
		t.Errorf("expected AppVersion '0.7.0', got %q", p.AppVersion)
	}
	if p.Platform != types.ClientPlatformWeb {
		t.Errorf("expected Platform %q, got %q", types.ClientPlatformWeb, p.Platform)
	}
}

func TestDeviceGenerateWebogram(t *testing.T) {
	p := Webogram.Generate("test")
	if p.AppVersion != "0.7.0" {
		t.Errorf("expected AppVersion '0.7.0', got %q", p.AppVersion)
	}
}

func TestConcurrentGenerate(t *testing.T) {
	// All device types — exercises every lazy-init path concurrently.
	devices := []Device{Android, AndroidX, IOS, MacOS, Windows, Linux, Desktop}

	var wg sync.WaitGroup
	for range 200 {
		for _, d := range devices {
			wg.Add(1)
			go func(d Device) {
				defer wg.Done()
				p := d.Generate("concurrent-test")
				if p.DeviceModel == "" {
					t.Error("expected non-empty DeviceModel")
				}
			}(d)
		}
	}
	wg.Wait()
}
package main

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Проверка Windows XP
func isWindowsXP() bool {
	version := syscall.RtlGetVersion
	if version == nil {
		return false
	}

	type OSVERSIONINFOEX struct {
		dwOSVersionInfoSize uint32
		dwMajorVersion      uint32
		dwMinorVersion      uint32
		dwBuildNumber       uint32
		dwPlatformId        uint32
		szCSDVersion        [128]uint16
		wServicePackMajor   uint16
		wServicePackMinor   uint16
		wSuiteMask          uint16
		wProductType        byte
		wReserved           byte
	}

	var info OSVERSIONINFOEX
	info.dwOSVersionInfoSize = uint32(unsafe.Sizeof(info))

	version(&info)

	// Windows XP = 5.1
	if info.dwMajorVersion == 5 && info.dwMinorVersion == 1 {
		return true
	}

	// Windows Server 2003 / XP x64 = 5.2
	if info.dwMajorVersion == 5 && info.dwMinorVersion == 2 {
		return true
	}

	return false
}

// Реальное Windows XP окно ошибки
func showXPError() {
	user32 := syscall.NewLazyDLL("user32.dll")
	msgbox := user32.NewProc("MessageBoxW")

	text := syscall.StringToUTF16Ptr("Ошибка запуска!\n\nЭта версия Windows больше не поддерживается.\nТребуется Windows 7 или новее.")
	title := syscall.StringToUTF16Ptr("DualMonitor Client – Ошибка")

	msgbox.Call(
		0,
		uintptr(unsafe.Pointer(text)),
		uintptr(unsafe.Pointer(title)),
		0x00000010, // MB_ICONERROR
	)

	os.Exit(1)
}

func mainWindow(a fyne.App) {
	w := a.NewWindow("DualMonitor Client")
	w.Resize(fyne.NewSize(800, 500))

	w.SetContent(container.NewVBox(
		widget.NewLabel("Добро пожаловать!\nВторой монитор активен."),
		widget.NewButton("Выход", func() { a.Quit() }),
	))

	w.Show()
}

func main() {
	// Windows XP → ошибка
	if runtime.GOOS == "windows" && isWindowsXP() {
		showXPError()
	}

	// Windows 7+ → запуск программы
	a := app.New()
	mainWindow(a)
	a.Run()
}

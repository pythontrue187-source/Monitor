
# Клиент второго монитора (Windows 7/8/10/11).  
Windows XP → показывает реальное системное окно ошибки.

## Поддержка
✔ Windows 7  
✔ Windows 8  
✔ Windows 10  
✔ Windows 11  
✖ Windows XP (отображается ошибка)

## Сборка (Linux / Termux / Android / macOS)

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
CC=x86_64-w64-mingw32-gcc \
go build -o DualMonitorClient.exe main.go

## Сборка на Windows

go build -o DualMonitorClient.exe main.go

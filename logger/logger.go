package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	// Logs klasörünü oluştur
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}

	// Log dosyalarının isimlerini tarih ile oluştur
	currentTime := time.Now().Format("2006-01-02")
	infoLogFile := filepath.Join(logsDir, fmt.Sprintf("info-%s.log", currentTime))
	errorLogFile := filepath.Join(logsDir, fmt.Sprintf("error-%s.log", currentTime))

	// Log dosyalarını aç
	infoFile, err := os.OpenFile(infoLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open info log file:", err)
	}

	errorFile, err := os.OpenFile(errorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open error log file:", err)
	}

	// Logger'ları oluştur
	InfoLogger = log.New(io.MultiWriter(os.Stdout, infoFile),
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger = log.New(io.MultiWriter(os.Stderr, errorFile),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo bilgi loglarını kaydeder
func LogInfo(message string, v ...interface{}) {
	InfoLogger.Printf(message, v...)
}

// LogError hata loglarını kaydeder
func LogError(message string, v ...interface{}) {
	ErrorLogger.Printf(message, v...)
}

// LogRequest HTTP isteklerini loglar
func LogRequest(method, path, ip string, status int, duration time.Duration) {
	InfoLogger.Printf("Request: %s %s | IP: %s | Status: %d | Duration: %v",
		method, path, ip, status, duration)
}

// LogErrorWithContext hata loglarını context ile kaydeder
func LogErrorWithContext(context, message string, err error) {
	ErrorLogger.Printf("[%s] %s: %v", context, message, err)
} 
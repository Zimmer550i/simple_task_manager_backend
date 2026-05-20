package handlers

import (
	"os"
	"runtime"
	"time"

	"github.com/Zimmer550i/simple_task_manager_backend/utils"
	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status      string       `json:"status"`
	Message     string       `json:"message"`
	Service     ServiceInfo  `json:"service"`
	Server      ServerInfo   `json:"server"`
	Runtime     RuntimeInfo  `json:"runtime"`
	Memory      MemoryInfo   `json:"memory"`
	CheckedAt   string       `json:"checked_at"`
}

type ServiceInfo struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	BaseURL     string `json:"base_url"`
	Port        string `json:"port"`
}

type ServerInfo struct {
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	CPUs     int    `json:"cpus"`
}

type RuntimeInfo struct {
	GoVersion    string `json:"go_version"`
	NumGoroutine int    `json:"num_goroutine"`
}

type MemoryInfo struct {
	AllocMB      uint64 `json:"alloc_mb"`
	TotalAllocMB uint64 `json:"total_alloc_mb"`
	SysMB        uint64 `json:"sys_mb"`
	NumGC        uint32 `json:"num_gc"`
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func HealthCheck(ctx *gin.Context) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	port := getEnv("PORT", "8080")
	baseURL := getEnv("BASE_URL", "http://localhost:"+port)
	environment := getEnv("APP_ENV", "development")
	serviceName := getEnv("APP_NAME", "simple-task-manager-backend")

	response := HealthResponse{
		Status:  "ok",
		Message: "Server is running",
		Service: ServiceInfo{
			Name:        serviceName,
			Environment: environment,
			BaseURL:     baseURL,
			Port:        port,
		},
		Server: ServerInfo{
			Hostname: hostname,
			OS:       runtime.GOOS,
			Arch:     runtime.GOARCH,
			CPUs:     runtime.NumCPU(),
		},
		Runtime: RuntimeInfo{
			GoVersion:    runtime.Version(),
			NumGoroutine: runtime.NumGoroutine(),
		},
		Memory: MemoryInfo{
			AllocMB:      memStats.Alloc / 1024 / 1024,
			TotalAllocMB: memStats.TotalAlloc / 1024 / 1024,
			SysMB:        memStats.Sys / 1024 / 1024,
			NumGC:        memStats.NumGC,
		},
		CheckedAt: time.Now().UTC().Format(time.RFC3339),
	}

	utils.Success(ctx, "Health check completed", response)
}

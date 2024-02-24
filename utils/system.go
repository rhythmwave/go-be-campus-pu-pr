package utils

import (
	"net/http"
	"runtime"
)

type MemoryUsage struct {
	Alloc      float64
	TotalAlloc float64
	Sys        float64
	NumGC      uint32
}

func GetMemoryUsage() MemoryUsage {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return MemoryUsage{
		Alloc:      bToMb(m.Alloc),
		TotalAlloc: bToMb(m.TotalAlloc),
		Sys:        bToMb(m.Sys),
		NumGC:      m.NumGC,
	}
}

func bToMb(b uint64) float64 {
	return float64(b) / 1024 / 1024
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

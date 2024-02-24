package models

import "time"

type GetAdminActivityLog struct {
	Id            string    `db:"id"`
	AdminId       string    `db:"admin_id"`
	AdminName     string    `db:"admin_name"`
	AdminUsername string    `db:"admin_username"`
	Module        string    `db:"module"`
	Action        string    `db:"action"`
	IpAddress     string    `db:"ip_address"`
	UserAgent     string    `db:"user_agent"`
	ExecutionTime float64   `db:"execution_time"`
	MemoryUsage   float64   `db:"memory_usage"`
	CreatedAt     time.Time `db:"created_at"`
}

type CreateAdminActivityLog struct {
	AdminId       string  `db:"admin_id"`
	AdminName     string  `db:"admin_name"`
	AdminUsername string  `db:"admin_username"`
	Module        string  `db:"module"`
	Action        string  `db:"action"`
	IpAddress     string  `db:"ip_address"`
	UserAgent     string  `db:"user_agent"`
	ExecutionTime float64 `db:"execution_time"`
	MemoryUsage   float64 `db:"memory_usage"`
}

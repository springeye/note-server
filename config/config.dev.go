//go:build !pro

package config

func init() {
	println("enable debug")
	DefaultConfig.Debug = true
}

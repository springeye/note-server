//go:build pro

package config

func init() {
	println("disable debug")
	DefaultConfig.Debug = false
}

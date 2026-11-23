package config_test

import (
	"maps"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/if0rget/my_lib/server/internal/config"
)

func TestLoad(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("DATABASE_URL", "postgres://x")
	t.Setenv("JWT_SECRET", "0123456789abcdef")
	t.Setenv("ADMIN_USERNAME", "")
	t.Setenv("ADMIN_PASSWORD", "")

	cfg, err := config.Load()
	require.NoError(t, err)
	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, "postgres://x", cfg.DatabaseURL)
	assert.Equal(t, "0123456789abcdef", cfg.JWTSecret)
	assert.Empty(t, cfg.AdminUsername)
	assert.Empty(t, cfg.AdminPassword)
	assert.Equal(t, 24*time.Hour, cfg.JWTTTL)
}

func TestLoad_CustomValues(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("DATABASE_URL", "postgres://x")
	t.Setenv("JWT_SECRET", "0123456789abcdef")
	t.Setenv("ADMIN_USERNAME", "admin")
	t.Setenv("ADMIN_PASSWORD", "secret-pass")

	cfg, err := config.Load()
	require.NoError(t, err)
	assert.Equal(t, "9090", cfg.Port)
	assert.Equal(t, "admin", cfg.AdminUsername)
	assert.Equal(t, "secret-pass", cfg.AdminPassword)
}

func cloneAndUpdate(m map[string]string, key string, value string) map[string]string {
	if m == nil {
		return nil
	}

	another := maps.Clone(m)
	another[key] = value

	return another
}

func TestLoad_Invalid(t *testing.T) {
	okENV := map[string]string{
		"PORT":         "8080",
		"DATABASE_URL": "postgres://x",
		"JWT_SECRET":   "0123456789abcdef",
	}

	cases := map[string]map[string]string{
		"missing DATABASE_URL":                 cloneAndUpdate(okENV, "DATABASE_URL", ""),
		"missing JWT_SECRET":                   cloneAndUpdate(okENV, "JWT_SECRET", ""),
		"JWT_SECRET must be at least 16 bytes": cloneAndUpdate(okENV, "JWT_SECRET", "0123456789abcde"),
		"PORT must be a number":                cloneAndUpdate(okENV, "PORT", "0x1"),
		"PORT cannot be zero":                  cloneAndUpdate(okENV, "PORT", "0"),
		"PORT cannot be negative":              cloneAndUpdate(okENV, "PORT", "-8080"),
		"PORT cannot have plus sign":           cloneAndUpdate(okENV, "PORT", "+8080"),
		"PORT cannot exceed 65535":             cloneAndUpdate(okENV, "PORT", "65536"),
		"PORT cannot be 99999":                 cloneAndUpdate(okENV, "PORT", "99999"),
	}

	for name, env := range cases {
		t.Run(name, func(t *testing.T) {
			for k, v := range env {
				t.Setenv(k, v)
			}
			_, err := config.Load()
			assert.Error(t, err)
		})
	}
}

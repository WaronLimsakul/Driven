package config

import (
	"os"
	"testing"
)

func TestNewServerConfig(t *testing.T) {
	type testStruct struct {
		name    string
		port    string
		env     string
		wantErr bool
	}

	testCases := []testStruct{
		{"normal", "5000", "production", false},
		{"no port", "", "production", true},
		{"no env", "5000", "", true},
	}

	for _, test := range testCases {
		os.Setenv("PORT", test.port)
		os.Setenv("ENV", test.env)
		config, err := NewServerConfig()
		if !test.wantErr {
			if err != nil {
				t.Fatalf("test %s: expect no error but found %v", test.name, err)
			} else if config.Env != test.env {
				t.Fatalf("test %s: expect env %s but found %s", test.name, test.env, config.Env)
			} else if config.Port != test.port {
				t.Fatalf("test %s: expect port %s but found %s", test.name, test.port, config.Env)
			}
		} else if test.wantErr && err == nil {
			t.Fatalf("test %s: expect error", test.name)
		}
	}

}

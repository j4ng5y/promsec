package promsec

import (
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
)

func tmpFile() *os.File {
	c := NewDefaultConfig()

	if err := c.SafeWriteConfigAs("tmp_config.yaml"); err != nil {
		log.Fatal().Err(err)
	}

	f, err := os.Open("tmp_config.yaml")
	if err != nil {
		log.Fatal().Err(err)
	}

	return f
}

func rmTmpFile() {
	if err := os.Remove("tmp_config.yaml"); err != nil {
		log.Fatal().Err(err)
	}
}

func Test_NewDefaultConfig(t *testing.T) {
	c := NewDefaultConfig()

	if c.GetString("version") != "v1" {
		t.Errorf("metadata version mismatch: expected \"v1\", got \"%s\"", c.GetString("version"))
	}
	if c.GetString("kind") != "PromsecConfig" {
		t.Errorf("metadata kind mismatch: expected \"PromsecConfig\", got \"%s\"", c.GetString("kind"))
	}
	if c.GetString("server.host") != "0.0.0.0" {
		t.Errorf("server host mismatch: expected \"0.0.0.0\", got \"%s\"", c.GetString("server.host"))
	}
	if c.GetUint("server.port") != 5001 {
		t.Errorf("server port mismatch: expected \"5001\", got \"%d\"", c.GetUint("server.port"))
	}
	if c.GetString("server.endpoint") != "/metrics" {
		t.Errorf("server endpoint mismatch: expected \"/metrics\", got \"%s\"", c.GetString("server.endpoint"))
	}
	if c.GetDuration("server.read_timeout") != 5*time.Second {
		t.Errorf("server read timeout mismatch: expected \"5s\", got \"%s\"", c.GetDuration("server.read_timeout").String())
	}
	if c.GetDuration("server.read_header_timeout") != 5*time.Second {
		t.Errorf("server read header timeout mismatch: expected \"5s\", got \"%s\"", c.GetDuration("server.read_header_timeout").String())
	}
	if c.GetDuration("server.write_timeout") != 5*time.Second {
		t.Errorf("server write timeout mismatch: expected \"5s\", got \"%s\"", c.GetDuration("server.write_timeout").String())
	}
	if c.GetDuration("server.idle_timeout") != 10*time.Second {
		t.Errorf("server idle timeout mismatch: expected \"10s\", got \"%s\"", c.GetDuration("server.idle_timeout").String())
	}
}

func Test_NewConfigFromFile(t *testing.T) {
	_, err := NewConfigFromFile(nil)
	if err == nil {
		t.Errorf("new config from nil file should have failed, but didn't")
	}

	f := tmpFile()
	defer rmTmpFile()

	c, err := NewConfigFromFile(f)
	if err != nil {
		t.Fatal(err)
	}

	if c.GetString("version") != "v1" {
		t.Errorf("metadata version mismatch: expected \"v1\", got \"%s\"", c.GetString("version"))
	}
	if c.GetString("kind") != "PromsecConfig" {
		t.Errorf("metadata kind mismatch: expected \"PromsecConfig\", got \"%s\"", c.GetString("kind"))
	}
	if c.GetString("server.host") != "0.0.0.0" {
		t.Errorf("server host mismatch: expected \"0.0.0.0\", got \"%s\"", c.GetString("server.host"))
	}
	if c.GetUint("server.port") != 5001 {
		t.Errorf("server port mismatch: expected \"5001\", got \"%d\"", c.GetUint("server.port"))
	}
	if c.GetString("server.endpoint") != "/metrics" {
		t.Errorf("server endpoint mismatch: expected \"/metrics\", got \"%s\"", c.GetString("server.endpoint"))
	}
	if c.GetDuration("server.read_timeout") != 5*time.Second {
		t.Errorf("server read timeout mismatch: expected \"5s\", got \"%s\"", c.GetDuration("server.read_timeout").String())
	}
	if c.GetDuration("server.read_header_timeout") != 5*time.Second {
		t.Errorf("server read header timeout mismatch: expected \"5s\", got \"%s\"", c.GetDuration("server.read_header_timeout").String())
	}
	if c.GetDuration("server.write_timeout") != 5*time.Second {
		t.Errorf("server write timeout mismatch: expected \"5s\", got \"%s\"", c.GetDuration("server.write_timeout").String())
	}
	if c.GetDuration("server.idle_timeout") != 10*time.Second {
		t.Errorf("server idle timeout mismatch: expected \"10s\", got \"%s\"", c.GetDuration("server.idle_timeout").String())
	}
}

package envconfig

import (
	"os"
	"testing"
)

type test_data_struct struct {
	MapField map[string]string `envconfig:"MAP_FIELD"`
}

func TestEmptyEnv(t *testing.T) {
	os.Clearenv()

	var s test_data_struct
	if err := Process("", &s); err != nil {
		t.Error(err)
	}

	if s.MapField != nil {
		t.Errorf("Map not nil %v", s.MapField)
	}
}

func TestEmptyEnvVar(t *testing.T) {
	os.Clearenv()
	os.Setenv("MAP_FIELD", "")

	var s test_data_struct
	if err := Process("", &s); err != nil {
		t.Error(err)
	}

	if len(s.MapField) != 0 {
		t.Errorf("Map not nil %v", s.MapField)
	}
}

func TestWhitespaceEnvVar(t *testing.T) {
	os.Clearenv()
	os.Setenv("MAP_FIELD", " ")

	var s test_data_struct
	if err := Process("", &s); err != nil {
		t.Error(err)
	}

	if len(s.MapField) != 0 {
		t.Errorf("Map not nil %v", s.MapField)
	}
}

func TestInvisibleEnvVar(t *testing.T) {
	os.Clearenv()
	os.Setenv("MAP_FIELD", "\u2063")

	var s test_data_struct
	if err := Process("", &s); err != nil {
		t.Error(err)
	}

	if len(s.MapField) != 0 {
		t.Errorf("Map not nil %v", s.MapField)
	}
}

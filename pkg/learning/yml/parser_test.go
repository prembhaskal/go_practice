package yml_test

import (
	"testing"
	"github.com/prembhaskal/go_practice/pkg/learning/yml"
)

func TestParsingYaml(t *testing.T) {
	yml.ParseYaml()
}

func TestDuplicateKeys(t *testing.T) {
	yml.ParseYamlWithDuplicateKeys()
}

func TestDuplicateKeysV3(t *testing.T) {
	err := yml.ParseYamlWithDuplicateKeysWithV3()
	if err == nil {
		t.Fail()
	}
}
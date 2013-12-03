package backend

import (
	"testing"
)

func TestKey(t *testing.T) {

	components := []string{"test", "key", "components"}

	if key := Key(components...); key != "test/key/components" {
		t.Errorf("Expected test/key/components but got %s", key)
	}
}

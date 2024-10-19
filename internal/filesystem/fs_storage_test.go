package filesystem

import (
	"bytes"
	"os"
	"testing"
)

func TestStorageMethods(t *testing.T) {
	st, err := Open("./testdata")
	if err != nil {
		t.Fatal(err)
	}

	name := "mytestfile"
	data := []byte("contents of my test file")

	defer os.Remove(name) // just in case

	if err := st.Save(name, data); err != nil {
		t.Fatal(err)
	}

	if err := st.Save(name, data); err != nil {
		t.Errorf("rewrite operatoin failed: %v", err)
	}

	b, err := st.Load(name)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, b) {
		t.Error("loaded file differs from original")
	}

	if err := st.Delete(name); err != nil {
		t.Errorf("delete failed: %v", err)
	}

	if err := st.Delete(name); err != nil {
		t.Errorf("delete of non-existent failed: %v", err)
	}

}

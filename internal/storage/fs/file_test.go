package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndGet(t *testing.T) {
	dir, err := ioutil.TempDir("", "test_save_and_get")
	if err != nil {
		t.Fatalf("error creating temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, "test.db")
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("error creating test file: %v", err)
	}
	defer file.Close()

	fs, err := NewFs(file)
	if err != nil {
		t.Fatalf("error creating fs: %v", err)
	}
	defer fs.Close()

	original1 := "https://example.com/long/url1"
	short1, err := fs.Save(original1)
	if err != nil {
		t.Fatalf("error saving original1: %v", err)
	}

	if short1 == "" {
		t.Fatalf("expected short1 to not be empty, got %q", short1)
	}

	if fs.Get(short1) != original1 {
		t.Fatalf("cache[short1] is not original1")
	}

	original2 := "https://example.com/long/url2"
	short2, err := fs.Save(original2)
	if err != nil {
		t.Fatalf("error saving original2: %v", err)
	}

	if short2 == "" {
		t.Fatalf("expected short2 to not be empty, got %q", short2)
	}

	if fs.Get(short2) != original2 {
		t.Fatalf("cache[short2] is not original2")
	}

	if err := fs.Close(); err != nil {
		t.Fatalf("error closing file: %v", err)
	}

	file2, err := os.Open(filename)
	if err != nil {
		t.Fatalf("error opening test file: %v", err)
	}
	defer file2.Close()

	fs2, err := NewFs(file2)
	if err != nil {
		t.Fatalf("error creating file system: %v", err)
	}

	for short, original := range fs.cache {
		if fs2.Get(short) != original {
			t.Errorf("cache[%s] is not %s", short, original)
		}
	}
}

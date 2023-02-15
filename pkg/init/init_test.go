package init

import (
	"os"
	"testing"
)

func TestInitBook(t *testing.T) {
	dir, _ := os.Getwd()

	InitBook("中国", true, "git", dir)
}

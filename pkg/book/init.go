package book

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/TaceyWong/mdbook/pkg/utils"
	"github.com/spf13/viper"
)

func InitBook(title string, isVCS bool, vsc string, root string) error {
	if err := utils.CreateDirIfNotExist(root); err != nil {
		return nil
	}
	bookPath := filepath.Join(root, "book")
	if err := utils.CreateDirIfNotExist(bookPath); err != nil {
		return nil
	}
	srcPath := filepath.Join(root, "src")
	if err := utils.CreateDirIfNotExist(srcPath); err != nil {
		return err
	} else {
		ioutil.WriteFile(filepath.Join(srcPath, "chapter_1.md"),
			[]byte("# Chapter 1"), os.ModePerm)
		ioutil.WriteFile(filepath.Join(srcPath, "SUMMARY.md"),
			[]byte("# Summary\n\n- [Chapter 1](./chapter_1.md)"), os.ModePerm)
	}
	if isVCS {
		vcsIgnore := ""
		switch vsc {
		case "git":
			vcsIgnore = ".gitignore"
		default:
			vcsIgnore = ".gitignore"
		}
		f, err := os.Create(filepath.Join(root, vcsIgnore))
		if err != nil {
			return err
		}
		defer f.Close()
	}
	v := viper.New()
	u, _ := user.Current()
	v.Set("book.authors", []string{u.Username})
	v.Set("book.language", "en")
	v.Set("book.multilingual", false)
	v.Set("book.src", "src")
	v.Set("book.title", title)

	confgiPath := filepath.Join(root, "book.toml")
	fmt.Println(v.WriteConfigAs(confgiPath))
	return nil
}

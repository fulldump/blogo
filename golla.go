package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var update bool
var parallel bool

func init() {
	flag.BoolVar(&update, "update", false,
		"Force download all dependencies again")
	flag.BoolVar(&parallel, "parallel", true,
		"Download all dependencies in parallel")

	flag.Parse()
}

func main() {

	clones := sync.WaitGroup{}

	ForEachLine("./golla", func(line string) (err error) {

		// Clean blanks
		line = strings.TrimSpace(line)

		// Remove blank lines
		if "" == line {
			return
		}

		// Remove comments
		if strings.HasPrefix(line, "#") {
			return
		}

		parts := strings.SplitN(line, "->", 2)
		if len(parts) != 2 {
			return errors.New("Missing destination path after -> ")
		}

		repo := strings.TrimSpace(parts[0])
		dest := strings.TrimSpace(parts[1])

		if parallel {
			clones.Add(1)
			go func(repo, dest string) {
				defer clones.Done()
				if err = Clone(repo, dest); nil != err {
					fmt.Printf("Error cloning %s: %s\n", repo, err)
					return
				}
			}(repo, dest)
			return
		}

		if err = Clone(repo, dest); nil != err {
			return
		}

		return
	})

	clones.Wait()

}

func Clone(repo, dest string) (err error) {

	if !update && FileExists(dest) {
		fmt.Println("Repo already cloned:", repo)
		return
	}

	fmt.Println("Cloning repo:", repo, "to:", dest)

	tempDir, err := ioutil.TempDir("", "golla")
	if nil != err {
		return
	}
	defer os.RemoveAll(tempDir)

	branch := ""

	repo_parts := strings.SplitN(repo, "#", 2)
	if len(repo_parts) == 2 {
		repo = repo_parts[0]
		branch = repo_parts[1]
	}

	subpath := ""
	subpath_parts := strings.SplitN(repo, ">", 2)
	if len(subpath_parts) == 2 {
		repo = subpath_parts[0]
		subpath = subpath_parts[1]
	}

	// Find git binary
	binary, lookErr := exec.LookPath("git")
	if lookErr != nil {
		panic(lookErr)
	}

	// Compose git command
	args := []string{
		"clone",
		repo,
		tempDir,
	}

	if branch != "" {
		args = append(args, "--branch", branch)
	}

	// Launch git command
	gitCmd := exec.Command(binary, args...)
	gitCmd.Stdin = os.Stdin
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr

	err = gitCmd.Start()
	if nil != err {
		return
	}

	if err = gitCmd.Wait(); nil != err {
		return
	}

	// Remove .git:
	os.RemoveAll(path.Join(tempDir, ".git"))

	if update {
		os.RemoveAll(dest)
	}

	// Copy files:
	if subpath == "" {
		files, e := ioutil.ReadDir(tempDir)
		if nil != e {
			return e
		}
		for _, file := range files {
			err = Copy(
				filepath.Join(tempDir, file.Name()),
				filepath.Join(dest, file.Name()),
			)
			if nil != err {
				return
			}
		}
	} else {
		// Or only one file
		err = Copy(
			filepath.Join(tempDir, subpath),
			dest,
		)
	}

	return
}

func FileExists(dir string) bool {

	_, err := os.Stat(dir)
	if nil != err {
		return false
	}

	return true
}

func ForEachLine(filename string, onReadLine func(line string) (err error)) {

	f, err := os.Open(filename)
	if nil != err {
		panic(err)
	}
	defer f.Close()

	i := -1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i++
		if err := onReadLine(scanner.Text()); nil != err {
			fmt.Printf("Error on line %d: %s\n", i, err)
		}
	}
}

// Copy copies src to dest, doesn't matter if src is a directory or a file
func Copy(src, dest string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	return copy(src, dest, info)
}

// "info" must be given here, NOT nil.
func copy(src, dest string, info os.FileInfo) error {
	if info.IsDir() {
		return dcopy(src, dest, info)
	}
	return fcopy(src, dest, info)
}

func fcopy(src, dest string, info os.FileInfo) error {

	if err := os.MkdirAll(path.Dir(dest), os.ModeDir|os.ModePerm); nil != err {
		return err
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = os.Chmod(f.Name(), info.Mode()); err != nil {
		return err
	}

	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = io.Copy(f, s)
	return err
}

func dcopy(src, dest string, info os.FileInfo) error {

	if err := os.MkdirAll(dest, info.Mode()); err != nil {
		return err
	}

	infos, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, info := range infos {
		if err := copy(
			filepath.Join(src, info.Name()),
			filepath.Join(dest, info.Name()),
			info,
		); err != nil {
			return err
		}
	}

	return nil
}

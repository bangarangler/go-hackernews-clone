// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	goDotEnvVar "github.com/bangarangler/go-hackernews-clone/internal/pkg/db/pg"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	// cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return sh.Run("go", "build", "-o", "go_hackernews_clone", ".")
}

func StartDocker() error {
	fmt.Println("Mage starting postgres container...")
	return sh.Run("docker-compose", "up", "-d")
	// return cmd.Run()
}

// Stop Running Postgres Docker Container.  under the hood running docker
// compose down
func StopDocker() error {
	fmt.Println("Mage stoping postgres container...")
	return sh.Run("docker", "compose", "down")
	// return cmd.Run()
}

func MigratePG() error {
	fmt.Println("Mage ... go-migrate running migrations...")
	test := goDotEnvVar("POSTGRES_URL")
	fmt.Println("test", test)
	// return sh.Run("migrate", "-database", goDotEnvVar("POSTGRES_URL"), "-path", "internal/pkg/db/migrations/pg", "up")
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./go_hackernews_clone", "/usr/bin/go_hackernews_clone")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	sh.Run("go", "mod", "tidy")
	cmd := exec.Command("go", "get", "github.com/bangarangler/go_hackernews_clone")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}

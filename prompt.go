package main

import (
	"bufio"
	"fmt"
	ehmgr "github.com/EndlessHosting/ehmgr/pkg"
	"github.com/sethvargo/go-password/password"
	"os"
	"strings"
)
import "github.com/antham/strumt/v2"

func RunPrompt() DAUser {
	user := DAUser{}
	genPass, _ := password.Generate(20, 5, 5, false, false)
	p := strumt.NewPromptsFromReaderAndWriter(bufio.NewReader(os.Stdin), os.Stdout)
	p.AddLinePrompter(&StringPrompt{&user.Username, "Account Username:", "username", "email", "username"})
	p.AddLinePrompter(&StringPrompt{&user.Email, "Account Email:", "email", "password", "email"})
	p.AddLinePrompter(&StringPromptDefault{&user.Password, genPass, "Account Password:", "password", "package", "password"})
	pkgs, err := DA.ListPackages()
	if err != nil {
		panic("Unable to load packages: " + err.Error())
	}
	p.AddLinePrompter(&PackagePrompt{&user.Package, pkgs, "Account Package:", "package", "", "package"})
	p.SetFirst("username")

	p.Run()

	return user
}

type StringPrompt struct {
	store *string
	prompt string
	currentID string
	nextPrompt string
	nextPromptOnError string
}

func (s *StringPrompt) ID() string {
	return s.currentID
}

func (s *StringPrompt) PromptString() string {
	return s.prompt
}

func (s *StringPrompt) Parse(value string) error {
	if value == "" {
		return fmt.Errorf("Please give a non-empty value")
	}

	*(s.store) = value

	return nil
}

func (s *StringPrompt) NextOnSuccess(value string) string {
	return s.nextPrompt
}

func (s *StringPrompt) NextOnError(err error) string {
	return s.nextPromptOnError
}

type StringPromptDefault struct {
	store *string
	def string
	prompt string
	currentID string
	nextPrompt string
	nextPromptOnError string
}

func (s *StringPromptDefault) ID() string {
	return s.currentID
}

func (s *StringPromptDefault) PromptString() string {
	return fmt.Sprintf("%s (%s)", s.prompt, s.def)
}

func (s *StringPromptDefault) Parse(value string) error {
	if strings.TrimSpace(value) == "" {
		*(s.store) = s.def
		return nil
	}

	*(s.store) = value

	return nil
}

func (s *StringPromptDefault) NextOnSuccess(value string) string {
	return s.nextPrompt
}

func (s *StringPromptDefault) NextOnError(err error) string {
	return s.nextPromptOnError
}

type PackagePrompt struct {
	store *string
	packages ehmgr.PackageList
	prompt string
	currentID string
	nextPrompt string
	nextPromptOnError string
}

func (s *PackagePrompt) ID() string {
	return s.currentID
}

func (s *PackagePrompt) PromptString() string {
	return fmt.Sprintf("%s %+q", s.prompt, s.packages)
}

func (s *PackagePrompt) Parse(value string) error {
	if value == "" {
		return fmt.Errorf("Please give a non-empty value")
	}

	if !contains(s.packages, value) {
		return fmt.Errorf("Package not found")
	}

	*(s.store) = value

	return nil
}

func (s *PackagePrompt) NextOnSuccess(value string) string {
	return s.nextPrompt
}

func (s *PackagePrompt) NextOnError(err error) string {
	return s.nextPromptOnError
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
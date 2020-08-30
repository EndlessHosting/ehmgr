package main

import (
	ehmgr "ehmgr/pkg"

	"github.com/jawher/mow.cli"
	"github.com/sethvargo/go-password/password"

	"bytes"
	"fmt"
	"os"
	"text/template"
)

const (
	version = "v0.1.0"
	descrip = "An interface to DAProxy to assist in creating accounts"
)

var DA *ehmgr.Client

func main() {
	// Create an API client
	if loadConfig() {
		DA = ehmgr.NewClient(k.String("endpoint"), k.String("key"))

		// Create a cli
		app := cli.App("ehmgr", descrip)
		app.Version("v version", version)

		// ehmgr new
		app.Command("new", "Create a new account", newCmd)

		// ehmgr prompt
		app.Command("prompt", "Run the prompt", promptCmd)

		// ehmgr pkgs
		app.Command("pkgs", "List available packages", pkgsCmd)

		app.Before = func() {
			fmt.Println(CreateCustomBanner(descrip, version))
		}

		app.Run(os.Args)
	}
}

func newCmd(cmd *cli.Cmd) {
	cmd.Spec = "-u=<username> -e=<email> [-i=<ip> | -p=<password> | --package=<package>]"
	var (
		username = cmd.StringOpt("u username", "", "Account username")
		email    = cmd.StringOpt("e email", "", "Account email")
		ip       = cmd.StringOpt("i ip", k.String("ip"), "Public IP")
		genPass,_= password.Generate(20, 5, 5, false, false)
		pass 	 = cmd.StringOpt("p password", genPass, "Account password")
		pack     = cmd.StringOpt("package", k.String("package"), "Account package")
	)
	cmd.Action = func() {
		mTmpl, err := template.New("message").Parse(ehmgr.MessageTemplate())
		if err != nil {
			fmt.Printf("Error parsing message generation template\n%v\n", err)
			os.Exit(1)
		}
		user := &DAUser{
			Username: *username,
			Email:    *email,
			Password: *pass,
			Package:  *pack,
		}

		nu := ehmgr.CreateNewUser(user.Username, user.Email, user.Password, user.Package, k.String("domain"), *ip)
		resp, err := DA.CreateUser(nu)
		if err != nil {
			panic(err)
		}
		if !resp.Error {
			fmt.Println("Account created successfully!")
			msg := bytes.NewBufferString("")
			err = mTmpl.Execute(msg, nu)
			if err != nil {
				fmt.Println("Could not create message: ", err)
				return
			}
			fmt.Println("BEGIN MESSAGE:")
			fmt.Println()
			fmt.Println(msg.String())
			fmt.Println()
			fmt.Println("END MESSAGE")
		} else {
			fmt.Println("Unable to create account: ")
			fmt.Println(resp.Details)
		}
	}
}

func promptCmd(cmd *cli.Cmd) {

	cmd.Action = func () {
		mTmpl, err := template.New("message").Parse(ehmgr.MessageTemplate())
		if err != nil {
			fmt.Printf("Error parsing message generation template\n%v\n", err)
			os.Exit(1)
		}
		user := RunPrompt()

		nu := ehmgr.CreateNewUser(user.Username, user.Email, user.Password, user.Package, k.String("domain"), k.String("ip"))
		resp, err := DA.CreateUser(nu)
		if err != nil {
			panic(err)
		}

		if !resp.Error {
			fmt.Println("Account created successfully!")
			msg := bytes.NewBufferString("")
			err = mTmpl.Execute(msg, nu)
			if err != nil {
				fmt.Println("Could not create message: ", err)
				return
			}
			fmt.Println("BEGIN MESSAGE:")
			fmt.Println()
			fmt.Println(msg.String())
			fmt.Println()
			fmt.Println("END MESSAGE")
		} else {
			fmt.Println("Unable to create account: ")
			fmt.Println(resp.Details)
		}
	}
}

func pkgsCmd(cmd *cli.Cmd) {
	cmd.Action = func () {
		pkgs, err := DA.ListPackages()
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nPackages:\n%+q\n", pkgs)
	}
}
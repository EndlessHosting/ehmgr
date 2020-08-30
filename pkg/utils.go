package ehmgr

import "strings"

func MessageTemplate() string {
	var tmpl = strings.Builder{}
	tmpl.WriteString("Hello,\n\n")

	tmpl.WriteString("I'm pleased to say that I have created your account! ")
	tmpl.WriteString("Your credentials to log into our panel are below. ")
	tmpl.WriteString("Please note that your account includes a subdomain of our services ")
	tmpl.WriteString("({{.Domain}}). If you want to add your own domain, you ")
	tmpl.WriteString("will need to do it under Domain Setup in DirectAdmin and point ")
	tmpl.WriteString("it to our servers. Our website explains how to do this. ")
	tmpl.WriteString("If you ever need help, our support staff can be reached ")
	tmpl.WriteString("at https://support.theendlessweb.com\n\n")

	tmpl.WriteString("Panel URL: https://da.theendlessweb.com:2222/\n")
	tmpl.WriteString("Username: {{.Username}}\n")
	tmpl.WriteString("Password: {{.Password}}\n\n")

	tmpl.WriteString("Please note it is HEAVILY RECOMMENDED you change your password!\n\n")

	tmpl.WriteString("Remember to also join our Discord server! https://discord.gg/jVa63vC\n")
	tmpl.WriteString("We also have an e2e encrypted Rocket.Chat instance for those who wish to use it: https://chat.itsendless.org\n\n")

	tmpl.WriteString("On behalf of all Endless Hosting staff, welcome!\n")
	tmpl.WriteString("- EH Signups Team")

	return tmpl.String()
}
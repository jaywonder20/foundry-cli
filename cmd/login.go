package cmd

import (
	// "bytes"
	// "encoding/json"
	// "errors"
	// "fmt"
	// "io/ioutil"
	// "net/http"

	"context"
	"log"
	"os"

	"foundry/cli/auth"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

var (
	loginCmd = &cobra.Command{
		Use: 		"login",
		Short: 	"Log to your Foundry account",
		Long:		"",
		Run:		runLogin,
	}

	qs = []*survey.Question{
			{
					Name:     "email",
					Prompt:   &survey.Input{Message: "Email:"},
					Validate: survey.Required,
			},
			{
					Name: "pass",
					Prompt:   &survey.Password{Message: "Password:"},
					Validate: survey.Required,
			},
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

func runLogin(cmd *cobra.Command, args []string) {
	creds := struct {
		Email string 	`survey:"email`
		Pass	string	`survey:"pass`
	}{}
	err := survey.Ask(qs, &creds)
	// Without this specific "if" SIGINT (Ctrl+C) would only
	// interrupt the survey's prompt and not the whole program
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else if err != nil {
		log.Println(err)
	}

	a := auth.New()
	if err = a.SignIn(context.Background(), creds.Email, creds.Pass); err != nil {
		color.Red("⨯ Error")
		log.Println("HTTP request error", err)
		return
	}

	if a.Error != nil {
		color.Red("⨯ Error")
		log.Println("Auth error", a.Error)
		return
	}

	if err = a.SaveTokens(); err != nil {
		color.Red("⨯ Error")
		log.Println("Save tokens error", err)
		return
	}

	color.Green("✔ Signed in")
}

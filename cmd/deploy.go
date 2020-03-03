// Copyright © 2020 tdonovic
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var Repo string
var Ref string
var AutoMerge bool
var Environment string
var User string
var Owner string
var Token string

// deploy represents the deploy command
var deploy = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		task := "deploy"
		client := github.NewClient(tc)
		emptyList := []string{}
		depReq := github.DeploymentRequest{
			Ref:              &Ref,
			Task:             &task,
			AutoMerge:        &AutoMerge,
			RequiredContexts: &emptyList,
			Environment:      &Environment,
		}

		// list all repositories for the authenticated user
		_, response, err := client.Repositories.CreateDeployment(ctx, User, Repo, &depReq)
		fmt.Println(response.Response.StatusCode)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(deploy)
	deploy.Flags().StringVarP(&Repo, "repo", "r", "", "Repo to deploy from")
	deploy.Flags().StringVarP(&Ref, "ref", "", "", "Ref to deploy")
	deploy.Flags().StringVarP(&Environment, "environment", "", "dev", "Enviornment to deploy to")
	deploy.Flags().BoolVarP(&AutoMerge, "auto-merge", "", false, "Automerge after commit?")
	deploy.Flags().StringVarP(&User, "user-id", "", "", "Your Github username")
	deploy.Flags().StringVarP(&Token, "token", "", "", "Personal Access token")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploy.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploy.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

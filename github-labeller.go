package main

import (
    // core
    "context"
    "flag"
    "fmt"
    "os"
    "strings"

    // 3rd party
    "golang.org/x/oauth2"
    "github.com/BurntSushi/toml"
    "github.com/google/go-github/github"
)

const VERSION = "UNRELEASED"

const helpString = `

Usage:

github-labeller create <label name> <color>
github-labeller delete <label name>

For proper operation of the tool, its best if the label name and color
are quoted when provided. For example:
    github-labeller create "foo label" "fff"

You can create an API token for use with this app going to
https://github.com/settings/tokens, logging in, and creating a new
"Personal Access Token" with the "repo" permission.

`
const tokenError = `

No GitHub API token found. Token must be provided either via -token=<api token>
or via 'token=your api token' in the $HOME/.github-labeller config file.

`

func usage() {
    fmt.Printf(helpString)
    flag.PrintDefaults()
    fmt.Printf("\n")
    os.Exit(0)
}

type organization struct {
    Repositories []string
}

type Config struct {
    Token string
    Orgs map[string]organization
}

func createLabel(client *github.Client, orgName string, repo string, label *github.Label) {
    ctx := context.Background()
    _, _, err := client.Issues.GetLabel(ctx, orgName, repo, *label.Name)
    // If label doesn't exist, create it
    if err != nil && strings.Contains(err.Error(), "404") {
        newLabel, _, err := client.Issues.CreateLabel(ctx, orgName, repo, label)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
        } else {
            fmt.Printf("CREATED: %s/%s#%s\n", orgName, repo, *newLabel.Name)
        }
    // If label does exist, update it
    } else {
        newLabel, _, err := client.Issues.EditLabel(ctx, orgName, repo, *label.Name, label)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
        } else {
            fmt.Printf("UPDATED: %s/%s#%s\n", orgName, repo, *newLabel.Name)
        }
    }
}

func deleteLabel(client *github.Client, orgName string, repo string, labelName string) {
    ctx := context.Background()
    _, err := client.Issues.DeleteLabel(ctx, orgName, repo, labelName)
    if err != nil && !strings.Contains(err.Error(), "404") {
        fmt.Printf("Error: %s\n", err)
    } else {
        fmt.Printf("DELETED: %s/%s#%s\n", orgName, repo, labelName)
    }
}

func main() {
    var (
        isHelp bool
        isVersion bool
        tokenFlag string
    )

    flag.BoolVar(&isHelp, "help", false, "Print help string and exit")
    flag.BoolVar(&isVersion, "version", false, "Print version string and exit")
    flag.StringVar(&tokenFlag, "token", "", "GitHub API token to authenticate with -- overrides token set via config")
    flag.Parse()

    if isVersion {
        fmt.Printf("%s\n", VERSION)
        os.Exit(0)
    }

    if isHelp || flag.NArg() < 3 {
        usage()
    }

    operation := flag.Arg(0)
    labelName := flag.Arg(1)
    labelColor := flag.Arg(2)

    var config Config
    // TODO: Make home dir query cross-platform
    filename := os.Getenv("HOME") + "/.github-labeller"
    fmt.Printf("\nReading config from %s...\n\n", filename)
    if _, err := toml.DecodeFile(filename, &config); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    apiToken := ""
    if tokenFlag != "" {
        apiToken = tokenFlag
    } else if config.Token != "" {
        apiToken = config.Token
    } else {
        fmt.Printf(tokenError)
        os.Exit(1)
    }

    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: apiToken},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)
    client := github.NewClient(tc)

    label := &github.Label{URL: nil, Name: &labelName, Color: &labelColor}

    for orgName, org := range config.Orgs {
        for _, repo := range org.Repositories {
            if operation == "create" {
                createLabel(client, orgName, repo, label)
            } else if operation == "delete" {
                deleteLabel(client, orgName, repo, *label.Name)
            }
        }
    }
}

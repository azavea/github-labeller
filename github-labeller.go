package main

import (
    // core
    "flag"
    "fmt"
    "os"

    // 3rd party
    "golang.org/x/oauth2"
    "github.com/BurntSushi/toml"
    "github.com/google/go-github/github"
)

const helpString = `

Usage: github-labeller <options> <label name> <color>

Must provide one of either -create or -delete, otherwise this tool is a noop.

For proper operation of the tool, its best if the label name and color
are quoted. For example:
    github-labeller -create "foo label" "#fff"

You can create an API token for use with this app by logging into GitHub with
your user and going to https://github.com/settings/tokens, and creating a new
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
}

type Config struct {
    Token string
    Repos []string
}

func main() {
    var (
        isHelp bool
        isCreate bool
        isDelete bool
        tokenFlag string
    )

    flag.BoolVar(&isHelp, "help", false, "Print help string and exit")
    flag.BoolVar(&isCreate, "create", false, "If provided, create the specified label")
    flag.BoolVar(&isDelete, "delete", false, "If provided, delete the specified label")
    flag.StringVar(&tokenFlag, "token", "", "GitHub API token to authenticate with -- overrides token set via config")
    flag.Parse()

    if isHelp || flag.NArg() != 2 {
        usage()
        os.Exit(0)
    }

    var config Config
    // TODO: Make home dir query cross-platform
    var filename = os.Getenv("HOME") + "/.github-labeller"
    fmt.Printf("Reading config from: %s\n", filename)
    if _, err := toml.DecodeFile(filename, &config); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    var labelName = flag.Arg(0)
    var labelColor = flag.Arg(1)
    var apiToken = ""
    if tokenFlag != "" {
        apiToken = tokenFlag
    } else if config.Token != "" {
        apiToken = config.Token
    } else {
        fmt.Printf(tokenError)
        os.Exit(1)
    }

    fmt.Printf("Label: %s, Color: %s\n", labelName, labelColor)
    fmt.Printf("Token: %s\n", apiToken)
    fmt.Printf("Repos: %s\n", config.Repos)

    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: apiToken},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)

    client := github.NewClient(tc)

    // TODO: Update config so that it allows multiple orgs
    labels, _, err := client.Issues.ListLabels("azavea", config.Repos[0], nil)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("Labels: %s\n", labels)
}

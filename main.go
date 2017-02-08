package main;


import (
	"os"
	"fmt"
	"github.com/mkideal/cli"
	"gopkg.in/yaml.v2"
	DeployCli "./src"
	"path/filepath"
	"io/ioutil"
)

var (
	apiEndpoint string
 	apiEmail string
 	apiPass string
	version string = "0.0.1"
)


type Config struct {
	Name string
	Method string
	Vars string
}


func main() {

	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(runit),
		cli.Tree(plugins),
		cli.Tree(thisversion),
		cli.Tree(updater),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	filename, _ := filepath.Abs("examples/docker.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

}

var help = cli.HelpCommand("show all the useful commands")

// root command
type rootT struct {
	cli.Helper
	//Version  bool `cli:"v,version" usage:"get this version of storj-go"`
}

var root = &cli.Command{
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		//argv := ctx.Argv().(*rootT)
		ctx.String("\n==========  "+(ctx.Color().Blue("deploy"))+"  ==========\n")
		ctx.String("================================\n")
		ctx.String("Try this out: '" + (ctx.Color().Magenta("deploy help")) + "'\n\n")
		return nil
	},
}


type argT struct {
	cli.Helper
	Username string `cli:"u,username" usage:"storj account" prompt:"Email"`
	Password string `pw:"p,password" usage:"password to storj account" prompt:"Password"`
}

var runit = &cli.Command{
	Name: "run",
	Desc: "run a deployment config file",
	Fn: func(ctx *cli.Context) error {
		cli.Run(new(argT), func(ctx *cli.Context) error {
			argv := ctx.Argv().(*argT)
			DeployCli.ApiEmail = argv.Username;
			DeployCli.ApiPass = DeployCli.EncryptPassword(argv.Password);
			DeployCli.ApiEndpoint = "https://api.storj.io/";
			response := DeployCli.SetAPIInfo(apiEmail,apiPass, apiEndpoint)
			ctx.String((ctx.Color().Yellow(response+"\n")))
			return nil
		})
		return nil
	},
}


var plugins = &cli.Command{
	Name: "plugins",
	Desc: "view all plugins installed",
	Fn: func(ctx *cli.Context) error {
		ctx.String((ctx.Color().Yellow("Fetching your plugins...\n")))
		userEmail := DeployCli.GetUser()
		ctx.String(userEmail)
		ctx.String(DeployCli.ApiEmail)
		//ctx.String(buckets)
		return nil
	},
}


var thisversion = &cli.Command{
	Name: "version",
	Desc: "current deploy version",
	Fn: func(ctx *cli.Context) error {
		ctx.String("deploy Version: " + (ctx.Color().Red(version)) + "\n")
		return nil
	},
}


var updater = &cli.Command{
	Name: "update",
	Desc: "check for an update",
	Fn: func(ctx *cli.Context) error {
		ctx.String("Checking for updates...\n")
		newVersion := DeployCli.GetNewVersion();
		if newVersion != version {
			ctx.String("Newest Version: " + (ctx.Color().Red(newVersion)) + "\n")
			ctx.String("Using Version:  " + (ctx.Color().Red(version)) + "\n")
			ctx.String(ctx.Color().Red("You should do an update!\n"))
			ctx.String("Run command: \n" + (ctx.Color().White("curl https://raw.githubusercontent.com/deploymentcli/deploy/master/install.sh | bash")) + "\n")
		} else {
			ctx.String("Newest Version: " + (ctx.Color().Green(newVersion)) + "\n")
			ctx.String("Using Version:  " + (ctx.Color().Green(version)) + "\n")
			ctx.String(ctx.Color().Green("deploy is up to date!\n"))
		}

		return nil
	},
}
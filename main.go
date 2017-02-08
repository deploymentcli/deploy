package main;


import (
	"os"
	"fmt"
	"github.com/mkideal/cli"
	DeployCli "./src"
)

var (
	apiEndpoint string
 	apiEmail string
 	apiPass string
	version string = "0.0.1"
)



func main() {

	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(login),
		cli.Tree(buckets),
		cli.Tree(thisversion),
		cli.Tree(updater),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
		ctx.String("\n==========  "+(ctx.Color().Blue("storj-go"))+"  ==========\n")
		ctx.String("======= "+(ctx.Color().Yellow("API for Storj.io"))+" =======\n")
		ctx.String("================================\n")
		ctx.String("Try this out: '" + (ctx.Color().Magenta("storj-go help")) + "'\n\n")
		return nil
	},
}


type argT struct {
	cli.Helper
	Username string `cli:"u,username" usage:"storj account" prompt:"Storj Email"`
	Password string `pw:"p,password" usage:"password to storj account" prompt:"Storj Password"`
}

var login = &cli.Command{
	Name: "login",
	Desc: "login to Storj.io for access",
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


var buckets = &cli.Command{
	Name: "buckets",
	Desc: "view all of my buckets",
	Fn: func(ctx *cli.Context) error {
		ctx.String((ctx.Color().Yellow("Fetching your Buckets...\n")))
		userEmail := DeployCli.GetUser()
		ctx.String(userEmail)
		ctx.String(DeployCli.ApiEmail)
		//ctx.String(buckets)
		return nil
	},
}


var thisversion = &cli.Command{
	Name: "version",
	Desc: "current storj-go version",
	Fn: func(ctx *cli.Context) error {
		ctx.String("storj-go Version: " + (ctx.Color().Red(version)) + "\n")
		ctx.String("Check for an update: '" + (ctx.Color().Magenta("storj-go update")) + "'\n")
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
			ctx.String("Run command: \n" + (ctx.Color().White("curl https://raw.githubusercontent.com/hunterlong/storj-go/master/install.sh | bash")) + "\n")
		} else {
			ctx.String("Newest Version: " + (ctx.Color().Green(newVersion)) + "\n")
			ctx.String("Using Version:  " + (ctx.Color().Green(version)) + "\n")
			ctx.String(ctx.Color().Green("storj-go is up to date!\n"))
		}

		return nil
	},
}
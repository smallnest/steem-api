package main

import (
	"github.com/kr/pretty"

	// Stdlib
	"flag"
	"log"

	// RPC
	client "github.com/smallnest/steem-api"

	// Vendor
	"github.com/pkg/errors"
)

// private key to wif by cli_wallet:
//  locked >>> get_private_key_from_password username active "password"

var (
	voter = "xeroc"
	key   = "5JLw5dgQAx6rhZEgNN5C2ds1V47RweGshynFSWFbaMohsYsBvE8"
)

func main() {
	cls, err := client.NewClient([]string{"wss://gtg.steem.house:8090"}, "steem")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer cls.Close()

	cls.SetKeys(&client.Keys{PKey: []string{key}})

	if err := run(cls); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run(cls *client.Client) (err error) {
	flag.Parse()
	// Process args.
	args := flag.Args()

	if len(args) != 2 {
		return errors.New("2 arguments required")
	}
	author, permlink := args[0], args[1]

	resp, err := cls.Vote(voter, author, permlink, 10000)
	if err != nil {
		return
	}
	log.Println(pretty.Sprint(resp))

	return nil
}

package cmd

import (
    "bufio"
    "fmt"
	"github.com/springeye/note-server/db"
	"github.com/urfave/cli/v2"
    "log"
    "os"
    "strings"
)
func askForConfirmation(s string) bool {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Printf("%s [y/n]: ", s)

        response, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        response = strings.ToLower(strings.TrimSpace(response))

        if response == "y" || response == "yes" {
            return true
        } else if response == "n" || response == "no" {
            return false
        }
    }
}

func AddUser(c *cli.Context) error {
	if c.Args().Len() != 2 {
		cli.ShowSubcommandHelpAndExit(c, 1)
	}
	username := c.Args().Get(0)
	password := c.Args().Get(1)
	user := db.User{
		Username: username,
		Password: password,
	}
	db.Setup()
	var count int64
	db.Connection.Model(&user).Where("username = ?", username).Count(&count)
	if count > 0 {
		return cli.Exit(fmt.Sprintf("User %s already exists\n", username), 1)
	}
	err := db.Connection.Create(&user).Error
	if err != nil {
		return cli.Exit(fmt.Sprintf("create user error:%s\n", err.Error()), 1)
	}
	return nil
}
func DeleteUser(c *cli.Context) error {
	if c.Args().Len() != 1 {
		cli.ShowSubcommandHelpAndExit(c, 1)
	}
    username:=c.Args().Get(0)
    if askForConfirmation(fmt.Sprintf("Are you sure delete user [%s]",username)){
        db.Setup()
        return db.Connection.Where("username = ?",c.Args().Get(0)).Delete(&db.User{}).Error
    }else{
        println("cancel")
    }
	return nil
}
func SetPassword(c *cli.Context) error {

	return nil
}

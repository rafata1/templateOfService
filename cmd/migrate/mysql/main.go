package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "github.com/spf13/cobra"
    "log"
    "os"
)

func migrateUpCmd() *cobra.Command {
    return &cobra.Command{
        Use: "up",
        Run: func(cmd *cobra.Command, args []string) {
            dsn := fmt.Sprintf("mysql://%s", os.Getenv("MYSQL"))
            fmt.Println("up up up", dsn)
        }}
}

func migrateDownCmd() *cobra.Command {
    return &cobra.Command{
        Use: "down",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("down down down")
        },
        //TODO: implement down migration
    }
}

func migrateCreateCmd() *cobra.Command {
    return &cobra.Command{
        Use: "create",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println(args[0])
            fmt.Println("create create create")
            //TODO: implement create migration
        },
    }
}

func initCmd() *cobra.Command {
    cmd := &cobra.Command{}
    cmd.AddCommand(
        migrateUpCmd(),
        migrateDownCmd(),
        migrateCreateCmd())
    return cmd
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading env: %s", err.Error())
    }

    rootCmd := initCmd()
    err = rootCmd.Execute()
    if err != nil {
        log.Fatalf("Error migrating data %s", err.Error())
    }
}

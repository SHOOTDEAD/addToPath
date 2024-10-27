/*
Copyright © 2024 NAME HERE KALAISELVAN574DB@GMAIL.COM
*/
package cmd

import (
	"addToPath/logger"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "addToPath",
	Short: "A command-line tool for easily adding paths to your Linux user's .bashrc file.",
	Long: `This CLI tool simplifies the process of managing environment paths in your Linux user account by modifying the .bashrc file. 
	It allows users to append new paths quickly and safely, creating a backup of the original .bashrc before making any changes.
	If needed, the tool also provides the functionality to reverse the last change, restoring the previous state from the backup. 
	This utility is perfect for developers and users who frequently adjust their environment settings for various applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		SetPath()
	},
}

var path string
var reverse bool

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1) // Exit with a non-zero status on error
	}
}

// Logger is an instance of the custom logger
var Logger = logger.Logger

// SetPath handles the logic to append a new path to .bashrc
func SetPath() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		Logger.Error("Can't find path to user's home directory") // Log an error if unable to find home directory
		return
	}

	bashrcPath := homeDir + "/.bashrc"
	backupPath := homeDir + "/.bashrc.bak"

	if reverse {
		Undo(bashrcPath, backupPath) // Revert changes if the reverse flag is set
		return
	}

	if path == "" {
		Logger.Warn("Path parameter is required") // Warn if no path is provided
		return
	}

	Logger.Info("Attempting to back up .bashrc to:", backupPath)
	if _, err := os.Stat(bashrcPath); err == nil {
		// Backup the existing .bashrc file if it exists
		inputFile, err := os.Open(bashrcPath)
		if err != nil {
			Logger.Error("Error opening .bashrc for backup:", err) // Log error if opening fails
			return
		}
		defer inputFile.Close()

		outputFile, err := os.Create(backupPath)
		if err != nil {
			Logger.Error("Error creating backup file:", err) // Log error if creating backup fails
			return
		}
		defer outputFile.Close()

		if _, err := io.Copy(outputFile, inputFile); err != nil {
			Logger.Error("Error copying .bashrc to backup:", err) // Log error if copying fails
			return
		}

		Logger.Info("Backup created:", backupPath) // Confirm backup creation
	} else {
		Logger.Warn(".bashrc not found, creating a new one.") // Warn if .bashrc does not exist
	}

	// Open .bashrc to append the new path
	file, err := os.OpenFile(bashrcPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Logger.Error("Error opening .bashrc:", err) // Log error if opening .bashrc fails
		return
	}
	defer file.Close()

	Logger.Info("Writing to .bashrc:", path) // Log the path being added
	if _, err := file.WriteString(path + "\n"); err != nil {
		Logger.Error("Error writing to .bashrc:", err) // Log error if writing fails
	} else {
		Logger.Info("Path added to .bashrc successfully:", path, bashrcPath) // Confirm successful addition
	}
}

// Undo reverts changes to .bashrc using the backup file
func Undo(bashrcPath string, backupPath string) {
	Logger.Info("Attempting to reverse changes using backup:", backupPath)
	if _, err := os.Stat(backupPath); err == nil {
		// Restore from backup if it exists
		backupFile, err := os.Open(backupPath)
		if err != nil {
			Logger.Error("Error opening backup file:", err) // Log error if opening backup fails
			return
		}
		defer backupFile.Close()

		// Create or truncate the .bashrc file
		file, err := os.OpenFile(bashrcPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			Logger.Error("Error opening .bashrc for writing:", err) // Log error if opening .bashrc for writing fails
			return
		}
		defer file.Close()

		if _, err := io.Copy(file, backupFile); err != nil {
			Logger.Error("Error restoring from backup:", err) // Log error if restoring fails
			return
		}

		Logger.Info("Changes reversed: restored from", backupPath) // Confirm successful restoration
	} else {
		Logger.Warn("Backup not found, cannot reverse changes.") // Warn if backup is missing
	}
}

// init initializes command flags for the CLI
func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "", "Path to add to .bashrc")
	rootCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Revert the last change to .bashrc")
}

Go CLI Reminder Tool

This Go CLI Reminder Tool is a command-line application that allows users to set reminders based on specified times and display notifications.

Features

Parses time in hh:mm format using a natural language parser.
Displays desktop notifications for reminders.
Allows setting reminders to trigger at specified times.

To set a reminder, run the executable with the following command-line arguments:

go run main.go <hh:mm> <reminder_message>

Environment Variables

GOLANG_CLI_REMINDER: Set this environment variable to 1 to trigger reminders directly from the current instance of the tool.

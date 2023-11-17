> **note**
> The main purpose of this project is for me to learn Go, feedback is much appreciated!

# Gezondheid

Gezondheid `/ɣəˈzɔntˌɦɛi̯t/` (Dutch for "Health") is a simple CLI tool to periodically check the health of URLs.

👷‍♂️ This project (and readme) is under construction

# Usage

## Monitor
To monitor the configured endpoint run the following command: 
```shell
gezondheid monitor -u <url>
```

- -u <url>: Replace <url> with the URL of the endpoint you want to monitor. Make sure to provide the complete URL, including the protocol (e.g., http:// or https://).

This command allows you to actively monitor the health and status of an endpoint.

## Add endpoint to watch
To add a new endpoint for monitoring, use the following command: 
```shell
gezondheid add -n <name> -u <url>
```

- -n <name>: Replace <name> with a descriptive name for the endpoint you are adding. This name is used to identify the endpoint in your monitoring configuration.
- -u <url>: Replace <url> with the URL of the endpoint you want to monitor. Make sure to provide the complete URL, including the protocol (e.g., http:// or https://).

This command will add the specified endpoint to your existing settings.yaml file, or generate a new one if it doesn't exist.

## List an endpoint (WIP)
To list all configured endpoint for monitoring, use the following command: 
```shell
gezondheid list
```

not yet implemented.

## Remove an endpoint
To remove an endpoint for monitoring, use the following command: 
```shell
gezondheid remove -n <name>
```

- -n <name>: Replace <name> with a the name of the endpoint which configuration you want to remove.

This command allows you to easily eliminate endpoints that are no longer needed in your monitoring setup.

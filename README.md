# sshmenu written on GO

**sshmenu** is a simple tool for connecting to remote hosts via ssh written on GO.
Great if you have trouble remembering IP addresses, hostnames, usernames or path to a key file.

## Quick Setup
- `go build sshmenu.go`
- `chmod +x sshmenu`
- `./sshmenu`

## Configuration
Edit config.json file

        {
            "host": "192.168.0.1",
            "friendly":"Server name",
            "options": [""]
        },

### Options examples

| Option | Description                    |
| ------------- | ------------------------------ |
| `-lroot`      | username       |
| `-i~/.ssh/id_rsa`   | path to key file     |

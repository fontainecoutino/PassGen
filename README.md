# PassGen

Password Generator CLI is an application written in Go. Built to generate secure passwords based on users desires. It can generate multiple types of passwords; safe, strong, strongest, and unsafe. Custom length and character type can be specified as well. PassGen can also turn a word or phrase into a password.


## Installation

#### 1. Download Go
Make sure you have the latest version of [Go](https://go.dev/) installed. You can download latest the version [here](https://go.dev/doc/install). The version required is a least 1.19.1.

To check your version.
```bash
go version
```

#### 2. Build project
To build the project.
```bash
go build -o passgen main.go
```

#### Optional
Currently, the executable just lives in any specific directory you put it in. If you would like to make it an alias for your enviroment, here are some resources to help you with that: [Mac](https://superuser.com/questions/963784/command-line-for-application-alias-in-mac-os-x), [Linux](https://stackoverflow.com/questions/5137726/creating-permanent-executable-aliases).

<!-- [Windows](https://superuser.com/questions/560519/how-to-set-an-alias-in-windows-command-line) -->

## Usage
Run the application and get information for a more in-depth description.
```bash
./passgen
```

#### Get
Generate secure password.
```bash
./passgen get [-m mode] [-l length] [-t type]
```

#### Tokenize
Tokenizes a word. Converts letters into special characters.
```bash
./passgen tokenize [word] [flags]
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
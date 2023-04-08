# Table of contents
1. [Installation](#introduction)
2. [Project Motivation](#paragraph1)
3. [File Descriptions](#paragraph2)
4. [Interaction](#paragraph3)
5. [Author](#paragraph4)

### Installation <a name="introduction"></a>

**Libraries:**<br/>
- github.com/pelletier/go-toml

**Tools:**<br/>
- Go (version 1.16 or higher)

### Project Motivation <a name="paragraph1"></a>
This project is a simple Go application that reads a TOML configuration file, extracts the "postgres" section, and prints out the user and password fields.

### File Descriptions <a name="paragraph2"></a>
- `main.go`: The main source file containing the Go code that reads the configuration file and prints the user and password.

- `config.toml`: The TOML configuration file containing the postgres section with user and password.

### Interaction <a name="paragraph3"></a>
1. Create a `config.toml` file with the following format:

[postgres]
user = "your_username"<br/>
password = "your_password"

2. Run the Go program using the following command:<br/>
go run main.go

3. The output will display the user and password:<br/>
User: your_username<br/>
Password: your_password

### Author <a name="paragraph4"></a>
Sergio Lima, sergiosouzalima@gmail.com


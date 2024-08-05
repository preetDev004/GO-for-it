# Go Setup Information

## Introduction
Go, also known as Golang, is an open-source programming language designed for simplicity and efficiency. This file provides a step-by-step guide to setting up Go on your local environment.

## Prerequisites
- A computer running Windows, macOS, or Linux.
- A stable internet connection.

## Step 1: Download Go
1. Visit the official Go website: https://golang.org/dl/
2. Download the installer for your operating system:
   - **Windows**: `.msi` file
   - **macOS**: `.pkg` file
   - **Linux**: `.tar.gz` file

## Step 2: Install Go
### Windows
1. Run the downloaded `.msi` file.
2. Follow the installation prompts.
3. The installer will set the Go environment variables automatically.

### macOS
1. Open the downloaded `.pkg` file.
2. Follow the installation instructions.

### Linux
1. Open a terminal.
2. Extract the downloaded archive:
```bash
   tar -C /usr/local -xzf go1.x.linux-amd64.tar.gz
```

## (Optional) Step 3: Add Go to your PATH
1. adding the following line to your .bashrc or .bash_profile:
```bash
export PATH=$PATH:/usr/local/go/bin
```
2. Reload your profile:
```bash
source ~/.bashrc
```

## Step 4: Verify Installation
1. To verify that Go is installed correctly, open a terminal or command prompt and run:
```go
go version
```

## Step 5: Create a .new GO file
1. Open main.go in your preferred text editor and add the following code:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Step 6: Build & Run your Go Program
1. To build your program into an executable, run:
```bash
go build
```
2. This will create an executable file in the current directory. You can run it with:
```bash
./hello
```
# Dot

## Motivation

When working on multiple projects at the same time, environment management can become a real mess. Let's start with a real life example.

Let's say you're working on a side project, plus for a company on a daily basis. In the scope of this project, you need to deal with a custom environment, and therefore add new variables. You got two options there:
* add an `export` clause to your `.zshrc` file for each environment variable, carefully ensuring that you do not have any overlapping with any other env var from your company (which can be pretty harassing in case third-party libraries such as AWS clients source predefined env vars by default)
* prefix all your commands with `ENV_VAR=xxx` everytime you need to run anything

Otherwise, you can create one configuration file for each project, and source them accordingly. It might probably go well until you open/close hundreds of terminal panes.

## Workspaces

`dot` introduces the notion of workspaces. Dot workspaces are registries of environment variables and custom configuration files to source everytime you'll switch context. See it as `kubectl` for environments.

In practice, this is what `dot` proposes:

```sh
# First, create a new workspace
$ dot create workspace companyname
# Then switch your currently active workspace
$ dot use workspace companyname
# Attach environment variables
$ dot set AWS_ACCESS_KEY xxxxxxxxxxxxxxxxxxx
$ dot set AWS_SECRET_ACCESS_KEY xxxxxxxxxxxxxxxxxxx
# And attach a custom file to source to get your environment ready
$ dot add source ~/github.com/companyname/stack/.zshrc 
```

Then, if you start working on a new project, you can simply create a new workspace and attach new variables and files to it:

```sh
# First, create a new workspace
$ dot create workspace myproject
# Then switch your currently active workspace
$ dot use workspace myproject
# Attach environment variables
$ dot set AWS_ACCESS_KEY xxxxxxxxxxxxxxxxxxx
$ dot set AWS_SECRET_ACCESS_KEY xxxxxxxxxxxxxxxxxxx
```

What happens under the hood is that `dot` regenerates a `.dotrc` file everytime you perform any change to your active context. So changes take effect immediately, even when switching workspace.

## Sync accross multiple devices

`dotfiles` help a lot when it comes to synchronizing information accross multiple devices. A lot of people work with multiple machines ans to their best to keep things synced.

`dot` also aims to address this kind of concerns. Indeed, `dot` exposes a `sync` interface which allows you to:
* attach remote state stores
* sync your configuration accross all your devices, encrypted

Currently only a git provider has been implemented. Encryption leverages the AES algorithm, with a 32 characters long key.

Please find below an usage example:

```sh
# For git provider, initializes a git repository. Other providers such as S3 could create a bucket or whatever
$ dot sync init
# Add a new remote with name `origin` and url `git@github.com:user/repo.git`
$ dot sync add remote origin git@github.com:user/repo.git
# Push your changes
$ dot sync push -e <key>
```

Also note that files are being encrypted one after another and your remote state will match the same structure as your local one, expect that all files will be suffixed with `.encrypted`.

## Getting started

Install `dot` using Go:

```
$ go get -u https://github.com/emaincourt/dotstate
```

Then, run the following command to initialize `dot`:

```
$ dot init
```

For more information about initialization options, run:

```
λ  dot init --help
Initialises dot by creating the configuration folder and associated files

Usage:
  dot init [flags]

Flags:
  -h, --help   help for init

Global Flags:
  -c, --config string           Path to the dot configuration file (default "$HOME/.dot/dot.yaml")
  -e, --encryption-key string   The secret to use to encrypt/decrypt data on sync
      --log-level string        The verbosity of output: debug, info, warn, error (default "info")
  -n, --no-regenerate           Whether or not the rc file should be regenerated after any operation
  -f, --rc-file-path string     The path to the rc file to regenerate (default "$HOME/.dot/.dotrc")
```

## Documentation

This CLI was built on top of the amazing [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) libraries of [Steve Francia](https://github.com/spf13). So using `--help` from anywhere should be fine.

Please find below the root help of the `dot` binary:

```sh
λ  dot --help
Usage:
  dot [flags]
  dot [command]

Available Commands:
  add         Appends a new value to a list of existing values
  create      Creates a new resource
  delete      Deletes a resource
  get         List a type of resources
  help        Help about any command
  init        Initializes dot by creating the configuration folder and associated files
  set         Sets a new environment variable for current workspace
  sync        Syncs resources with a distant state store
  tidy        Cleans up workspaces
  unset       Unsets an environment variable for current workspace
  use         Switch use of a resource

Flags:
  -c, --config string           Path to the dot configuration file (default "$HOME/.dot/dot.yaml")
  -e, --encryption-key string   The secret to use to encrypt/decrypt data on sync
  -h, --help                    help for dot
      --log-level string        The verbosity of output: debug, info, warn, error (default "info")
  -n, --no-regenerate           Whether or not the rc file should be regenerated after any operation
  -f, --rc-file-path string     The path to the rc file to regenerate (default "$HOME/.dot/.dotrc")

Use "dot [command] --help" for more information about a command.
```
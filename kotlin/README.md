# Kotlin

### Installation

Install Command Line Compiler
`curl -s https://get.sdkman.io | bash`

Please open a new terminal, or run the following in the existing one:

`source "/Users/user/.sdkman/bin/sdkman-init.sh"`

Next, install kotlin

`sdk install kotlin`

Alternatively, install with homebrew

```bash
brew update

brew install kotlin
```

## Creating an Application

Compile the application using the Kotlin compiler (`.kt` file ext):

`kotlinc hello.kt -include-runtime -d hello.jar`

Run the application:

`java -jar hello.jar`

Compiling a library (exclude `-include-runtime`)

You can also use the kotlin script to run binaries produced by the Kotlin compiler:

`kotlin -classpath hello.jar HelloKt`

`HelloKt` is the main class name that the Kotlin compiler generates for the file named `hello.kt`.

#### Run a script

`kotlinc -script list_folders.kts -- -d <path_to_folder_to_inspect>`

#### Running the REPL

`bin/kotlinc-jvm`
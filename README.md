# Curricula

> Terminal tools for interacting with the Canvas LMS through the API

---

## Features

- Command Line based tool to interact with Canvas API endpoints
- Securly stores API Token using OS KeyRing or using ENV Variable (with option to pass via a cli flag)
- Outputs data to file as JSON or CSV or to console
- Save default values for flags in a configuration file
- Allow for using JSON files as POST body
- Geneate a log file of API interactions for each run

---

## Technologies Used

- [Go](https://golang.org/) – backend and CLI foundation using [Cobra](https://github.com/spf13/cobra)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) – (planned) TUI interface
- [Charmbracelet ecosystem](https://charm.sh/) – for rich terminal experiences
- [Canvas LMS API](https://canvas.instructure.com/doc/api/) – core data source

---

## Installation

## Usage

## Contributing

### Windows: Build Module

This project includes a PowerShell module to provide Make style build commands. This module will generate and set the VERSION, BUILD_NUMBER, BUILD_STATE, and BUILD_TIME variables for the application.
The module is included within the 'tools' directory of this repository.

To use the module, you will need to import it into your PowerShell session. You can do this by running the following command:

```powershell
Import-Module .\tools\Build.psm1
```

This module must be imported every time you open a new PowerShell session. There is no automatic importing of this module.
Once the module is imported, you will have access to the following commands:

- 'Set-Version' - Sets the version number for the application. This updates the VERSION variable in the repository.
- 'Invoke-Clean' - Cleans the project from build and testing artifacts. This command takes one parameter:
  - 'TestResults' - A switch parameter that, if set, will also remove the test files, such as coverage.out and any test binaries. The default is $true.
- 'Invoke-Build' - Builds the project and generates the binary executable. This command takes two parameters:
  - 'Configuration' - The build configuration to use. This can be either 'Debug' or 'Release'. The default is 'Release'. -- (Note: Currently this parameter is not used, but it is included for future use.)
  - 'OutputDirectory' - The directory where the built binary will be output. The default is '.\bin'.
- 'Invoke-Test' - Runs the tests for the project. This command takes two parameters:
  - 'Coverage' - A switch parameter that, if set, will generate a coverage report. The default is $true.
  - 'OutputDirectory' - The directory where the test results will be output. The default is '.\TestResults'.
- 'Start-Run' - Runs the application in its current state. This command takes no parameters.

## License

Apache License 2.0

This project is licenses under the [Apache 2.0 License](https://spdx.org/licenses/Apache-2.0.html) see the [LICENSE](LICENSE) file for more information.

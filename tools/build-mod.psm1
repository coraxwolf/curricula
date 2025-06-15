<#
  .SYNOPSIS
    Powershell Module to run Make type commands to build, develop, test, and manage project files.
  
  .DESCRIPTION
    This module provides a set of functions to streamline the process of building, developing, testing, and managing project files. It includes commands for initializing projects, running tests, building artifacts, and cleaning up generated files.
  
#>

# Requires -Version 7.5

$PROJECT_NAME = "curricula"
# $PROJECT_REPO = "github.com/coraxwolf/curricula"
$BIN_NAME = "curricula"

function getVersion {
  <#
    .SYNOPSIS
      Retrieves the version of the project from the VERSION file.
    
    .DESCRIPTION
      This function reads the VERSION file located in the project directory and returns its content as a string. If the file does not exist, it returns "0.0.0".
  #>
  [CmdletBinding()]
  param ()

  $versionFilePath = ".\VERSION"
  if (Test-Path -Path $versionFilePath) {
    $version = Get-Content -Path $versionFilePath -Raw
    return $version.Trim()
  }
  else {
    return "0.0.0"
  }
}

function getBuildNumber {
  <#
    .SYNOPSIS
      Generates a build number from the lastest git commit hash for the current branch.
    
    .DESCRIPTION
      This function gets the commit hash for the latest commit on the current branch and returns the first 7 characters as the build number. If git is not available or there are no commits, it returns "0000000".
  #>
  [CmdletBinding()]
  param ()

  try {
    $gitCommitHash = git rev-parse HEAD 2>$null
    if ($LASTEXITCODE -eq 0 -and $gitCommitHash) {
      return $gitCommitHash.Substring(0, 7)
    }
    else {
      return "0000000"
    }
  }
  catch {
    return "0000000"
  }
}

function getBuildState {
  <#
    .SYNOPSIS
      Checks the state of the git working tree and returns "dev" if there are uncommitted changes, otherwise returns "clean".
    
    .DESCRIPTION
      This function returns 'dev' if there are uncommitted changes in the git working tree, otherwise it returns 'clean'. If git is not available, it returns 'unknown'.
  #>
  [CmdletBinding()]
  param ()

  try {
    $gitStatus = git status --porcelain 2>$null
    if ($LASTEXITCODE -eq 0) {
      if ($gitStatus) {
        return "dev"
      }
      else {
        return "clean"
      }
    }
    else {
      return "unknown"
    }
  }
  catch {
    return "unknown"
  }
}

function Invoke-Clean {
  <#
    .SYNOPSIS
      Cleans build artifacts including binaries, temp files, and testing files.
    
    .DESCRIPTION
      This function removes any build artifacts that are generated during the build and testing processes. It ensures that the project directory is clean and free of any unnecessary files.
    
    .PARAMETER TestResults
      If set to true, the function will also remove test result and coverage files. (.out files)
  #>
  [CmdletBinding()]
  param (
    [bool]$TestResults = $true
  )

  Write-Host "Cleaning build artifacts..."
  if (Test-Path -Path ".\tmp") {
    Remove-Item -Path ".\tmp" -Recurse -Force
    Write-Host "Removed tmp directory."
  }
  else {
    Write-Host "No tmp directory found."
  }

  Write-Host "Cleaning Testing files..."
  Get-ChildItem -Path . -Recurse -Include "*.test.exe" | ForEach-Object {
    Remove-Item -Path $_.FullName -Force
    Write-Host "Removed $($_.FullName)"
  }

  if ($TestResults) {
    Write-Host "Cleaning Test Results files..."
    Get-ChildItem -Path . -Recurse -Include "*.out" | ForEach-Object {
      Remove-Item -Path $_.FullName -Force
      Write-Host "Removed $($_.FullName)"
    }
  }
  
  Write-Host "Clean completed."
}

function generateLDFlags {
  <#
    .SYNOPSIS
      Generates linker flags for embedding version information into binaries.
    
    .DESCRIPTION
      This function generates a string of linker flags that can be used to embed version information such as version number, build number, and build state into the compiled binaries. It retrieves the version from the VERSION file and the build number and state from git.
    
    .OUTPUTS
      A string containing the generated linker flags.
  #>
  [CmdletBinding()]
  param (
  )

  $Prefix = "github.com/coraxwolf/curricula/config"

  $version = getVersion
  $buildNumber = getBuildNumber
  $buildState = getBuildState
  $buildTime = (Get-Date).ToString("yyyy-MM-ddTHH:mm:ssZ")

  $ldFlags = @(
    "-X ${Prefix}.VERSION=$version"
    "-X ${Prefix}.BUILD_NUMBER=$buildNumber"
    "-X ${Prefix}.BUILD_STATE=$buildState"
    "-X ${Prefix}.BUILD_TIME=$buildTime"
  )

  return $ldFlags -join " "
}

function Invoke-Build {
  <#
    .SYNOPSIS
      Builds the project by compiling source files and generating necessary artifacts.
    
    .DESCRIPTION
      This function compiles the source files in the project directory and generates the necessary build artifacts. It ensures that all dependencies are resolved and that the build process is completed successfully.
    
    .PARAMETER Configuration
      Specifies the build configuration (e.g., Debug, Release). Default is 'Release'.
    
    .PARAMETER OutputDirectory
      Specifies the output directory for the build artifacts. Default is 'bin'.
  #>
  [CmdletBinding()]
  param (
    [string]$Configuration = "Release",
    [string]$OutputDirectory = "bin"
  )

  Write-Host "Building project $PROJECT_NAME with configuration: $Configuration"

  $APP_FILE = "$OutputDirectory\$BIN_NAME"
  
  if (-Not (Test-Path -Path $OutputDirectory)) {
    New-Item -ItemType Directory -Path $OutputDirectory | Out-Null
    Write-Host "Created output directory: $OutputDirectory"
  }
  $ldFlags = generateLDFlags
  Write-Host "Using linker flags: $ldFlags"
  go build -o $APP_FILE -ldflags $ldFlags .
  # Example build command (replace with actual build logic)
  # Here we just simulate a build process
  Start-Sleep -Seconds 2
  Write-Host "Build completed. Artifacts are in $OutputDirectory"
}

function Start-Run {
  <#
    .SYNOPSIS
      Runs the current version of the application.
    
    .DESCRIPTION
      This function executes the built application. It ensures that the application is run in the correct environment and captures any output or errors.
    
  #>
  [CmdletBinding()]
  param ()

  Write-Host "Running application $PROJECT_NAME ..."
  go run -ldflags "$(generateLDFlags)" .
}

function Invoke-Test {
  <#
    .SYNOPSIS
      Runs tests for the project and generates test reports.
    
    .DESCRIPTION
      This function executes the test suite for the project, capturing results and generating reports. It ensures that all tests are run in a clean environment and that any failures are reported.
    
    .PARAMETER Coverage
      If set to true, the function will generate a code coverage report.
    
    .PARAMETER OutputDirectory
      Specifies the output directory for test reports. Default is 'tmp'.
  #>
  [CmdletBinding()]
  param (
    [bool]$Coverage = $true,
    [string]$OutputDirectory = "tmp"
  )

  Write-Host "Running tests..."
  if (-Not (Test-Path -Path $OutputDirectory)) {
    New-Item -ItemType Directory -Path $OutputDirectory | Out-Null
    Write-Host "Created output directory: $OutputDirectory"
  }

  $testArgs = @("-v", "./...")
  if ($Coverage) {
    $coverageFile = "$OutputDirectory\coverage.out"
    $testArgs += @("-coverprofile=$coverageFile")
  }

  go test @testArgs

  if ($Coverage) {
    Write-Host "Generating coverage report..."
    go tool cover -html=$coverageFile -o "$OutputDirectory\coverage.html"
    Write-Host "Coverage report generated at $OutputDirectory\coverage.html"
  }

  Write-Host "Tests completed."
}

function Set-Version {
  <#
    .SYNOPSIS
      Sets the project version by updating the VERSION file.
    
    .DESCRIPTION
      This function updates the VERSION file with the specified version string. It ensures that the version format is valid and that the file is updated correctly.
    
    .PARAMETER Version
      The new version string to set in the VERSION file.
  #>
  [CmdletBinding()]
  param (
    [Parameter(Mandatory = $true)]
    [string]$Version
  )

  if ($Version -match '^\d+\.\d+\.\d+?$') {
    Set-Content -Path ".\VERSION" -Value $Version
    Write-Host "Version set to $Version"
  }
  else {
    Write-Error "Invalid version format. Use semantic versioning (e.g., 1.0.0 or 1.1.0 or 1.2.3)."
  }
}


Export-ModuleMember -Function *-*
Export-ModuleMember -Variable PROJECT_NAME, PROJECT_REPO, BIN_NAME
# Check if running as administrator
if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Warning "Please run this script as Administrator!"
    exit 1
}

# Get the current directory where the script is located
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $scriptDir

# Copy to a directory in your PATH
Copy-Item -Path ".\saika.exe" -Destination "C:\Windows\System32\" -Force

Write-Host "Saika has been installed to C:\Windows\System32\saika.exe"
Write-Host "You can now run it from anywhere by typing 'saika'"

# Ask if user wants to clean up
$cleanup = Read-Host "Do you want to remove the downloaded archive and extracted files after installation? (y/n)"

if ($cleanup -eq "y" -or $cleanup -eq "Y") {
    # Get parent directory
    $parentDir = Split-Path -Parent $scriptDir
    
    # Go up one level and remove the directory
    Set-Location ..
    Remove-Item -Recurse -Force $scriptDir
    
    # Find and remove the archive file
    $archiveFile = Get-ChildItem -Path $parentDir -Filter "saika-*.zip" | Select-Object -First 1
    if ($archiveFile) {
        Remove-Item $archiveFile.FullName
        Write-Host "Cleaned up downloaded files."
    }
    
    Write-Host "Installation complete and cleanup done."
} else {
    Write-Host "Installation complete. Downloaded files were kept."
}
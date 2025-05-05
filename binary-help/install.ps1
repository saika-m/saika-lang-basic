# Check if running as administrator
if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Warning "Please run this script as Administrator!"
    exit 1
}

# Copy to a directory in your PATH
Copy-Item -Path ".\saika.exe" -Destination "C:\Windows\System32\" -Force

Write-Host "Saika has been installed to C:\Windows\System32\saika.exe"
Write-Host "You can now run it from anywhere by typing 'saika'"
Write-Host "Installation complete!"
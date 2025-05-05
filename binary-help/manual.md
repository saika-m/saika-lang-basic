**Linux (x86_64):**
```bash
# Extract the archive
tar -xzf saika-linux-amd64.tar.gz
cd linux-amd64

# Make executable
chmod +x saika

# Copy to system path
sudo cp saika /usr/local/bin/

# Verify installation
saika
```

**Linux (ARM64):**
```bash
tar -xzf saika-linux-arm64.tar.gz
cd linux-arm64
chmod +x saika
sudo cp saika /usr/local/bin/
saika
```

**Linux (386):**
```bash
tar -xzf saika-linux-386.tar.gz
cd linux-386
chmod +x saika
sudo cp saika /usr/local/bin/
saika
```

**macOS (Intel):**
```bash
tar -xzf saika-macos-amd64.tar.gz
cd macos-amd64
chmod +x saika
sudo cp saika /usr/local/bin/
saika
```

**macOS (Apple Silicon):**
```bash
tar -xzf saika-macos-arm64.tar.gz
cd macos-arm64
chmod +x saika
sudo cp saika /usr/local/bin/
saika
```

**Windows (x86_64):**
```powershell
Expand-Archive saika-windows-amd64.zip
cd saika-windows-amd64
copy saika.exe C:\Windows\System32\
saika
```

**Windows (ARM64):**
```powershell
Expand-Archive saika-windows-arm64.zip
cd saika-windows-arm64
copy saika.exe C:\Windows\System32\
saika
```

**Windows (386):**
```powershell
Expand-Archive saika-windows-386.zip
cd saika-windows-386
copy saika.exe C:\Windows\System32\
saika
```
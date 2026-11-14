$ErrorActionPreference = "Stop"

$Version = if ($env:DSGO_VERSION) { $env:DSGO_VERSION } else { "v0.1.0" }
$Repo = "wong4ga1ming4/dsgo"
$Name = "dsgo-windows-amd64.exe"
$Url = "https://github.com/$Repo/releases/download/$Version/$Name"

if ($env:DSGO_BIN) {
    $BinDir = $env:DSGO_BIN
} else {
    $BinDir = Join-Path $env:LOCALAPPDATA "dsgo"
}
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null

$Dest = Join-Path $BinDir "dsgo.exe"
Invoke-WebRequest -Uri $Url -OutFile $Dest

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (-not $UserPath) { $UserPath = "" }
$Parts = $UserPath -split ";" | Where-Object { $_ -and ($_ -ne $BinDir) }
$NewPath = (@($BinDir) + $Parts) -join ";"
[Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
$env:Path = "$BinDir;" + $env:Path

Write-Host "已安装 $Dest ($Version)"
Write-Host "新开一个终端后可直接运行 dsgo"

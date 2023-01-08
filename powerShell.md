```javascript
# 自定义prompt, 显示git信息
Set-PSReadLineOption -Colors @{
    Command            = 'Green'
    Number             = 'Yellow'
    Member             = 'Yellow'
    Operator           = 'White'
    Type               = 'Green'
    Variable           = 'Yellow'
    Parameter          = 'Yellow'
    ContinuationPrompt = 'Yellow'
    Default            = 'Yellow'
}

function getTimeStamp([int]$useStamp) {
    $(if ($useStamp -ne 0) {
        $(Get-Date -Format ' [yy-MM-dd HH:mm:ss]')
    } else { "" })
}

function getAdminPrefix {
    $identity = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = [Security.Principal.WindowsPrincipal] $identity
    $adminRole = [Security.Principal.WindowsBuiltInRole]::Administrator

    $(if ($principal.IsInRole($adminRole)) {
        "[ADMIN]"
    } else { "" })
}

function getParentGitBranchName([String]$rootPath) {
    $currentLevelPath = $rootPath

    while (-not [string]::IsNullOrEmpty($currentLevelPath)) {
        $gitPath = Join-Path $currentLevelPath ".git"
        $gitHeadPath = Join-Path $gitPath "HEAD"

        if ((Test-Path $gitPath) -and (Test-Path $gitHeadPath)) {
            $content = Get-Content $gitHeadPath
            $gitBranchRgx = "ref\:\s*refs\/heads\/([a-zA-Z0-9\/\\]+)"

            if ((-not [string]::IsNullOrEmpty($content)) -and ($content -match $gitBranchRgx)) {
                return $Matches.1
            } else {
                return $content
            }
        }
        $currentLevelPath = Split-Path -Parent $currentLevelPath
    }
    return ""
}

function prompt {
    $currentPath = ($pwd).Path
    $currentUserName = $env:USERNAME
    $computerName = $(Get-WMIObject Win32_ComputerSystem).name
    $gitBranchName = $(getParentGitBranchName $currentPath)
    $adminPrefix = $(getAdminPrefix)

    # 使用Unix风格用户目录符号 ~, 与 C:/Users/username/ 目录等价
    $pattern = $HOME -replace '\\', '\\'
    $currentPath = $currentPath -replace $pattern, "~"
    $ESC = $([char]27)

    $adminPrefix +
    "$ESC[35;1m$currentUserName$ESC[0m" + '@' + "$ESC[32;1m$computerName" +
    "$ESC[36;1m $currentPath$ESC[0m" +
    $(if (-not [string]::IsNullOrEmpty($gitBranchName)) { "$ESC[32;1m ($gitBranchName)$ESC[0m" }) +
    $(getTimeStamp 1) +
    "$ESC[36;1m`n$ $ESC[0m"
}

# 代码页设置为65001，可以删掉
chcp 65001# 自定义prompt, 显示git信息
Set-PSReadLineOption -Colors @{
    Command            = 'Green'
    Number             = 'Yellow'
    Member             = 'Yellow'
    Operator           = 'White'
    Type               = 'Green'
    Variable           = 'Yellow'
    Parameter          = 'Yellow'
    ContinuationPrompt = 'Yellow'
    Default            = 'Yellow'
}

function getTimeStamp([int]$useStamp) {
    $(if ($useStamp -ne 0) {
        $(Get-Date -Format ' [yy-MM-dd HH:mm:ss]')
    } else { "" })
}

function getAdminPrefix {
    $identity = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = [Security.Principal.WindowsPrincipal] $identity
    $adminRole = [Security.Principal.WindowsBuiltInRole]::Administrator

    $(if ($principal.IsInRole($adminRole)) {
        "[ADMIN]"
    } else { "" })
}

function getParentGitBranchName([String]$rootPath) {
    $currentLevelPath = $rootPath

    while (-not [string]::IsNullOrEmpty($currentLevelPath)) {
        $gitPath = Join-Path $currentLevelPath ".git"
        $gitHeadPath = Join-Path $gitPath "HEAD"

        if ((Test-Path $gitPath) -and (Test-Path $gitHeadPath)) {
            $content = Get-Content $gitHeadPath
            $gitBranchRgx = "ref\:\s*refs\/heads\/([a-zA-Z0-9\/\\]+)"

            if ((-not [string]::IsNullOrEmpty($content)) -and ($content -match $gitBranchRgx)) {
                return $Matches.1
            } else {
                return $content
            }
        }
        $currentLevelPath = Split-Path -Parent $currentLevelPath
    }
    return ""
}

function prompt {
    $currentPath = ($pwd).Path
    $currentUserName = $env:USERNAME
    $computerName = $(Get-WMIObject Win32_ComputerSystem).name
    $gitBranchName = $(getParentGitBranchName $currentPath)
    $adminPrefix = $(getAdminPrefix)

    # 使用Unix风格用户目录符号 ~, 与 C:/Users/username/ 目录等价
    $pattern = $HOME -replace '\\', '\\'
    $currentPath = $currentPath -replace $pattern, "~"
    $ESC = $([char]27)

    $adminPrefix +
    "$ESC[35;1m$currentUserName$ESC[0m" + '@' + "$ESC[32;1m$computerName" +
    "$ESC[36;1m $currentPath$ESC[0m" +
    $(if (-not [string]::IsNullOrEmpty($gitBranchName)) { "$ESC[32;1m ($gitBranchName)$ESC[0m" }) +
    $(getTimeStamp 1) +
    "$ESC[36;1m`n$ $ESC[0m"
}

# 代码页设置为65001，可以删掉
chcp 65001
```

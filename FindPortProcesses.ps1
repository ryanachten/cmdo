param([switch]$kill)
$ports = @(3000, 5099, 5000, 11119, 5173, 8080, 1111, 1112, 1113, 1114, 1115)

foreach ($port in $ports) {
    $foundProcesses = netstat -ano | findstr :$port
    $activePortPattern = ":$port\s.+LISTENING\s+\d+$"
    $pidNumberPattern = "\d+$"

    if ($foundProcesses | Select-String -Pattern $activePortPattern -Quiet) {
        $matches = $foundProcesses | Select-String -Pattern $activePortPattern
        $firstMatch = $matches.Matches.Get(0).Value

        $pidNumber = [regex]::match($firstMatch, $pidNumberPattern).Value

        if($kill) {
            taskkill /pid $pidNumber /f
        }
    } else {
        Write-Host "No process found running on port $port"
    }
}
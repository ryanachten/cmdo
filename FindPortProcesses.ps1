param([switch]$kill)
$ports = @(3000, 5099, 5000, 11119, 5173, 8080)

foreach ($port in $ports) {
    $connection = Get-NetTCPConnection -LocalPort $port -ErrorAction SilentlyContinue
    if($connection) {
        $process = Get-Process -Id ($connection).OwningProcess
        if ($process -and $kill) {
            Write-Host "Killing process $($process.Id) running on port $port"
            Stop-Process -Id $process.Id -Force
            continue
        }
        elseif ($process) {
            Write-Host "Found process $($process.Id) running on port $port"
            continue
        }
    }
    Write-Host "No process found running on port $port"
}
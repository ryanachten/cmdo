using System.Diagnostics;

namespace Commando.Models;

public class CommandTask
{
    public required string Name { get; init; }
    public required string FileName { get; init; }
    public required string Arguments { get; init; }
    public required string WorkingDirectory { get; init; }
    public ConsoleColor Color { get; set; }

    public async Task<Process> Start(CancellationToken token)
    {
        var processInfo = new ProcessStartInfo()
        {
            FileName = FileName,
            Arguments = Arguments,
            CreateNoWindow = true,
            RedirectStandardInput = true,
            RedirectStandardError = true,
            RedirectStandardOutput = true,
            WorkingDirectory = WorkingDirectory,
        };
        var process = Process.Start(processInfo);
        if (process == null) throw new TaskCanceledException($"Could not start process {Name}");

        Console.ForegroundColor = Color;
        Console.WriteLine($"Starting {Name} process with ID {process.Id}");

        process.EnableRaisingEvents = true;
        process.OutputDataReceived += (sender, e) =>
        {
            if (!string.IsNullOrEmpty(e.Data))
            {
                Console.ForegroundColor = Color;
                Console.WriteLine($"[{Name}] {e.Data}");
            }
        };

        process.ErrorDataReceived += (sender, e) =>
        {
            if (!string.IsNullOrEmpty(e.Data))
            {
                Console.ForegroundColor = ConsoleColor.Red;
                Console.Error.WriteLine($"[{Name}] Error: {e.Data}");
            }
        };

        process.BeginOutputReadLine();
        process.BeginErrorReadLine();

        // Will await for either the process to exit or the cancellation token to be cancelled
        await Task.WhenAny(process.WaitForExitAsync(), Task.Delay(-1, token));

        return process;
    }
}

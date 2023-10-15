using System.Diagnostics;
using Commando.Models;

namespace Commando.Services;

public static class CommandService
{
    private static readonly ConsoleColor[] _colors = new[] { ConsoleColor.Green, ConsoleColor.Magenta, ConsoleColor.Blue, ConsoleColor.Cyan };

    public static async Task Run(List<CommandTask> commands)
    {
        using var tokenSource = new CancellationTokenSource();
        var token = tokenSource.Token;

        // Cancel all child processes when user cancels console
        Console.CancelKeyPress += (sender, e) =>
        {
            e.Cancel = true;
            tokenSource.Cancel();
        };

        var processes = new List<Task<Process>>();
        for (int i = 0; i < commands.Count; i++)
        {
            var command = commands[i];
            command.Color = _colors[i];
            processes.Add(command.Start(token));
        }

        // Keep main thread alive until all processes finish
        while (processes.Count > 0)
        {
            // If any process exits, we'll close all other processes
            await Task.WhenAny(processes);

            Console.ResetColor();

            // Terminate all processes and their process tree
            foreach (var process in processes.Select(p => p.Result))
            {
                Console.WriteLine($"Stopping process: {process.Id}");
                process.Kill(true);
            }

            processes.Clear();
        }
    }
}

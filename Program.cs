using Commando.Configuration;
using Commando.Services;
using System.CommandLine;
using System.Text.Json;

internal class Program
{
    private static async Task Main(string[] args)
    {
        var configuration = await ParseArguments(args);
        if (configuration == null) return;

        await CommandService.Run(configuration.Commands);
    }

    private static async Task<Configuration?> ParseArguments(string[] arguments)
    {
        var configOption = new Option<string>(name: "--config", description: "Configuration file path containing commands to be executed");
        var rootCommand = new RootCommand("Runs commands in parallel");
        rootCommand.AddOption(configOption);

        Configuration? configuration = null;
        rootCommand.SetHandler(
            async (pathName) =>
            {
                configuration = await Deserialize<Configuration?>(pathName);
            },
            configOption
        );

        await rootCommand.InvokeAsync(arguments);

        return configuration;
    }

    private static async Task<T?> Deserialize<T>(string filePath)
    {
        using FileStream stream = File.OpenRead(filePath);
        return await JsonSerializer.DeserializeAsync<T>(stream, new JsonSerializerOptions
        {
            PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
        });
    }
}
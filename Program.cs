using Microsoft.Extensions.Configuration;
using System.Configuration;
using Commando.Models;
using Commando.Services;

internal class Program
{
    private static async Task Main(string[] args)
    {
        var configuration = new ConfigurationBuilder()
         .AddJsonFile("appsettings.json").Build();

        var commands = configuration.GetRequiredSection("Commands").Get<List<CommandTask>>();
        if (commands == null) throw new ConfigurationErrorsException("Missing tasks configuration");

        await CommandService.Run(commands);
    }
}
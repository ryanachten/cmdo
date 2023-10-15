using Commando.Models;

namespace Commando.Configuration;

public class Configuration
{
    public required List<CommandTask> Commands { get; init; }
}

using System.CommandLine;
using System.CommandLine.Binding;

namespace Commando.Configuration;

public class ConfigurationBinder : BinderBase<Configuration>
{
    private readonly Option<Configuration> _configuration;

    public ConfigurationBinder(Option<Configuration> configuration)
    {
        _configuration = configuration;
    }
    protected override Configuration GetBoundValue(BindingContext bindingContext)
    {
        var result = bindingContext.ParseResult.GetValueForOption(_configuration);
        return result ?? new Configuration() { Commands = new() };
    }
}

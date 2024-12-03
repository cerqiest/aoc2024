var input = File.ReadLines("../../../input.txt");
var safeReports = 0;

bool ReportIsSafe(List<string> levels)
{
    bool? increasing = null;
    int? lastLevel = null;

    foreach (var level in levels)
    {
        var levelInt = int.Parse(level);
        
        if (lastLevel.HasValue)
        {
            increasing ??= levelInt > lastLevel;
            
            if ((increasing ?? true) && levelInt <= lastLevel)
            {
                return false;
            }
            
            if (!(increasing ?? true) && levelInt >= lastLevel)
            {
                return false;
            }
            
            if (Math.Abs(levelInt - lastLevel.Value) > 3)
            {
                return false;
            }
        }
        
        lastLevel = levelInt;
    }

    return true;
}

bool WouldBeDampened(List<string> levels)
{
    for (int i = 0; i < levels.Count; i++)
    {
        var levelsCopy = new List<string>(levels);
        levelsCopy.RemoveAt(i);
        if (ReportIsSafe(levelsCopy))
        {
            return true;
        }
    }
    
    return false;
}

foreach (var line in input)
{
    var levels = line.Split(" ").ToList(); 
    if (ReportIsSafe(levels))
    {
        safeReports++;
        continue;
    }
    
    if (WouldBeDampened(levels))
    {
        safeReports++;
    }
}

Console.WriteLine(safeReports);
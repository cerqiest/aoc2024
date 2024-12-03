var input = File.ReadLines("../../../input.txt");
var safeReports = 0;

foreach (var line in input)
{
    var levels = line.Split(" ");
    var safeReport = true;
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
                safeReport = false;
                break;
            }
            
            if (!(increasing ?? true) && levelInt >= lastLevel)
            {
                safeReport = false;
                break;
            }
            
            if (Math.Abs(levelInt - lastLevel.Value) > 3)
            {
                safeReport = false;
                break;
            }
        }
        
        lastLevel = levelInt;
    }
    
    if (safeReport)
    {
        safeReports++;
    }
}

Console.WriteLine(safeReports);
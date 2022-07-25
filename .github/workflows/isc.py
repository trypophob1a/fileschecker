import sys
import re

args = sys.argv[1:3]
file = open(sys.argv[1], "r")
MIN_DEFAULT_PERCENT = 50

try:
    args[1]
except IndexError:
     args.append(MIN_DEFAULT_PERCENT)

try:
    min_percent = int(args[1])
except ValueError:
    print(f"\033[1m\033[91m ERROR: min percent: <<{args[1]}>> not is number:\n")
    sys.exit(1)

lines = file.readlines()
if not lines:
   print(f"\033[92m test there are no files for testing")
   os.environ["NO_FILE"] = "1"
   sys.exit(0)

regex = r"LF:(\d+)|LH:(\d+)"
full = 0
part = 0
matcher = re.compile(regex, flags=re.IGNORECASE)

for line in lines:
   matchValue = matcher.match(line.strip())
   if matchValue:
      full+= int(matchValue.group(1)) if matchValue.group(1) != None else 0
      part+= int(matchValue.group(2)) if matchValue.group(2) != None else 0
file.close

percent = round((part * 100) / full)
if  percent < min_percent:
   print(f"\033[1m\033[91m ERROR: insufficient coverage {percent}% needs {min_percent}%\n")
   sys.exit(1)
else:
    print(f"\033[92m coverage is: {percent}%")
import re

with open("./source.html", "r") as f:
    lines = f.readlines()


with open("out", "w") as f:
    for line in lines:
        striped = line.strip()
        if (striped[:4] == "<li>"):
            tmp = striped[4:-5]
            tmp = re.sub('\d+_', '', tmp)
            tmp = re.sub('\.wav', '', tmp)
            tmp = re.sub('\(.*?\)', '', tmp)
            tmp = re.sub('（.*?）', '', tmp)
            if len(tmp) > 0:
                f.write(tmp + '\n')

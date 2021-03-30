import json
import argparse

parser = argparse.ArgumentParser(
    prog='change-sarif'
)

parser.add_argument(
    '--ruleid',
    help='Input ruleid deliminator seperated',
    required=True
)
parser.add_argument(
    '--severity',
    help='Input severity',
    required=True
)

args = parser.parse_args()
ruleids = args.ruleid.split('|')
severity = args.severity

print(ruleids)
print(severity)
# Opening JSON file
with open('go-builtin.sarif', 'r') as f:

    # returns JSON object as
    # a dictionary
    sarif = json.load(f)

    for run in sarif.get('runs', []):
        rules = run.get('tool').get('driver').get('rules')
        for rule in rules:
            id = rule.get('id')
            for ruleid in ruleids:
                if(ruleid == id):
                    rule['properties']['problem.severity'] = "hello world"
f.close

with open('go-builtin.sarif', 'w') as f:
    json.dump(sarif, f, indent=2)
f.close

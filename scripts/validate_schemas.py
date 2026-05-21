import os
import json
from jsonschema import Draft7Validator

def validate_schemas(schema_dir):
    for filename in os.listdir(schema_dir):
        if filename.endswith(".json"):
            path = os.path.join(schema_dir, filename)
            print(f"Validating {path}...")
            with open(path, "r") as f:
                try:
                    schema = json.load(f)
                    Draft7Validator.check_schema(schema)
                    print(f"  {filename} is a valid JSON schema.")
                except Exception as e:
                    print(f"  Error validating {filename}: {e}")
                    return False
    return True

if __name__ == "__main__":
    schema_dir = "02_docs/schemas"
    if validate_schemas(schema_dir):
        print("All schemas are valid.")
    else:
        exit(1)

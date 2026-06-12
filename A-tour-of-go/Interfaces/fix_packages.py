import os
import re

dir_path = "/home/divyansh-rawat/backend-in-golang/A-tour-of-go/Interfaces"


def repl(content, old, new):
    return re.sub(r"\b" + old + r"\b", new, content)


for i in range(1, 14):
    filename = f"i{i}.go"
    filepath = os.path.join(dir_path, filename)
    if not os.path.exists(filepath):
        continue

    with open(filepath, "r") as f:
        content = f.read()

    # package main -> package interfaces
    content = re.sub(
        r"^package main", "package interfaces", content, flags=re.MULTILINE
    )

    # main -> I{i}
    content = re.sub(r"func main\(\)", f"func I{i}()", content)

    # Specific type renames to avoid conflicts
    if i in [2, 3, 4, 5, 6]:
        content = repl(content, "I", f"I_{i}")
        # Fix %I_x back to %T if it was messed up, but 'I' won't match '%T'
    if i in [2, 3, 4, 5]:
        content = repl(content, "T", f"T_{i}")
        # Fix fmt.Printf("%T") -> fmt.Printf("%T_i") which is wrong
        content = content.replace(f"%T_{i}", "%T")

    if i == 3:
        content = repl(content, "F", "F_3")
        content = content.replace("%F_3", "%F")  # Just in case

    if i in [3, 4, 5, 6, 7]:
        content = repl(content, "describe", f"describe_{i}")

    if i == 6:
        content = repl(content, "MyType", "MyType_6")
        content = repl(content, "describe2", "describe2_6")
        content = repl(content, "runExample2", "runExample2_6")

    if i == 9:
        content = repl(content, "do", "do_9")

    if i == 10:
        content = repl(content, "Person", "Person_10")

    if i == 11:
        content = repl(content, "MyError", "MyError_11")
        content = re.sub(r"\brun\b", "run_11", content)

    with open(filepath, "w") as f:
        f.write(content)

print("Python script finished.")

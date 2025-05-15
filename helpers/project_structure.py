import os


EXCLUDE = {'__pycache__', '.git', '.venv'}

def is_excluded(path):
    for excluded in EXCLUDE:
        if path.startswith(excluded):
            return True
    return False

def list_sorted(path):
    items = os.listdir(path)
    return sorted(items, key=lambda name: (
        not os.path.isdir(os.path.join(path, name)),
        name.lower()
    ))

def scan_folder(path, root):
    result = {}
    for name in list_sorted(path):
        full_path = os.path.join(path, name)
        rel_path = os.path.relpath(full_path, root)

        if is_excluded(rel_path):
            continue

        if os.path.isdir(full_path):
            result[name] = scan_folder(full_path, root)
        else:
            result[name] = None
    return result

def ASCII_tree(tree, prefix=""):
    entries = list(tree.items())
    for idx, (name, subtree) in enumerate(entries):
        is_last = (idx == len(entries) - 1)
        branch = "└── " if is_last else "├── "
        print(f"{prefix}{branch}{name}")
        if subtree is not None:
            extension = "    " if is_last else "│   "
            ASCII_tree(subtree, prefix + extension)


root_dir = os.path.abspath(os.path.join(__file__, "..", ".."))
structure = {os.path.basename(root_dir): scan_folder(root_dir, root_dir)}
ASCII_tree(structure)

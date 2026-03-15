import os
import random
import re
import subprocess
import platform
from pathlib import Path

# --- CONFIGURATION ---
CONFIG = {
    "BASE_DIR": "./",  # Where your organized folders are
    "NUM_QUESTIONS": 5,               # Total questions to pick
    "OPEN_README": True,              # Automatically open the README files?
    "LEVELS_TO_SKIP": [],             # e.g., ["05_percent"] if you're past it
}


def get_sorted_levels(base_dir):
    """Returns a list of level folders sorted numerically by the percentage."""
    p = Path(base_dir)
    levels = [d for d in p.iterdir() if d.is_dir() and d.name not in CONFIG["LEVELS_TO_SKIP"]]

    def natural_sort_key(path_obj):
        # Find all sequences of digits in the folder name
        numbers = re.findall(r'\d+', path_obj.name)
        # Return the first number found as an integer (e.g., "100" from "100_percent")
        return int(numbers[0]) if numbers else 0

    return sorted(levels, key=natural_sort_key)

def open_file(filepath):
    """Opens a file using the default system application."""
    if platform.system() == "Darwin":       # macOS
        subprocess.call(("open", filepath))
    elif platform.system() == "Windows":    # Windows
        os.startfile(filepath)
    else:                                   # Linux
        subprocess.call(("xdg-open", filepath))

def generate_mock_exam():
    base_path = Path(CONFIG["BASE_DIR"])
    if not base_path.exists():
        print(f"Error: {CONFIG['BASE_DIR']} does not exist. Run the organizer script first!")
        return

    levels = get_sorted_levels(CONFIG["BASE_DIR"])

    # We can only pick as many questions as we have levels
    limit = min(CONFIG["NUM_QUESTIONS"], len(levels))
    selected_levels = levels[:limit]

    print(f"--- 🚀 GENERATING MOCK EXAM ({limit} Questions) ---")

    for i, level_path in enumerate(selected_levels, 1):
        # Get all subjects in this level
        subjects = [s for s in level_path.iterdir() if s.is_dir()]

        if not subjects:
            print(f"Level {level_path.name}: No subjects found. Skipping...")
            continue

        # Pick one random subject from this level
        chosen_subject = random.choice(subjects)
        readme_path = chosen_subject / "README.md"

        print(f"Question {i} (Level {level_path.name}): {chosen_subject.name}")

        if CONFIG["OPEN_README"]:
            if readme_path.exists():
                open_file(readme_path)
            else:
                print(f"  ! Warning: No README.md found in {chosen_subject.name}")

    print("\nGood luck with your practice! Focus on the logic first.")

if __name__ == "__main__":
    generate_mock_exam()

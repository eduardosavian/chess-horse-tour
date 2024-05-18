import tkinter as tk
from tkinter import ttk, messagebox, BooleanVar
import subprocess
import json
import matplotlib.pyplot as plt
import matplotlib.animation as animation
import seaborn as sns
import platform
import numpy as np


class KnightsTourApp:
    def __init__(self, root):
        self.root = root
        self.root.title("Knight's Tour")

        self.board_size_var = tk.StringVar(value="8")
        self.start_x_var = tk.StringVar(value="0")
        self.start_y_var = tk.StringVar(value="0")
        self.algorithm_var = tk.StringVar(value="backtrack")
        self.animate_var = BooleanVar()

        self.create_widgets()


    def create_widgets(self):
        frm = ttk.Frame(self.root, padding=10)
        frm.pack()

        ttk.Label(frm, text="Board Size:").grid(column=0, row=0, sticky="e")
        ttk.Entry(frm, textvariable=self.board_size_var).grid(column=1, row=0)

        ttk.Label(frm, text="Start X:").grid(column=0, row=1, sticky="e")
        ttk.Entry(frm, textvariable=self.start_x_var).grid(column=1, row=1)

        ttk.Label(frm, text="Start Y:").grid(column=0, row=2, sticky="e")
        ttk.Entry(frm, textvariable=self.start_y_var).grid(column=1, row=2)

        ttk.Label(frm, text="Algorithm:").grid(column=0, row=3, sticky="e")
        algorithm_combo = ttk.Combobox(frm, textvariable=self.algorithm_var, values=["warnsdorff", "backtrack", "shuffle", "highDegree"])
        algorithm_combo.grid(column=1, row=3)
        algorithm_combo.current(0)

        ttk.Checkbutton(frm, text="Animate", variable=self.animate_var).grid(column=1, row=4, sticky="w")

        ttk.Button(frm, text="Find Tour", command=self.find_tour).grid(column=0, row=5, columnspan=2)


    def find_tour(self):
        board_size = int(self.board_size_var.get())
        start_x = int(self.start_x_var.get())
        start_y = int(self.start_y_var.get())
        algorithm = self.algorithm_var.get()
        animate = self.animate_var.get()

        if board_size <= 0 or start_x < 0 or start_y < 0 or start_x >= board_size or start_y >= board_size:
            messagebox.showerror("Error", "Invalid board size or start coordinates.")
            return

        if platform.system() == 'Windows':
            executable = "./libs/knight_tour.exe"
        else:
            executable = "./libs/knight_tour"

        result = subprocess.run(
            [executable, str(start_x), str(start_y), str(board_size), algorithm],
            capture_output=True,
            text=True
        ).stdout

        try:
            data = json.loads(result)
            if 'board' not in data:
                raise ValueError("The JSON does not contain the 'board' key.")
            board = data['board']
            if not all(isinstance(row, list) for row in board) or not all(isinstance(num, int) for row in board for num in row):
                raise ValueError("The 'board' key does not contain a valid 2D list of integers.")
        except json.JSONDecodeError:
            raise ValueError("Failed to decode JSON.")

        if animate:
            self.plot_animated(board)
        else:
            self.plot_static(board)


    def plot_static(self, board):
        plt.figure(figsize=(10, 8))
        sns.heatmap(board, annot=True, fmt="d", cmap="Reds", cbar=False, square=True)
        plt.title("Knight's Tour Heatmap")
        plt.xlabel("X Coordinate")
        plt.ylabel("Y Coordinate")
        plt.show()


    def plot_animated(self, board):
        max_size = max(len(board), len(board[0]))
        padded_board = np.zeros((max_size, max_size), dtype=int)
        padded_board[:len(board), :len(board[0])] = board

        fig, ax = plt.subplots(figsize=(10, 8))
        ax.set_title("Knight's Tour Animation")
        ax.set_xlabel("X Coordinate")
        ax.set_ylabel("Y Coordinate")


        def animate(frame):
            ax.clear()
            ax.set_title("Knight's Tour Animation")
            ax.set_xlabel("X Coordinate")
            ax.set_ylabel("Y Coordinate")
            sns.heatmap(padded_board, annot=True, fmt="d", cmap="Reds", cbar=False, square=True, ax=ax)

            current_num = frame + 1
            current_x, current_y = np.where(padded_board == current_num)

            if current_num > 1:
                previous_num = current_num - 1
                previous_x, previous_y = np.where(padded_board == previous_num)
                ax.plot([previous_y + 0.5, current_y + 0.5], [previous_x + 0.5, current_x + 0.5], color='blue', alpha=0.5)

            ax.add_patch(plt.Circle((current_y + 0.5, current_x + 0.5), 0.4, color='blue', alpha=0.5))

        num_frames = len(board) * len(board[0])
        ani = animation.FuncAnimation(fig, animate, frames=num_frames, repeat=False)
        plt.show()


if __name__ == "__main__":
    root = tk.Tk()
    app = KnightsTourApp(root)
    root.mainloop()


from collections import namedtuple
import json
import copy


class Point(object):
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    def __repr__(self):
        return f"Point({self.x}, {self.y}, {self.z})"


def make_rectangle(A: Point, W: Point, H: Point, r, g, b) -> dict:
    return {
        "type": "rectangle",
        "a": {"x": A.x, "y": A.y, "z": A.z},
        "w": {"x": W.x, "y": W.y, "z": W.z},
        "h": {"x": H.x, "y": H.y, "z": H.z},
        "mat": {
            "type": "metal",
            "albedo": {"x": r / 255, "y": g / 255, "z": b / 255},
            "fuzz": 0,
        },
    }


def make_chessboard(starting: Point, w: int, l: int) -> str:
    config = {"random": False}
    spheres = [
        {
            "type": "sphere",
            "center": {"x": 0, "y": 0.5, "z": -1},
            "radius": 0.5,
            "mat": {"type": "lambertian", "albedo": {"x": 0.1, "y": 0.2, "z": 0.5}},
        },
        {
            "type": "sphere",
            "center": {"x": 1, "y": 0.5, "z": -1},
            "radius": 0.5,
            "mat": {
                "type": "metal",
                "albedo": {"x": 0.8, "y": 0.6, "z": 0.2},
                "fuzz": 0,
            },
        },
        {
            "type": "sphere",
            "center": {"x": -1, "y": 0.5, "z": -1},
            "radius": 0.5,
            "mat": {"type": "dielectric", "refindex": 1.5},
        },
    ]
    cur = copy.copy(starting)
    square_dim = 0.5
    board = []
    row_start_color = 0xFF

    while cur.z < (starting.z + l):
        cur.x = starting.x
        color = row_start_color
        while cur.x < (starting.x + w):
            board.append(
                make_rectangle(
                    cur,
                    Point(cur.x + square_dim, cur.y, cur.z),
                    Point(0, 0, square_dim),
                    color,
                    color,
                    color,
                )
            )
            color ^= 0xFF
            cur.x += square_dim
        cur.z += square_dim
        row_start_color ^= 0xFF
    config["static"] = board + spheres 
    return json.dumps(config, indent=4)


if __name__ == "__main__":
    chessboard = make_chessboard(Point(-8, 0, -8), 16, 16)
    print(chessboard)

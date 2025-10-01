from __future__ import annotations


class Zombie:
    count = 0
    STARTING_HP = 3

    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y
        self.hp = Zombie.STARTING_HP
        Zombie.count += 1

    def hit(self) -> Zombie:
        """Hit the zombie, subtracting 1 of its HP."""
        self.hp -= 1
        return self

    def is_alive(self) -> bool:
        """Determine if the zombie is alive."""
        return self.hp > 0

    @classmethod
    def from_list(cls, coordinates: list[tuple[int, int]]) -> list[Zombie]:
        """Return a list of zombies with the given coordinates."""
        return [cls(x, y) for x, y in coordinates]

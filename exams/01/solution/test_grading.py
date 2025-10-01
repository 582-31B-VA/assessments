from pytest import raises
from main import Zombie


def test_zombie_initialization():
    m = Zombie(1, 2)
    assert m.x == 1
    assert m.y == 2
    assert m.hp == 3
    assert Zombie.count == 1


def test_is_alive_false_when_hp_positive():
    m = Zombie(2, 2)
    m.hit()
    assert not m.is_alive()


def test_is_alive():
    m = Zombie(3, 3)
    m.hp = 3
    assert m.is_alive()
    m.hp = 0
    assert not m.is_alive()


def test_from_list_creates_zombies():
    lst = [(1, 2), (3, 4)]
    zombies = Zombie.from_list(lst)
    assert len(zombies) == 2
    assert isinstance(zombies[0], Zombie)
    assert zombies[0].x == 1 and zombies[0].y == 2
    assert zombies[1].x == 3 and zombies[1].y == 4
    assert Zombie.count == 2

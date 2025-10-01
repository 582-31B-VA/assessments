from enum import Enum


class Status(Enum):
    PROCESSING = 0
    SHIPPED = 1
    DELIVERED = 2


class Item:
    def __init__(self, name: str, price: int):
        self.name = name
        self.price = price

    @property
    def price(self) -> int:
        """Price of the item."""
        return self._price

    @price.setter
    def price(self, value: int) -> None:
        if value < 0:
            raise ValueError(f"{value} is negative: price must be positive")
        self._price = value


class Order:
    def __init__(self, items: list[Item]):
        self.items = items
        self.status = Status.PROCESSING

    @property
    def total(self) -> int:
        """The order's total in cents."""
        return sum(item.price for item in self.items)

    @property
    def status_message(self) -> str:
        """Message for the current status."""
        if self.status is Status.PROCESSING:
            return "We are processing your order."
        if self.status is Status.SHIPPED:
            return "Your order has been shipped."
        else:
            return "Your order has been delivered."

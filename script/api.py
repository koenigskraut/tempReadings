from dataclasses import dataclass, field
from datetime import datetime
import requests

API_ROOT = "https://koenigskraut.ru/api"

LAST_N_READINGS_URL = f"{API_ROOT}/lastNReadings"
FIRST_N_READINGS_URL = f"{API_ROOT}/firstNReadings"
AVERAGE_READINGS_URL = f"{API_ROOT}/averageReadings"
MINMAX_READINGS_URL = f"{API_ROOT}/minMaxReadings"


@dataclass
class Reading:
    inside: float
    radiator: float
    outside: float
    added: datetime

    def __post_init__(self):
        self.added = datetime.strptime(self.added, "%Y-%m-%dT%H:%M:%SZ")


@dataclass
class MinMaxReading:
    inside_min: float
    inside_max: float
    radiator_min: float
    radiator_max: float
    outside_min: float
    outside_max: float
    added: datetime

    def __post_init__(self):
        self.added = datetime.strptime(self.added, "%Y-%m-%dT%H:%M:%SZ")


def last_n_readings(limit: int, offset: int = 0) -> list[Reading]:
    r = requests.post(LAST_N_READINGS_URL, json={"limit": limit, "offset": offset})
    return [Reading(**i) for i in r.json()]


def first_n_readings(limit: int, offset: int = 0) -> list[Reading]:
    r = requests.post(FIRST_N_READINGS_URL, json={"limit": limit, "offset": offset})
    return [Reading(**i) for i in r.json()]


# seconds is averaging interval in seconds
def average_readings(seconds: int) -> list[Reading]:
    r = requests.post(AVERAGE_READINGS_URL, json={"seconds": seconds})
    return [Reading(**i) for i in r.json()]


# seconds is minmax search interval in seconds
def minmax_readings(seconds: int) -> list[MinMaxReading]:
    r = requests.post(MINMAX_READINGS_URL, json={"seconds": seconds})
    return [MinMaxReading(**i) for i in r.json()]

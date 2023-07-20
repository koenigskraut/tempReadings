from api import average_readings, minmax_readings
import matplotlib.pyplot as plt
from datetime import datetime
from plot import plot
from api import MinMaxReading


def title_string_avg(minutes: int) -> str:
    if minutes == 1:
        return "Температура на радиаторе, среднее за минуту"
    str_end = ""
    if minutes % 10 == 1 and minutes % 100 != 11:
        str_end = "у"
    elif 2 <= minutes % 10 <= 4 and not 12 <= minutes % 100 <= 14:
        str_end = "ы"
    min_fmt = f"{minutes} минут{str_end}"
    title = f"Температура на радиаторе, среднее за каждые {min_fmt}"
    return title


def title_string_mm(minutes: int) -> str:
    if minutes == 1:
        return "Температура на радиаторе, min/max за минуту"
    str_end = ""
    if minutes % 10 == 1 and minutes % 100 != 11:
        str_end = "у"
    elif 2 <= minutes % 10 <= 4 and not 12 <= minutes % 100 <= 14:
        str_end = "ы"
    min_fmt = f"{minutes} минут{str_end}"
    title = f"Температура на радиаторе, min/max за каждые {min_fmt}"
    return title


def avg_plot(seconds: int):
    avg = average_readings(seconds)[50:][:100]
    added = [i.added for i in avg]
    radiator = [i.radiator for i in avg]
    title = title_string_avg(seconds // 60)
    filename = f"radiator_every_{seconds//60}_min.png"
    plot(added, radiator, title, "радиатор", filename, color="#ff7f0e")


def minmax_process(data: list[tuple[float, float]]) -> list[float]:
    data[0] = data[0][0]
    for i in range(len(data) - 1):
        if abs(data[i + 1][0] - data[i]) > abs(data[i + 1][1] - data[i]):
            data[i + 1] = data[i + 1][0]
        else:
            data[i + 1] = data[i + 1][1]
    return data


def minmax_plot(seconds: int):
    mm = minmax_readings(seconds)[50:][:100]
    added = [i.added for i in mm]
    radiator = minmax_process([(i.radiator_min, i.radiator_max) for i in mm])
    title = title_string_mm(seconds // 60)
    filename = f"radiator_mm_every_{seconds//60}_min.png"
    plot(added, radiator, title, "радиатор", filename, color="#ff7f0e")


if __name__ == "__main__":
    avg_plot(300)
    minmax_plot(300)

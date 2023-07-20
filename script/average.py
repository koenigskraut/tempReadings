from api import average_readings
import matplotlib.pyplot as plt
from datetime import datetime
from plot import plot, ticks_and_labels


def title_string(minutes: int) -> str:
    if minutes == 1:
        return "Температура на улице, среднее за минуту"
    str_end = ""
    if minutes % 10 == 1 and minutes % 100 != 11:
        str_end = "у"
    elif 2 <= minutes % 10 <= 4 and not 12 <= minutes % 100 <= 14:
        str_end = "ы"
    min_fmt = f"{minutes} минут{str_end}"
    title = f"Температура на улице, среднее за каждые {min_fmt}"
    return title


def avg_plot(seconds: int):
    avg = average_readings(seconds)
    added = [i.added for i in avg]
    outside = [i.outside for i in avg]
    title = title_string(seconds // 60)
    filename = f"outside_every_{seconds//60}_min.png"
    plot(added, outside, title, "улица", filename)


if __name__ == "__main__":
    avg_plot(120)
    avg_plot(600)

import matplotlib.pyplot as plt
from datetime import datetime


def ticks_and_labels(dates: list[datetime]) -> tuple[list[datetime], list[str]]:
    xticks = [i for i in dates[:: len(dates) // 18]]
    xlabels = (
        [xticks[0].strftime("%H:%M\n%d.%m")]
        + [i.strftime("%H:%M") for i in xticks[1:-1]]
        + [xticks[-1].strftime("%H:%M\n%d.%m")]
    )
    prev = xticks[0]
    for i, date in enumerate(xticks):
        if date.day != prev.day:
            xlabels[i] = xticks[i].strftime("%H:%M\n%d.%m")
            prev = date
    return (xticks, xlabels)


def plot(dataX, dataY, title, label, filename, color="#1f77b4"):
    fig, ax = plt.subplots()
    fig.set_size_inches(13, 7.31)
    fig.set_dpi(120)
    ax.plot(dataX, dataY, label=label, color=color)

    plt.xticks(*ticks_and_labels(dataX))
    plt.title(title)
    plt.legend()
    plt.grid(visible=True)
    plt.tight_layout()
    plt.savefig(filename)

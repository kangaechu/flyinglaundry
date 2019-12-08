import requests
from bs4 import BeautifulSoup


class Tenki():
    def __init__(self, url):
        res = requests.get(url)
        self.soup = BeautifulSoup(res.text, 'html.parser')

    def weather(self):
        """
        天気を取得
        @return: 1日の24時間分の天気情報のリスト
        """
        css_class = '#forecast-point-1h-today > tr.weather td'
        return [h.get_text().strip() for h in self.soup.select(css_class)]

    def temperature(self):
        """
        気温を取得
        @return: 1日の24時間分の天気情報のリスト
        """
        css_class = '#forecast-point-1h-today > tr.temperature td'
        return [h.get_text().strip() for h in self.soup.select(css_class)]

    def prob_precip(self):
        """
        降水確率を取得
        @return: 1日の24時間分の天気情報のリスト
        """
        css_class = '#forecast-point-1h-today > tr.prob-precip td'
        return [h.get_text().strip() for h in self.soup.select(css_class)]

    def precipitation(self):
        """
        降水量を取得
        @return: 1日の24時間分の天気情報のリスト
        """
        css_class = '#forecast-point-1h-today > tr.precipitation td'
        return [h.get_text().strip() for h in self.soup.select(css_class)]

    def humidity(self):
        """
        湿度を取得
        @return: 1日の24時間分の天気情報のリスト
        """
        css_class = '#forecast-point-1h-today > tr.humidity td'
        return [h.get_text().strip() for h in self.soup.select(css_class)]

    def wind_blow(self):
        """
        風向を取得
        @return: 1日の24時間分の天気情報のリスト
        """
        css_class = '#forecast-point-1h-today > tr.wind-blow td'
        return [h.get_text().strip() for h in self.soup.select(css_class)]

    def wind_speed(self):
        """
        風速を取得
        @return: 1日の24時間分の天気情報のリスト
        """
        css_class = '#forecast-point-1h-today > tr.wind-speed td'
        return [h.get_text().strip() for h in self.soup.select(css_class)]


if __name__ == '__main__':
    url = 'https://tenki.jp/forecast/1/2/1400/1100/1hour.html'

    tenki = Tenki(url)
    print(tenki.weather())
    print(tenki.temperature())
    print(tenki.prob_precip())
    print(tenki.precipitation())
    print(tenki.humidity())
    print(tenki.wind_blow())
    print(tenki.wind_speed())

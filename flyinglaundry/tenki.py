import requests
from bs4 import BeautifulSoup


class Tenki():
    def __init__(self, url):
        res = requests.get(url)
        self.soup = BeautifulSoup(res.text, 'html.parser')

    def weather(self):
        """
        天気を取得
        @return: 48時間分の天気情報のリスト
        """
        css_class_today = '#forecast-point-1h-today > tr.weather td'
        today = [h.get_text().strip() for h in self.soup.select(css_class_today)]
        css_class_tomorrow = '#forecast-point-1h-tomorrow > tr.weather td'
        tomorrow = [h.get_text().strip() for h in self.soup.select(css_class_tomorrow)]
        return today + tomorrow

    def temperature(self):
        """
        気温を取得
        @return: 48時間分の天気情報のリスト
        """
        css_class_today = '#forecast-point-1h-today > tr.temperature td'
        today = [h.get_text().strip() for h in self.soup.select(css_class_today)]
        css_class_tomorrow = '#forecast-point-1h-tomorrow > tr.temperature td'
        tomorrow = [h.get_text().strip() for h in self.soup.select(css_class_tomorrow)]
        return today + tomorrow

    def prob_precip(self):
        """
        降水確率を取得
        @return: 48時間分の天気情報のリスト
        """
        css_class_today = '#forecast-point-1h-today > tr.prob-precip td'
        today = [h.get_text().strip() for h in self.soup.select(css_class_today)]
        css_class_tomorrow = '#forecast-point-1h-tomorrow > tr.prob-precip td'
        tomorrow = [h.get_text().strip() for h in self.soup.select(css_class_tomorrow)]
        return today + tomorrow

    def precipitation(self):
        """
        降水量を取得
        @return: 48時間分の天気情報のリスト
        """
        css_class_today = '#forecast-point-1h-today > tr.precipitation td'
        today = [h.get_text().strip() for h in self.soup.select(css_class_today)]
        css_class_tomorrow = '#forecast-point-1h-tomorrow > tr.precipitation td'
        tomorrow = [h.get_text().strip() for h in self.soup.select(css_class_tomorrow)]
        return today + tomorrow

    def humidity(self):
        """
        湿度を取得
        @return: 48時間分の天気情報のリスト
        """
        css_class_today = '#forecast-point-1h-today > tr.humidity td'
        today = [h.get_text().strip() for h in self.soup.select(css_class_today)]
        css_class_tomorrow = '#forecast-point-1h-tomorrow > tr.humidity td'
        tomorrow = [h.get_text().strip() for h in self.soup.select(css_class_tomorrow)]
        return today + tomorrow

    def wind_blow(self):
        """
        風向を取得
        @return: 48時間分の天気情報のリスト
        """
        css_class_today = '#forecast-point-1h-today > tr.wind-blow td'
        today = [h.get_text().strip() for h in self.soup.select(css_class_today)]
        css_class_tomorrow = '#forecast-point-1h-tomorrow > tr.wind-blow td'
        tomorrow = [h.get_text().strip() for h in self.soup.select(css_class_tomorrow)]
        return today + tomorrow

    def wind_speed(self):
        """
        風速を取得
        @return: 48時間分の天気情報のリスト
        """
        css_class_today = '#forecast-point-1h-today > tr.wind-speed td'
        today = [h.get_text().strip() for h in self.soup.select(css_class_today)]
        css_class_tomorrow = '#forecast-point-1h-tomorrow > tr.wind-speed td'
        tomorrow = [h.get_text().strip() for h in self.soup.select(css_class_tomorrow)]
        return today + tomorrow


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

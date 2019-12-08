from unittest import TestCase

from tenki import Tenki


class TestTenki(TestCase):
    def setUp(self):
        url = 'https://tenki.jp/forecast/1/2/1400/1100/1hour.html'
        self.tenki = Tenki(url)

    def test_weather(self):
        weather = self.tenki.weather()
        self.assertEqual(len(weather), 24)

    def test_temperature(self):
        temperature = self.tenki.temperature()
        self.assertEqual(len(temperature), 24)

    def test_prob_precip(self):
        prob_precip = self.tenki.prob_precip()
        self.assertEqual(len(prob_precip), 24)

    def test_precipitation(self):
        precipitation = self.tenki.precipitation()
        self.assertEqual(len(precipitation), 24)

    def test_humidity(self):
        humidity = self.tenki.humidity()
        self.assertEqual(len(humidity), 24)

    def test_wind_blow(self):
        wind_blow = self.tenki.wind_blow()
        self.assertEqual(len(wind_blow), 24)

    def test_wind_speed(self):
        wind_speed = self.tenki.wind_speed()
        self.assertEqual(len(wind_speed), 24)


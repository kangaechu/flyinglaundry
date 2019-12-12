from tenki import Tenki
from datetime import datetime

if __name__ == '__main__':
    url = 'https://tenki.jp/forecast/1/2/1400/1100/1hour.html'
    tenki = Tenki(url)

    current_hour = datetime.now().hour
    wind_speed_24h = tenki.wind_speed()[current_hour:current_hour + 24]
    wind_blow_24h = tenki.wind_blow()[current_hour:current_hour + 24]

    flying_laundry = False
    for wind_speed, wind_blow in zip(wind_speed_24h, tenki.wind_blow()):
        # 北風で風速5m以上だと洗濯物が吹っ飛ぶ
        if '北' in wind_blow and int(wind_speed) >= 5:
            flying_laundry = True
    if flying_laundry:
        print("洗濯物ふっとび注意")

from tenki import Tenki

if __name__ == '__main__':
    url = 'https://tenki.jp/forecast/1/2/1400/1100/1hour.html'
    tenki = Tenki(url)
    flying_laundry = False
    for wind_speed, wind_blow in zip(tenki.wind_speed(), tenki.wind_blow()):
        # 北風で風速5m以上だと洗濯物が吹っ飛ぶ
        if '北' in wind_blow and int(wind_speed) >= 5:
            flying_laundry = True
    if flying_laundry:
        print("洗濯物ふっとび注意")


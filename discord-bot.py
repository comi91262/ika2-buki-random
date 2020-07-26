# coding: UTF-8

import traceback
import discord
import os
import random

client = discord.Client()

# 環境変数からBOTのTOKENと自分のユーザIDを取得
TOKEN = 'NjkxNTkyMjI3OTk0MjA2MjI5.XniNdA.GR4yGGrYyYb2EpRca5-jIDV7yIg'

bukiList = [
        'スプラシューター'
        ,'スプラシューターコラボ'
        ,'ヒーローシューター レプリカ'
        ,'オクタシューター レプリカ'
        ,'スプラシューターベッチュー'
        ,'プライムシューター'
        ,'プライムシューターコラボ'
        ,'プライムシューターベッチュー'
        ,'シャープマーカー'
        ,'シャープマーカーネオ'
        ,'わかばシューター'
        ,'もみじシューター'
        ,'おちばシューター'
        ,'プロモデラーMG'
        ,'プロモデラーRG'
        ,'プロモデラーPG'
        ,'N-ZAP85'
        ,'N-ZAP89'
        ,'N-ZAP83'
        ,'.52ガロン'
        ,'.52ガロンデコ'
        ,'.52ガロンベッチュー'
        ,'ジェットスイーパー'
        ,'ジェットスイーパーカスタム'
        ,'L3リールガン'
        ,'L3リールガンD'
        ,'L3リールガンベッチュー'
        ,'.96ガロン'
        ,'.96ガロンデコ'
        ,'H3リールガン'
        ,'H3リールガンD'
        ,'H3リールガンチェリー'
        ,'ボールドマーカー'
        ,'ボールドマーカーネオ'
        ,'ボールドマーカー7'
        ,'ボトルガイザー'
        ,'ボトルガイザーフォイル'
        ,'ホットブラスター'
        ,'ホットブラスターカスタム'
        ,'ヒーローブラスター レプリカ'
        ,'ラピッドブラスター'
        ,'ラピッドブラスターデコ'
        ,'ラピッドブラスターベッチュー'
        ,'ノヴァブラスター'
        ,'ノヴァブラスターネオ'
        ,'ノヴァブラスターベッチュー'
        ,'クラッシュブラスター'
        ,'クラッシュブラスターネオ'
        ,'Rブラスターエリート'
        ,'Rブラスターエリートデコ'
,'ロングブラスター'
,'ロングブラスターカスタム'
,'ロングブラスターネクロ'
,'スプラローラー'
,'スプラローラーコラボ'
,'ヒーローローラー レプリカ'
,'スプラローラーベッチュー'
,'ダイナモローラー'
,'ダイナモローラーテスラ'
,'ダイナモローラーベッチュー'
,'カーボンローラー'
,'カーボンローラーデコ'
,'ヴァリアブルローラー'
,'ヴァリアブルローラーフォイル'
,'ホクサイ'
,'ヒーローブラシ レプリカ'
,'ホクサイ・ヒュー'
,'ホクサイベッチュー'
,'パブロ'
,'パブロ・ヒュー'
,'パーマネント・パブロ'
,'14式竹筒銃・甲'
,'14式竹筒銃・乙'
,'14式竹筒銃・丙'
,'バケットスロッシャー'
,'ヒーロースロッシャー レプリカ'
,'バケットスロッシャーデコ'
,'バケットスロッシャーソーダ'
,'ヒッセン'
,'ヒッセン・ヒュー'
,'スクリュースロッシャー'
,'スクリュースロッシャーネオ'
,'スクリュースロッシャーベッチュー'
,'エクスプロッシャー'
,'エクスプロッシャーカスタム'
,'オーバーフロッシャー'
,'オーバーフロッシャーデコ'
,'バレルスピナー'
,'ヒーロースピナー レプリカ'
,'バレルスピナーデコ'
,'バレルスピナーリミックス'
,'スプラスピナー'
,'スプラスピナーコラボ'
,'スプラスピナーベッチュー'
,'ハイドラント'
,'ハイドラントカスタム'
,'クーゲルシュライバー'
,'クーゲルシュライバー・ヒュー'
,'ノーチラス47'
,'ノーチラス79'
,'スプラマニューバー'
,'スプラマニューバーコラボ'
,'ヒーローマニューバー レプリカ'
,'スプラマニューバーベッチュー'
,'スパッタリー'
,'スパッタリー・ヒュー'
,'スパッタリークリア'
,'デュアルスイーパー'
,'デュアルスイーパーカスタム'
,'ケルビン525'
,'ケルビン525デコ'
,'ケルビン525ベッチュー'
,'クアッドホッパーブラック'
,'クアッドホッパーホワイト'
,'パラシェルター'
,'ヒーローシェルター レプリカ'
,'パラシェルターソレーラ'
,'キャンピングシェルター'
,'キャンピングシェルターソレーラ'
,'キャンピングシェルターカーモ'
,'スパイガジェット'
,'スパイガジェットソレーラ'
,'スパイガジェットベッチュー' 
]

chargerList = [
        'スプラチャージャー'
        ,'スプラスコープ'
        ,'スプラチャージャーコラボ'
        ,'スプラスコープコラボ'
        ,'ヒーローチャージャー レプリカ'
        ,'スプラチャージャーベッチュー'
        ,'スプラスコープベッチュー'
        ,'リッター4K'
        ,'4Kスコープ'
        ,'リッター4Kカスタム'
        ,'4Kスコープカスタム'
        ,'ソイチューバー'
        ,'ソイチューバーカスタム'
        ,'スクイックリンα'
        ,'スクイックリンβ'
        ,'スクイックリンγ'
        ]

@client.event
async def on_ready():
    print('ログインしました')


@client.event
async def on_message(message):
    # メッセージ送信者がBotだった場合は無視する
    if message.author.bot:
        return
    if message.content == '/buki':
        chargerOrElse = random.randint(0, 99)
        if chargerOrElse == 50:
            randomInt = random.randint(0, 15)
            await message.channel.send("大当たり! " + chargerList[randomInt])
        else:
            randomInt = random.randint(0, 122)
            await message.channel.send(bukiList[randomInt])




client.run(TOKEN)


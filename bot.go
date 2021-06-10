package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"strings"
	"time"
)

var TOKEN = ""

var bukiList = []string{
	"スプラシューター",
	"スプラシューターコラボ",
	"ヒーローシューター レプリカ",
	"オクタシューター レプリカ",
	"スプラシューターベッチュー",
	"プライムシューター",
	"プライムシューターコラボ",
	"プライムシューターベッチュー",
	"シャープマーカー",
	"シャープマーカーネオ",
	"わかばシューター",
	"もみじシューター",
	"おちばシューター",
	"プロモデラーMG",
	"プロモデラーRG",
	"プロモデラーPG",
	"N-ZAP85",
	"N-ZAP89",
	"N-ZAP83",
	".52ガロン",
	".52ガロンデコ",
	".52ガロンベッチュー",
	"ジェットスイーパー",
	"ジェットスイーパーカスタム",
	"L3リールガン",
	"L3リールガンD",
	"L3リールガンベッチュー",
	".96ガロン",
	".96ガロンデコ",
	"H3リールガン",
	"H3リールガンD",
	"H3リールガンチェリー",
	"ボールドマーカー",
	"ボールドマーカーネオ",
	"ボールドマーカー7",
	"ボトルガイザー",
	"ボトルガイザーフォイル",
	"ホットブラスター",
	"ホットブラスターカスタム",
	"ヒーローブラスター レプリカ",
	"ラピッドブラスター",
	"ラピッドブラスターデコ",
	"ラピッドブラスターベッチュー",
	"ノヴァブラスター",
	"ノヴァブラスターネオ",
	"ノヴァブラスターベッチュー",
	"クラッシュブラスター",
	"クラッシュブラスターネオ",
	"Rブラスターエリート",
	"Rブラスターエリートデコ",
	"ロングブラスター",
	"ロングブラスターカスタム",
	"ロングブラスターネクロ",
	"スプラローラー",
	"スプラローラーコラボ",
	"ヒーローローラー レプリカ",
	"スプラローラーベッチュー",
	"ダイナモローラー",
	"ダイナモローラーテスラ",
	"ダイナモローラーベッチュー",
	"カーボンローラー",
	"カーボンローラーデコ",
	"ヴァリアブルローラー",
	"ヴァリアブルローラーフォイル",
	"ホクサイ",
	"ヒーローブラシ レプリカ",
	"ホクサイ・ヒュー",
	"ホクサイベッチュー",
	"パブロ",
	"パブロ・ヒュー",
	"パーマネント・パブロ",
	"14式竹筒銃・甲",
	"14式竹筒銃・乙",
	"14式竹筒銃・丙",
	"バケットスロッシャー",
	"ヒーロースロッシャー レプリカ",
	"バケットスロッシャーデコ",
	"バケットスロッシャーソーダ",
	"ヒッセン",
	"ヒッセン・ヒュー",
	"スクリュースロッシャー",
	"スクリュースロッシャーネオ",
	"スクリュースロッシャーベッチュー",
	"エクスプロッシャー",
	"エクスプロッシャーカスタム",
	"オーバーフロッシャー",
	"オーバーフロッシャーデコ",
	"バレルスピナー",
	"ヒーロースピナー レプリカ",
	"バレルスピナーデコ",
	"バレルスピナーリミックス",
	"スプラスピナー",
	"スプラスピナーコラボ",
	"スプラスピナーベッチュー",
	"ハイドラント",
	"ハイドラントカスタム",
	"クーゲルシュライバー",
	"クーゲルシュライバー・ヒュー",
	"ノーチラス47",
	"ノーチラス79",
	"スプラマニューバー",
	"スプラマニューバーコラボ",
	"ヒーローマニューバー レプリカ",
	"スプラマニューバーベッチュー",
	"スパッタリー",
	"スパッタリー・ヒュー",
	"スパッタリークリア",
	"デュアルスイーパー",
	"デュアルスイーパーカスタム",
	"ケルビン525",
	"ケルビン525デコ",
	"ケルビン525ベッチュー",
	"クアッドホッパーブラック",
	"クアッドホッパーホワイト",
	"パラシェルター",
	"ヒーローシェルター レプリカ",
	"パラシェルターソレーラ",
	"キャンピングシェルター",
	"キャンピングシェルターソレーラ",
	"キャンピングシェルターカーモ",
	"スパイガジェット",
	"スパイガジェットソレーラ",
	"スパイガジェットベッチュー",
}

var chargerList = []string{
	"スプラチャージャー",
	"スプラスコープ",
	"スプラチャージャーコラボ",
	"スプラスコープコラボ",
	"ヒーローチャージャー レプリカ",
	"スプラチャージャーベッチュー",
	"スプラスコープベッチュー",
	"リッター4K",
	"4Kスコープ",
	"リッター4Kカスタム",
	"4Kスコープカスタム",
	"ソイチューバー",
	"ソイチューバーカスタム",
	"スクイックリンα",
	"スクイックリンβ",
	"スクイックリンγ",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	discord, err := discordgo.New("Bot " + TOKEN)

	if err != nil {
		log.Println("Error logging in")
	}

	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Listening...")

	<-make(chan bool)
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID) //チャンネル取得
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}
	if m.Author.Bot {
		return
	}
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	var userName string
	if m.Member.Nick == "" {
		userName = m.Author.Username
	} else {
		userName = m.Member.Nick
	}

	switch {
	case strings.HasPrefix(m.Content, "/buki"):
		chargerOrElse := rand.Intn(100)
		if chargerOrElse == 50 {
			randomInt := rand.Intn(16)
			msg := userName + " << 大当たり! " + chargerList[randomInt]
			sendMessage(s, c, msg)
		} else {
			randomInt := rand.Intn(123)
			msg := userName + " << " + bukiList[randomInt]
			sendMessage(s, c, msg)
		}
	}
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

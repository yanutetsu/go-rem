package rem

import (
	"fmt"
	"math/rand"
	"time"
)

var remsays = [...]string{
	"姉様姉様。どうやら少し混乱されているみたいですお客様",
	"鬼がかってますね。",
	"嘘をついていることぐらいレムにもわかります。その嘘の理由が話せないでいることも、わかります。だから信じさせようだとか、嘘で丸め込もうだとか、そんな風に自分を追い詰めたりする必要、どこにもないんですよ。だってレムは　丸ごと信じていますから。",
	"いえ、そんなことは・・・レムも少し、かなり少し、とても少し気になるのは事実ですから",
	"穀潰しの発言ですよ。聞きました姉様？",
}

// Print prints a message with rem's random message.
// It works like fmt.Print.
// You can use it only by changing from `fmt` to `rem`.
func Print(a ...interface{}) {
	msg := fmt.Sprint(a...)

	rand.Seed(time.Now().UnixNano())
	fmt.Printf("レム＜%s %s", remsays[rand.Intn(5)], msg)
}
